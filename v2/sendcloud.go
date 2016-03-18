package v2

import (
	"net/url"
	"bytes"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"mime/multipart"
	"io"
)

const (
	SEND_CLOUD_SEND_TEMPLATE_API_URL   = "http://api.sendcloud.net/apiv2/mail/sendtemplate"
	SEND_CLOUD_MAIL_TASK_INFO_API_URL  = "http://api.sendcloud.net/apiv2/mail/taskinfo"
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
// invokeName   string 是   邮件模板调用名称
// from         string 是   发件人地址
// fromName     string 否   发件人名称
// replyTo      string 否   设置用户默认的回复邮件地址. 如果 replyTo 没有或者为空, 则默认的回复邮件地址为 from
// subject      string *    邮件标题
func SendTemplateMail(invokeName, from, fromName, replyTo, subject string, toList []map[string]string, filename string) (bool, error, string) {
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
		"templateInvokeName": {invokeName},
		"xsmtpapi":    {substitutionVars},
	}

	return doRequestWithFile(SEND_CLOUD_SEND_TEMPLATE_API_URL, params, "attachments", filename)
}

////////////////////////////////////////////////////////////////////////////////
// SendTemplateMailWithAddressList 向邮件地址列表发送邮件
// addressList  string 是   邮件地址列表
// invokeName   string 是   邮件模板调用名称
// from         string 是   发件人地址
// fromName     string 否   发件人名称
// replyTo      string 否   设置用户默认的回复邮件地址. 如果 replyTo 没有或者为空, 则默认的回复邮件地址为 from
// subject      string *    邮件标题
func SendTemplateMailToAddressList(addressList, invokeName, from, fromName, replyTo, subject, filename string) (bool, error, string) {
	params := url.Values{
		"to":       {addressList},
		"from":     {from},
		"fromName": {fromName},
		"replyTo":  {replyTo},
		"subject":  {subject},
		"templateInvokeName": {invokeName},
		"useAddressList": {"true"},
	}
	return doRequestWithFile(SEND_CLOUD_SEND_TEMPLATE_API_URL, params, "attachments", filename)
}

////////////////////////////////////////////////////////////////////////////////
// GetTaskInfo 获取邮件地址列表发送任务信息
// maillistTaskId   int  是  返回的maillistTaskId
func GetTaskInfo(maillistTaskId int) (bool, error, string) {
	params := url.Values{}
	params.Add("maillistTaskId", fmt.Sprintf("%d", maillistTaskId))
	return doRequest(SEND_CLOUD_MAIL_TASK_INFO_API_URL, params)
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

func doRequestWithFile(url string, params url.Values, fileField, filename string) (bool, error, string) {
	if len(MailApiKey) == 0 || len(MailApiUser) == 0 {
		return false, errors.New("请先配置 api 信息"), ""
	}
	params.Add("apiUser", MailApiUser)
	params.Add("apiKey", MailApiKey)

	var bodyBuf    = bytes.NewBufferString("")
	var bodyWriter = multipart.NewWriter(bodyBuf)

	for key, val := range params {
		_ = bodyWriter.WriteField(key, val[0])
	}

	_, err := bodyWriter.CreateFormFile(fileField, filename)
	if err != nil {
		return false, err, ""
	}

	var boundary = bodyWriter.Boundary()
	var closeBuf = bytes.NewBufferString(fmt.Sprintf("\r\n--%s--\r\n", boundary))

	var requestReader io.Reader
	var fileSize      int64
	if filename != "" {
		file, err := os.Open(filename)
		defer file.Close()
		if err != nil {
			return false, err, ""
		}
		requestReader = io.MultiReader(bodyBuf, file, closeBuf)

		fileInfo, err := file.Stat()
		if err != nil {
			return false, err, ""
		}

		fileSize = fileInfo.Size()
	} else {
		requestReader = io.MultiReader(bodyBuf, closeBuf)
	}

	request, err := http.NewRequest("POST", url, requestReader)
	if err != nil {
		return false, err, ""
	}
	request.Header.Add("Content-Type", "multipart/form-data; boundary="+boundary)
	request.ContentLength = fileSize + int64(bodyBuf.Len()) + int64(closeBuf.Len())

	responseHandler, err := http.DefaultClient.Do(request)
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