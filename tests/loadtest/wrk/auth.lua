token = nil
path  = "/api/v1/auth/login"
method = "POST"
headers = {}
body = '{"email":"37445","password":"9420","rememberMe":false}'

request = function()
    headers['Content-Type'] = "application/json"
    if token ~= nil then
        headers["Authorization"] = token
    end

    return wrk.format(method, path, headers, body)
end

response = function(status, headers, body)
    if not token and status == 200 then
        token = headers["Authorization"]
        path  = "/__auth__"
        method = "GET"
    end
end
