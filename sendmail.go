package sendcloud

import (
	"encoding/json"
	"fmt"
	"net/url"
)

const (
	kSendTemplate = "http://api.sendcloud.net/apiv2/mail/sendtemplate"
	kMailTaskInfo = "http://api.sendcloud.net/apiv2/mail/taskinfo"
)

// SendTemplateMail 模板发送 https://www.sendcloud.net/doc/email_v2/send_email/#_2
// invokeName   string 是   邮件模板调用名称
// from         string 是   发件人地址
// fromName     string 否   发件人名称
// replyTo      string 否   设置用户默认的回复邮件地址. 如果 replyTo 没有或者为空, 则默认的回复邮件地址为 from
// subject      string *    邮件标题
func (this *Client) SendTemplateMail(to *To, invokeName, from, fromName, replyTo, subject string, filename []string) (result *SendTemplateRsp, err error) {
	var toMap = map[string]interface{}{}
	var mailList = make([]string, to.Len())
	var sub = map[string][]string{}

	var index = 0
	for addr, param := range to.toList {
		mailList[index] = addr

		for key, value := range param {
			if _, ok := sub[key]; !ok {
				sub[key] = make([]string, to.Len())
			}
			sub[key][index] = value
		}

		index++
	}
	toMap["to"] = mailList
	if len(sub) > 0 {
		toMap["sub"] = sub
	}

	toBytes, err := json.Marshal(toMap)
	if err != nil {
		return nil, err
	}

	var values = url.Values{
		"from":               {from},
		"fromName":           {fromName},
		"replyTo":            {replyTo},
		"subject":            {subject},
		"templateInvokeName": {invokeName},
		"xsmtpapi":           {string(toBytes)},
	}

	if err = this.doRequestWithFile(kSendTemplate, values, "attachments", filename, &result); err != nil {
		return nil, err
	}

	if result.Result == false {
		return nil, fmt.Errorf("%d-%s", result.StatusCode, result.Message)
	}

	return result, err
}

// SendTemplateMailToAddressList 向邮件地址列表发送邮件 https://www.sendcloud.net/doc/email_v2/send_email/#_2
// addressList  string 是   邮件地址列表
// invokeName   string 是   邮件模板调用名称
// from         string 是   发件人地址
// fromName     string 否   发件人名称
// replyTo      string 否   设置用户默认的回复邮件地址. 如果 replyTo 没有或者为空, 则默认的回复邮件地址为 from
// subject      string *    邮件标题
func (this *Client) SendTemplateMailToAddressList(addressList, invokeName, from, fromName, replyTo, subject string, filename []string) (result *SendTemplateRsp, err error) {
	var values = url.Values{
		"to":                 {addressList},
		"from":               {from},
		"fromName":           {fromName},
		"replyTo":            {replyTo},
		"subject":            {subject},
		"templateInvokeName": {invokeName},
		"useAddressList":     {"true"},
	}
	if err = this.doRequestWithFile(kSendTemplate, values, "attachments", filename, &result); err != nil {
		return nil, err
	}

	if result.Result == false {
		return nil, fmt.Errorf("%d-%s", result.StatusCode, result.Message)
	}

	return result, err
}

// GetTaskInfo 获取邮件地址列表发送任务信息 https://www.sendcloud.net/doc/email_v2/send_email/#_4
// mailListTaskId   int  是  返回的mailListTaskId
func (this *Client) GetTaskInfo(mailListTaskId int) (result *GetTaskInfoRsp, err error) {
	var values = url.Values{}
	values.Add("maillistTaskId", fmt.Sprintf("%d", mailListTaskId))

	if err = this.doRequest(kMailTaskInfo, values, &result); err != nil {
		return nil, err
	}

	if result.Result == false {
		return nil, fmt.Errorf("%d-%s", result.StatusCode, result.Message)
	}

	return result, err
}
