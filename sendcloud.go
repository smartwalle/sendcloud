package sendcloud

import (
	"bytes"
	"encoding/json"
	"errors"
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
// doRequest 发起网络请求
func (this *Client) doRequest(url string, param url.Values, result interface{}) error {
	if len(this.apiKey) == 0 || len(this.apiUser) == 0 {
		return errors.New("请先配置 api 信息")
	}
	param.Add("apiUser", this.apiUser)
	param.Add("apiKey", this.apiKey)

	var body = bytes.NewBufferString(param.Encode())
	rsp, err := http.Post(url, "application/x-www-form-urlencoded", body)
	if rsp.Body != nil {
		defer rsp.Body.Close()
	}
	if err != nil {
		return err
	}

	bodyByte, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(bodyByte, result)
}

//func doRequestWithFile(url string, param url.Values, fileField string, filenames []string) (bool, error, string) {
//	if len(apiKey) == 0 || len(apiUser) == 0 {
//		return false, errors.New("请先配置 api 信息"), ""
//	}
//	param.Add("apiUser", apiUser)
//	param.Add("apiKey", apiKey)
//
//	var bodyBuf    = bytes.NewBufferString("")
//	var bodyWriter = multipart.NewWriter(bodyBuf)
//	defer bodyWriter.Close()
//
//	for key, value := range param {
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

func (this *Client) doRequestWithFile(url string, param url.Values, fileField string, filenames []string, result interface{}) error {
	if len(this.apiKey) == 0 || len(this.apiUser) == 0 {
		return errors.New("请先配置 api 信息")
	}

	param.Add("apiUser", this.apiUser)
	param.Add("apiKey", this.apiKey)

	var body = &bytes.Buffer{}
	var writer = multipart.NewWriter(body)

	for _, filename := range filenames {
		file, err := os.Open(filename)
		if err != nil {
			return err
		}

		fileWriter, err := writer.CreateFormFile(fileField, filename)
		if err != nil {
			return err
		}
		_, err = io.Copy(fileWriter, file)
		file.Close()
	}

	for key, value := range param {
		_ = writer.WriteField(key, value[0])
	}

	if err := writer.Close(); err != nil {
		return err
	}

	request, err := http.NewRequest("POST", url, body)
	request.Header.Set("Content-Type", writer.FormDataContentType())

	rsp, err := http.DefaultClient.Do(request)
	if err != nil {
		return err
	}
	defer rsp.Body.Close()

	bodyByte, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(bodyByte, result)

}
