package sendcloud

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
)

const (
	kSendTemplate = "http://api.sendcloud.net/apiv2/mail/sendtemplate"
	kMailTaskInfo = "http://api.sendcloud.net/apiv2/mail/taskinfo"
)

type Client struct {
	apiUser string
	apiKey  string
}

func New(apiUser, apiKey string) *Client {
	var c = &Client{}
	c.apiUser = apiUser
	c.apiKey = apiKey
	return c
}

////////////////////////////////////////////////////////////////////////////////
// SendMailWithTemplate 模板发送
// invokeName   string 是   邮件模板调用名称
// from         string 是   发件人地址
// fromName     string 否   发件人名称
// replyTo      string 否   设置用户默认的回复邮件地址. 如果 replyTo 没有或者为空, 则默认的回复邮件地址为 from
// subject      string *    邮件标题
func (this *Client) SendTemplateMail(invokeName, from, fromName, replyTo, subject string, toList []map[string]string, filename []string) (bool, error, string) {
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
		return false, err, ""
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

	return this.doRequestWithFile(kSendTemplate, params, "attachments", filename)
}

////////////////////////////////////////////////////////////////////////////////
// SendTemplateMailWithAddressList 向邮件地址列表发送邮件
// addressList  string 是   邮件地址列表
// invokeName   string 是   邮件模板调用名称
// from         string 是   发件人地址
// fromName     string 否   发件人名称
// replyTo      string 否   设置用户默认的回复邮件地址. 如果 replyTo 没有或者为空, 则默认的回复邮件地址为 from
// subject      string *    邮件标题
func (this *Client) SendTemplateMailToAddressList(addressList, invokeName, from, fromName, replyTo, subject string, filename []string) (bool, error, string) {
	params := url.Values{
		"to":                 {addressList},
		"from":               {from},
		"fromName":           {fromName},
		"replyTo":            {replyTo},
		"subject":            {subject},
		"templateInvokeName": {invokeName},
		"useAddressList":     {"true"},
	}
	return this.doRequestWithFile(kSendTemplate, params, "attachments", filename)
}

////////////////////////////////////////////////////////////////////////////////
// GetTaskInfo 获取邮件地址列表发送任务信息
// mailListTaskId   int  是  返回的mailListTaskId
func (this *Client) GetTaskInfo(mailListTaskId int) (bool, error, string) {
	params := url.Values{}
	params.Add("maillistTaskId", fmt.Sprintf("%d", mailListTaskId))
	return this.doRequest(kMailTaskInfo, params)
}

////////////////////////////////////////////////////////////////////////////////
// doRequest 发起网络请求
func (this *Client) doRequest(url string, params url.Values) (bool, error, string) {
	if len(this.apiKey) == 0 || len(this.apiUser) == 0 {
		return false, errors.New("请先配置 api 信息"), ""
	}
	params.Add("apiUser", this.apiUser)
	params.Add("apiKey", this.apiKey)

	var body = bytes.NewBufferString(params.Encode())
	rsp, err := http.Post(url, "application/x-www-form-urlencoded", body)
	if rsp.Body != nil {
		defer rsp.Body.Close()
	}
	if err != nil {
		return false, err, ""
	}

	bodyByte, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return false, err, string(bodyByte)
	}

	var result map[string]interface{}
	err = json.Unmarshal(bodyByte, &result)
	return result["result"] == true, err, string(bodyByte)
}

//func doRequestWithFile(url string, params url.Values, fileField string, filenames []string) (bool, error, string) {
//	if len(apiKey) == 0 || len(apiUser) == 0 {
//		return false, errors.New("请先配置 api 信息"), ""
//	}
//	params.Add("apiUser", apiUser)
//	params.Add("apiKey", apiKey)
//
//	var bodyBuf    = bytes.NewBufferString("")
//	var bodyWriter = multipart.NewWriter(bodyBuf)
//	defer bodyWriter.Close()
//
//	for key, value := range params {
//		_ = bodyWriter.WriteField(key, value[0])
//	}
//
//	var readers = make([]io.Reader, 0)
//	readers = append(readers, bodyBuf)
//
//	var fileSize int64
//	for _, filename := range filenames {
//		_, err := bodyWriter.CreateFormFile(fileField, filename)
//		if err != nil {
//			return false, err, ""
//		}
//
//		file, err := os.Open(filename)
//		if err != nil {
//			return false, err, ""
//		}
//
//		readers = append(readers, file)
//
//		fileInfo, err := file.Stat()
//		if err != nil {
//			return false, err, ""
//		}
//
//		fileSize = fileSize + fileInfo.Size()
//	}
//
//	var requestReader io.Reader
//	var boundary = bodyWriter.Boundary()
//
//	var closeBuf = bytes.NewBufferString(fmt.Sprintf("\r\n--%s--\r\n", boundary))
//
//	readers = append(readers, closeBuf)
//	requestReader = io.MultiReader(readers...)
//	request, err := http.NewRequest("POST", url, requestReader)
//	if err != nil {
//		return false, err, ""
//	}
//	request.Header.Add("Content-Type", bodyWriter.FormDataContentType())
//	request.ContentLength = fileSize + int64(bodyBuf.Len()) + int64(closeBuf.Len())
//
//	responseHandler, err := http.DefaultClient.Do(request)
//	if err != nil {
//		return false, err, ""
//	}
//	defer responseHandler.Body.Close()
//
//	bodyByte, err := ioutil.ReadAll(responseHandler.Body)
//	if err != nil {
//		return false, err, string(bodyByte)
//	}
//
//	var result map[string]interface{}
//	err = json.Unmarshal(bodyByte, &result)
//	return (result["result"] == true), err, string(bodyByte)
//}

func (this *Client) doRequestWithFile(url string, params url.Values, fileField string, filenames []string) (bool, error, string) {
	if len(this.apiKey) == 0 || len(this.apiUser) == 0 {
		return false, errors.New("请先配置 api 信息"), ""
	}

	params.Add("apiUser", this.apiUser)
	params.Add("apiKey", this.apiKey)

	var body = &bytes.Buffer{}
	var writer = multipart.NewWriter(body)

	for _, filename := range filenames {
		file, err := os.Open(filename)
		if err != nil {
			return false, err, ""
		}

		fileWriter, err := writer.CreateFormFile(fileField, filename)
		if err != nil {
			return false, err, ""
		}
		_, err = io.Copy(fileWriter, file)
		file.Close()
	}

	for key, value := range params {
		_ = writer.WriteField(key, value[0])
	}

	var err = writer.Close()
	if err != nil {
		return false, err, ""
	}

	request, err := http.NewRequest("POST", url, body)
	request.Header.Set("Content-Type", writer.FormDataContentType())

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
	return result["result"] == true, err, string(bodyByte)

}
