local cjson = require "cjson"

require "middleware";

local path = "/__auth__"

function authAdmin()
    if ngx.req.get_method() == "OPTIONS" then
        return nil
    end

    local res = ngx.location.capture(path)
    if res.header["X-User-Auth"] == nil then
        ngx.exit(ngx.HTTP_UNAUTHORIZED)
    end

    if string.find(res.header["X-User-Auth"], "ROLE_MASTER") == nil then
        ngx.exit(ngx.HTTP_FORBIDDEN)
    end
end

function auth(query)
    if query == nil then
        query = ""
    end

    local res = ngx.location.capture(path .. query)
    ngx.req.set_header("X-User-Id", res.header["X-User-Id"])
    ngx.req.set_header("X-User-Data", res.header["X-User-Data"])
    ngx.req.clear_header("Authorization")

    return res
end

function authBody(key, query)
    local res = auth(query)

    if ngx.req.get_method() == "POST" or ngx.req.get_method() == "PUT" then
        ngx.req.read_body()

        if ngx.var.request_body == nil then
            return nil
        end

        local ok, t = pcall(cjson.decode, ngx.var.request_body)
        if not ok then
            return nil
        end

        t[key] = tonumber(res.header["X-User-Id"])

        return ngx.req.set_body_data(cjson.encode(t))
    end

    return nil
end

function modelsManager()
    local res = auth("?data=position")

    ngx.log(ngx.STDERR, res.body)
    ngx.log(ngx.STDERR, res.status)
    ngx.log(ngx.STDERR, ngx.var.uri)
    ngx.log(ngx.STDERR, res.header["X-User-Data"])
    ngx.req.set_uri("/character/model/" .. res.header["X-User-Id"])


    return res
end

function authLogin()
    if ngx.req.get_method() == "POST" then
        ngx.req.read_body()

        if ngx.var.request_body == nil then
            return nil
        end

        local ok, req = pcall(cjson.decode, ngx.var.request_body)
        if not ok then
            return nil
        end

        if req.email ~= nil then
            req.login = req.email
        end

        local webRes = ngx.location.capture("/login", { method = ngx.HTTP_POST,
            body = cjson.encode(req) })

        local ok, res = pcall(cjson.decode, webRes.body)
        if not ok then
            return nil
        end

        if req.firebase_token ~= nil then
            local pushToken = {}
            pushToken.id = res["modelId"]
            pushToken.token = req.firebase_token

            local pushRes = ngx.location.capture("/save_token", { method = ngx.HTTP_PUT,
                body = cjson.encode(pushToken) })

            ngx.log(ngx.INFO, "push:" .. pushRes.status)

            local redis = require "resty.redis"
            local red = redis:new()
            red:set_timeout(1000)

            local redisHost = os.getenv("REDIS_HOST")
            local ok, err = red:connect(redisHost, 6379)
            if not ok then
                ngx.print("failed connect: ", err)
                return
            end

            local ok, err = red:lpush("tokens::" .. res["modelId"], res["token"])
            if not ok then
                ngx.print("failed to lpush token: ", err)
                return
            end
        end

        for k, v in pairs(webRes.header) do
            ngx.header[k] = v
        end
        ngx.status = webRes.status

        local result = {}
        result.id = res["modelId"]
        result.api_key = res.token

        ngx.say(cjson.encode(result))
    end

end

function authLogout()
    ngx.header['Set-Cookie'] = 'Authorization=; Path=/; Domain=evarun.ru;'
    return
end

function manaLevel()
    local res = auth("?data=position")

    if res.status ~= 200 then
        ngx.status = res.status
        ngx.print(res.Body)
        return
    end

    local ok, data = pcall(cjson.decode, res.header["X-User-Data"])
    if not ok or data.position == nil then
        ngx.header.content_type = 'application/json'
        ngx.status = 200
        ngx.print('{"id":0,"manaLevel":0}')
        return
    end

    ngx.header.content_type = 'application/json'
    ngx.print(cjson.encode(data.position))
    return
end

function config(key)
    if ngx.req.get_method() == "GET" or ngx.req.get_method() == "POST" then

        local redis = require "resty.redis"
        local red = redis:new()

        red:set_timeout(1000)

        local redisHost = os.getenv("REDIS_HOST")
        local ok, err = red:connect(redisHost, 6379)
        if err ~= nil then
            ngx.status = 500
            ngx.print(err)
            return
        end

        if not ok then
            ngx.status = 500
            ngx.print("")
            return
        end

        if ngx.req.get_method() == "GET" then
            local res, err = red:get(key);
            if res == ngx.null then
                ngx.status = 404
                ngx.print("")
                return
            else
                ngx.header.content_type = 'application/json'
                ngx.print(res)
                return
            end
        end

        if ngx.req.get_method() == "POST" then
            ngx.req.read_body()

            if ngx.var.request_body == nil then
                ngx.header.content_type = 'application/json'
                ngx.status = 400
                ngx.print('{"body":"empty"}')
                return
            end

            local ok, res = pcall(cjson.decode, ngx.var.request_body)
            if not ok then
                ngx.header.content_type = 'application/json'
                ngx.status = 400
                ngx.print('{"body":"invalid json"}')
                return
            end

            ok, err = red:set(key, ngx.var.request_body)
            if not ok then
                ngx.print("failed to set dog: ", err)
                return
            end
            ngx.status = 200
            ngx.print("")
            return
        end

        ngx.status = 404
        ngx.print("")
    else
        ngx.status = 405
        ngx.print("")
    end
end

function addCorsHeaders()
    local origin = "*"
    if ngx.req.get_headers()["Origin"] ~= nil then
        origin = ngx.req.get_headers()["Origin"]
    end

    ngx.header["Access-Control-Allow-Origin"] = origin
    ngx.header["Access-Control-Allow-Methods"] = "POST,GET,OPTIONS,PUT,DELETE,HEAD,PATCH"
    ngx.header["Access-Control-Allow-Credentials"] = "true"
    ngx.header["Access-Control-Allow-Headers"] = "Origin,X-Requested-With,Content-Type,Accept,Authorization"
end
