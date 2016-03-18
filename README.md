## sendcloud

对 [SendCloud](http://sendcloud.sohu.com) 发送邮件 API 的封装，目前只实现了模板发送，其它方式以后有用到在做考虑。

使用方法如下：

设置 API 信息：

    UpdateApiInfo("api_user", "api_key")

发送邮件：

    var to = make([]map[string]string, 1)
    to[0] = map[string]string{"to":"917996695@qq.com", "%url%": "http://www.baidu.com"}
    var ok, err, result = SendMailWithTemplate("template name", "from mail", "from name", "replay to email address", "subject", to)

返回数据说明：

    ok     发送是否成功
    err    错误信息
    result 接口返回的 json 数据
