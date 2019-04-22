package sendcloud

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
)

type TaskInfo struct {
	GMTCreated     string `json:"gmtCreated"`
	GMTtUpdated    string `json:"gmtUpdated"`
	MailListTaskId int    `json:"maillistTaskId"`
	ApiUser        string `json:"apiUser"`
	AddressList    string `json:"addressList"`
	MemberCount    int    `json:"memberCount"`
	Status         string `json:"status"`
}

// --------------------------------------------------------------------------------
type SendTemplateRsp struct {
	Result     bool   `json:"result"`
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
	Info       struct {
		EmailIdList    []string `json:"emailIdList"`
		MailListTaskId []int    `json:"maillistTaskId"`
	} `json:"info"`
}

type To struct {
	toList map[string]map[string]string
}

func NewTo() *To {
	var tl = &To{}
	tl.toList = make(map[string]map[string]string)
	return tl
}

func (this *To) Add(to string, param map[string]string) {
	this.toList[to] = param
}

func (this *To) Del(to string) {
	delete(this.toList, to)
}

func (this *To) Len() int {
	return len(this.toList)
}

// SendMailWithTemplate 模板发送
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

	substitutionVarsBytes, err := json.Marshal(toMap)
	if err != nil {
		return nil, err
	}

	var substitutionVars = string(substitutionVarsBytes)
	params := url.Values{
		"from":               {from},
		"fromName":           {fromName},
		"replyTo":            {replyTo},
		"subject":            {subject},
		"templateInvokeName": {invokeName},
		"xsmtpapi":           {substitutionVars},
	}

	if err = this.doRequestWithFile(kSendTemplate, params, "attachments", filename, &result); err != nil {
		return nil, err
	}

	if result.Result == false {
		return nil, errors.New(fmt.Sprintf("%d-%s", result.StatusCode, result.Message))
	}

	return result, err
}

// --------------------------------------------------------------------------------
// SendTemplateMailWithAddressList 向邮件地址列表发送邮件
// addressList  string 是   邮件地址列表
// invokeName   string 是   邮件模板调用名称
// from         string 是   发件人地址
// fromName     string 否   发件人名称
// replyTo      string 否   设置用户默认的回复邮件地址. 如果 replyTo 没有或者为空, 则默认的回复邮件地址为 from
// subject      string *    邮件标题
func (this *Client) SendTemplateMailToAddressList(addressList, invokeName, from, fromName, replyTo, subject string, filename []string) (result *SendTemplateRsp, err error) {
	params := url.Values{
		"to":                 {addressList},
		"from":               {from},
		"fromName":           {fromName},
		"replyTo":            {replyTo},
		"subject":            {subject},
		"templateInvokeName": {invokeName},
		"useAddressList":     {"true"},
	}
	if err = this.doRequestWithFile(kSendTemplate, params, "attachments", filename, &result); err != nil {
		return nil, err
	}

	if result.Result == false {
		return nil, errors.New(fmt.Sprintf("%d-%s", result.StatusCode, result.Message))
	}

	return result, err
}

// --------------------------------------------------------------------------------
type GetTaskInfoRsp struct {
	Result     bool   `json:"result"`
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
	Info       struct {
		Data *TaskInfo `json:"data"`
	} `json:"info"`
}

// GetTaskInfo 获取邮件地址列表发送任务信息
// mailListTaskId   int  是  返回的mailListTaskId
func (this *Client) GetTaskInfo(mailListTaskId int) (result *GetTaskInfoRsp, err error) {
	params := url.Values{}
	params.Add("maillistTaskId", fmt.Sprintf("%d", mailListTaskId))

	if err = this.doRequest(kMailTaskInfo, params, &result); err != nil {
		return nil, err
	}

	if result.Result == false {
		return nil, errors.New(fmt.Sprintf("%d-%s", result.StatusCode, result.Message))
	}

	return result, err
}
