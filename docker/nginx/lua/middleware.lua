local cjson = require "cjson"

function modelsManagerMiddleware()
    local res = auth("?data=position")

    if res.status ~= 200 then
        ngx.status = 401
        ngx.say('Unauthorized')
        return ngx.exit(401)
    end

    if ngx.req.get_method() == "POST" or ngx.req.get_method() == "PUT" then
        ngx.req.read_body()

        if ngx.var.request_body == nil then
            return nil
        end

        local ok, reqBody = pcall(cjson.decode, ngx.var.request_body)
        if not ok then
            return nil
        end

        local ok, resBody = pcall(cjson.decode, res.header["X-User-Data"])
        if not ok then
            return nil
        end

        if reqBody.data == nil then
            reqBody.data = {}
        end

        reqBody.data["location"] = resBody.position

        ngx.log(ngx.INFO, cjson.encode(reqBody))

        ngx.req.set_body_data(cjson.encode(reqBody))
    end

    ngx.req.set_uri("/character/model/" .. res.header["X-User-Id"])

    return res
end

function billingAccountInfoMiddleware()
    local res = auth()

    ngx.req.set_uri("/api/billing/info/getbalance")
    ngx.req.set_uri_args("characterId=" .. res.header["X-User-Id"])

    return res
end

function billingGetTransfersMiddleware()
    local res = auth()

    ngx.req.set_uri("/api/billing/info/gettransfers")
    ngx.req.set_uri_args("characterId=" .. res.header["X-User-Id"])

    return res
end

function billingTransferMiddleware()
    local res = auth()

    if ngx.req.get_method() == "POST" then
        ngx.req.read_body()

        if ngx.var.request_body == nil then
            return nil
        end

        local ok, reqBody = pcall(cjson.decode, ngx.var.request_body)
        if not ok then
            return nil
        end

        ngx.req.set_method(ngx.HTTP_GET)
        ngx.req.set_uri("/api/billing/transfer/maketransfersinsin")
        ngx.req.set_uri_args("character1=" .. res.header["X-User-Id"] .. "&character2=" .. reqBody.sin_to .. "&amount=" .. reqBody.amount .. "&comment=" .. reqBody.comment)
    end

    return res
end

function authLoginRequestMiddleware()
    if ngx.req.get_method() == "POST" then
        ngx.req.read_body()

        if ngx.var.request_body == nil then
            return nil
        end

        local ok, t = pcall(cjson.decode, ngx.var.request_body)
        if not ok then
            return nil
        end

        t["login"] = t["email"]

        return ngx.req.set_body_data(cjson.encode(t))
    end

    return nil
end

function authLoginResponseMiddleware()
    if ngx.req.get_method() == "POST" then

        if ngx.var.request_body == nil then
            return nil
        end

        local ok, req = pcall(cjson.decode, ngx.var.request_body)
        if not ok then
            return nil
        end

        local ok, res = pcall(cjson.decode, ngx.arg[1])
        if not ok then
            return nil
        end

        local pushToken = {}
        pushToken.id = res["modelId"]
        pushToken.token = req.firebase_token

        ngx.log(ngx.INFO, cjson.encode(pushToken))

        local pushRes = ngx.location.capture(os.getenv("PUSH_HOST") .. "/save_token", { method = ngx.HTTP_PUT,
            body = cjson.encode(pushToken) })

        ngx.log(ngx.INFO, pushRes.body)

        return nil
    end

    return nil
end