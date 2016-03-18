package v2

import (
	"net/url"
	"bytes"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"errors"
	"fmt"
)

const (
	SEND_CLOUD_TEMPLATE_MAIL_API_URL   = "http://api.sendcloud.net/apiv2/mail/sendtemplate"
	SEND_CLOUD_TEMPLATE_LIST_API_URL   = "http://api.sendcloud.net/apiv2/template/list"
	SEND_CLOUD_TEMPLATE_ADD_API_URL    = "http://api.sendcloud.net/apiv2/template/add"
	SEND_CLOUD_TEMPLATE_DELETE_API_URL = "http://api.sendcloud.net/apiv2/template/delete"
	SEND_CLOUD_TEMPLATE_UPDATE_API_URL = "http://api.sendcloud.net/apiv2/template/update"
)

var (
	MailApiUser = ""
	MailApiKey  = ""
)

func UpdateApiInfo(apiUser, apiKey string) {
	MailApiUser = apiUser
	MailApiKey  = apiKey
}

////////////////////////////////////////////////////////////////////////////////
// SendMailWithTemplate 模板发送
// templateInvokeName   string 是   邮件模板调用名称
// from                 string 是   发件人地址
// fromName             string 否   发件人名称
// replyTo              string 否   设置用户默认的回复邮件地址. 如果 replyTo 没有或者为空, 则默认的回复邮件地址为 from
// subject              string *    邮件标题
func SendMailWithTemplate(template, from, fromName, replyTo, subject string, toList []map[string]string) (bool, error, string) {
	var toMap = map[string]interface{}{}
	var toMailList = make([]string, len(toList))
	var sub = map[string][]string{}

	for index, item := range toList {
		for key, value := range item {
			if key == "to" {
				toMailList[index] = value
			} else {
				if _, ok := sub[key]; !ok {
					sub[key] = make([]string, len(toList))
				}
				sub[key][index] = value
			}
		}
	}
	toMap["to"] = toMailList
	if len(sub) > 0 {
		toMap["sub"] = sub
	}

	var substitutionVarsBytes, err = json.Marshal(toMap)
	if err != nil {
		return false ,err, ""
	}

	var substitutionVars  = string(substitutionVarsBytes)
	params := url.Values {
		"from":     {from},
		"fromName": {fromName},
		"replyTo":  {replyTo},
		"subject":  {subject},
		"templateInvokeName": {template},
		"xsmtpapi":    {substitutionVars},
	}

	return doRequest(SEND_CLOUD_TEMPLATE_MAIL_API_URL, params)
}

////////////////////////////////////////////////////////////////////////////////
// GetTemplateList 查询(批量查询)邮件模板
// invokeNam    string  否   邮件模板调用名称
// templateType int     否   邮件模板类型: 0(触发), 1(批量)
// templateStat int     否   邮件模板状态: -2(未提交审核), -1(审核不通过), 0(待审核), 1(审核通过)
// start        int     否   查询起始位置, 取值区间 [0-], 默认为 0
// limit        int     否   查询个数, 取值区间 [0-100], 默认为 100

const (
	SEND_CLOUD_TEMPLATE_TYPE_OF_ALL     = -1
	SEND_CLOUD_TEMPLATE_TYPE_OF_TRIGGER = 0
	SEND_CLOUD_TEMPLATE_TYPE_OF_BATCH   = 1
)

const (
	SEND_CLOUD_TEMPLATE_STATE_OF_ALL        = -3
	SEND_CLOUD_TEMPLATE_STATE_OF_NOT_REVIEW = -2   // 未提交审核
	SEND_CLOUD_TEMPLATE_STATE_OF_NOT_PASS   = -1   // 审核不通过
	SEND_CLOUD_TEMPLATE_STATE_OF_REVIEWING  = 0    // 待审核
	SEND_CLOUD_TEMPLATE_STATE_OF_PASSED     = 1    // 审核通过
)

