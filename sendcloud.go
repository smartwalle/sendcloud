package sendcloud

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type Client struct {
	apiUser string
	apiKey  string
	Client  *http.Client
}

func New(apiUser, apiKey string) *Client {
	var c = &Client{}
	c.apiUser = apiUser
	c.apiKey = apiKey
	c.Client = http.DefaultClient
	return c
}

// doRequest 发起网络请求
func (this *Client) doRequest(url string, values url.Values, result interface{}) error {
	if len(this.apiKey) == 0 || len(this.apiUser) == 0 {
		return errors.New("请先配置 api 信息")
	}
	values.Add("apiUser", this.apiUser)
	values.Add("apiKey", this.apiKey)

	rsp, err := http.Post(url, "application/x-www-form-urlencoded", strings.NewReader(values.Encode()))
	if rsp != nil {
		defer rsp.Body.Close()
	}
	if err != nil {
		return err
	}

	bodyBytes, err := io.ReadAll(rsp.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(bodyBytes, result)
}

func (this *Client) doRequestWithFile(url string, values url.Values, fieldname string, filenames []string, result interface{}) error {
	if len(this.apiKey) == 0 || len(this.apiUser) == 0 {
		return errors.New("请先配置 api 信息")
	}

	values.Add("apiUser", this.apiUser)
	values.Add("apiKey", this.apiKey)

	var bodyBuffer = &bytes.Buffer{}
	var bodyWriter = multipart.NewWriter(bodyBuffer)

	for _, filename := range filenames {
		fileContent, err := os.ReadFile(filename)
		if err != nil {
			return err
		}
		fileWriter, err := bodyWriter.CreateFormFile(fieldname, filename)
		if err != nil {
			return err
		}
		if _, err = fileWriter.Write(fileContent); err != nil {
			return err
		}
	}

	for key, value := range values {
		for _, v := range value {
			_ = bodyWriter.WriteField(key, v)
		}
	}

	if err := bodyWriter.Close(); err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, bodyBuffer)
	req.Header.Set("Content-Type", bodyWriter.FormDataContentType())

	rsp, err := this.Client.Do(req)
	if err != nil {
		return err
	}
	defer rsp.Body.Close()

	bodyBytes, err := io.ReadAll(rsp.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(bodyBytes, result)
}
