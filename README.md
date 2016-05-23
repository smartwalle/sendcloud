## sendcloud

对 [SendCloud](http://sendcloud.sohu.com) 发送邮件 API 的封装，现已将 API 接口切换到了 v2 版本，如果想使用 v1 版本的 API，请 Checkout v1 Tag。

v2 版本新功能：

* 1、模板管理（添加、删除、更新、获取）
* 2、地址列表管理（添加、删除、更新、获取）
* 3、地址列表成员管理（添加、删除、更新、获取）
* 4、邮件发送（模板发送、地址列表发送、附件支持）

使用方法如下：

设置 API 信息：

```
UpdateApiInfo("api_user", "api_key")
```

发送邮件：

```
var to = make([]map[string]string, 1)
to[0] = map[string]string{"to":"917996695@qq.com", "%url%": "http://www.baidu.com"}
var ok, err, result = SendTemplateMail("template name", "from mail", "from name", "replay to email address", "subject", to, nil)
```

返回数据说明：

    ok     发送是否成功
    err    错误信息
    result 接口返回的 json 数据
