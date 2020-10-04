set = nil
path  = "/api/v1/config/testkey"
method = "POST"
headers = {}
body = '{"key":"test"}'

request = function()
    headers['Content-Type'] = "application/json"

    return wrk.format(method, path, headers, body)
end

response = function(status, headers, body)
    if not set and status == 200 then
        method = "GET"
        body = ""
    end
end
