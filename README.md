## sendcloud

目前已实现以下功能：

* 1、模板管理（添加、删除、更新、获取）
* 2、地址列表管理（添加、删除、更新、获取）
* 3、地址列表成员管理（添加、删除、更新、获取）
* 4、邮件发送（模板发送、地址列表发送、附件支持）

使用方法如下：

设置 API 信息：

```
var c = sendcloud.New("api_user", "api_key")
```

发送邮件：

```
var to = sendcloud.NewTo()
to.Add("917996695@qq.com", nil)
var result, err = c.SendTemplateMail(to, "invoke name", "from mail", "from name", "replay to email address", "subject", nil)
```