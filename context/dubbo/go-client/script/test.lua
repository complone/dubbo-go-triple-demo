
local counter = 1
local threads  = {}

local json = require("json")

local req  = {

    group = "dubbo-test",
    version = "1.0.0",
    method= "GetUserByName",
    types = "types",
    values = "tcccccccc"
}

function setup(thread)
    -- 给每个线程设置一个id参数
    thread:set("id",counter)
    table.insert(threads,thread)
    counter = counter + 1

end

function init(args)
    -- 初始化两个参数 每个线程都有独立的 requests, response 参数
    requests = 0
    response = 0
    -- 打印线程被创建的消息，打印完后，线程正式开始运行
    local msg = " thread %d created"
    print(msg.format(counter))
end

function request()

    if method == "POST" then
        wrk.headers["Content-Type"] = "application/x-www-form-urlencoded; charset=UTF-8"
        wrk.headers["User-Agent"] = "wrk"
        wrk.headers["Connection"] = "keep-alive"
        wrk.body = "{\"types\":\"string\",\"values\":\"tc\" }"

        print("本次请求包体为: %s",string.len(wrk.body))
    end
    return wrk.format("POST",wrk.path,wrk.headers,json.encode(req))
end

function response(status,headers,body)
    if status ~= 200 then
        print(body)
    return
    end

    -- 打印响应体
    local resp = json.decode(body)
    print(json.encode(body)..'-->'..body)

end