func GetTemplateList(invokeName string, templateType, templateStat, start, limit int) (bool, error, string) {
	params := url.Values{}
	if templateType != SEND_CLOUD_TEMPLATE_TYPE_OF_ALL {
		params.Add("templateType", fmt.Sprintf("%d", templateType))
	}
	if templateStat != SEND_CLOUD_TEMPLATE_STATE_OF_ALL {
		params.Add("templateStat", fmt.Sprintf("%d", templateStat))
	}
	if start >= 0 {
		params.Add("start", fmt.Sprintf("%d", start))
	}
	if limit >= 1 {
		params.Add("limit", fmt.Sprintf("%d", limit))
	}

	return doRequest(SEND_CLOUD_TEMPLATE_LIST_API_URL, params)
}

////////////////////////////////////////////////////////////////////////////////
// AddTemplate 添加模板
// invokeName     string    是   邮件模板调用名称
// name           string    是   邮件模板名称
// html           string    是   html格式内容
// subject        string    是   模板标题
// templateType   int       是   邮件模板类型: 0(触发), 1(批量)
// isSubmitAudit  int       否   是否提交审核: 0(不提交审核), 1(提交审核). 默认为 1
func AddTemplate(invokeName, name, html, subject string, templateType int, isSubmitAudit bool) (bool, error, string) {
	params := url.Values{}
	params.Add("invokeName", invokeName)
	params.Add("name", name)
	params.Add("html", html)
	params.Add("subject", subject)
	params.Add("templateType", fmt.Sprintf("%d", templateType))
	if isSubmitAudit {
		params.Add("isSubmitAudit", "1")
	} else {
		params.Add("isSubmitAudit", "0")
	}
	return doRequest(SEND_CLOUD_TEMPLATE_ADD_API_URL, params)
}

////////////////////////////////////////////////////////////////////////////////
// DeleteTemplate 删除模板
// invokeName     string    是   邮件模板调用名称
func DeleteTemplate(invokeName string) (bool, error, string) {
	params := url.Values{}
	params.Add("invokeName", invokeName)
	return doRequest(SEND_CLOUD_TEMPLATE_DELETE_API_URL, params)
}

////////////////////////////////////////////////////////////////////////////////
// UpdateTemplate 修改模板信息
// invokeName     string   是  邮件模板调用名称
// name           string   否  邮件模板名称
// html           string   否  html格式内容
// subject        string   否  模板标题
// templateType   int      否  邮件模板类型: 0(触发), 1(批量)
// isSubmitAudit  int      否  是否提交审核: 0(不提交审核), 1(提交审核). 默认为 1
func UpdateTemplate(invokeName, name, html, subject string, templateType int, isSubmitAudit bool) (bool, error, string) {
	params := url.Values{}
	params.Add("invokeName", invokeName)
	params.Add("name", name)
	params.Add("html", html)
	params.Add("subject", subject)
	params.Add("templateType", fmt.Sprintf("%d", templateType))
	if isSubmitAudit {
		params.Add("isSubmitAudit", "1")
	} else {
		params.Add("isSubmitAudit", "0")
	}
	return doRequest(SEND_CLOUD_TEMPLATE_UPDATE_API_URL, params)
}

////////////////////////////////////////////////////////////////////////////////
// doRequest 发起网络请求
func doRequest(url string, params url.Values) (bool, error, string) {
	if len(MailApiKey) == 0 || len(MailApiUser) == 0 {
		return false, errors.New("请先配置 api 信息"), ""
	}

	params.Add("apiUser", MailApiUser)
	params.Add("apiKey", MailApiKey)

	var body = bytes.NewBufferString(params.Encode())
	responseHandler, err := http.Post(url, "application/x-www-form-urlencoded", body)
	if err != nil {
		return false, err, ""
	}
	defer responseHandler.Body.Close()

	bodyByte, err := ioutil.ReadAll(responseHandler.Body)
	if err != nil {
		return false, err, string(bodyByte)
	}

	var result map[string]interface{}
	err = json.Unmarshal(bodyByte, &result)
	return (result["result"] == true), err, string(bodyByte)
}