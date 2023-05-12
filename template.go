package sendcloud

import (
	"fmt"
	"net/url"
)

const (
	kTemplateList   = "http://api.sendcloud.net/apiv2/template/list"
	kTemplateGet    = "http://api.sendcloud.net/apiv2/template/get"
	kTemplateAdd    = "http://api.sendcloud.net/apiv2/template/add"
	kTemplateDelete = "http://api.sendcloud.net/apiv2/template/delete"
	kTemplateUpdate = "http://api.sendcloud.net/apiv2/template/update"
)

// GetTemplateList 查询(批量查询)邮件模板 https://www.sendcloud.net/doc/email_v2/template_do/#_1
// invokeNam    string  否   邮件模板调用名称
// templateType int     否   邮件模板类型: 0(触发), 1(批量)
// templateStat int     否   邮件模板状态: -2(未提交审核), -1(审核不通过), 0(待审核), 1(审核通过)
// start        int     否   查询起始位置, 取值区间 [0-], 默认为 0
// limit        int     否   查询个数, 取值区间 [0-100], 默认为 100
func (this *Client) GetTemplateList(invokeName string, templateType TemplateType, start, limit int) (result *TemplateListRsp, err error) {
	var values = url.Values{}
	if templateType != TemplateTypeAll {
		values.Add("templateType", fmt.Sprintf("%d", templateType))
	}
	if start >= 0 {
		values.Add("start", fmt.Sprintf("%d", start))
	}
	if limit >= 1 {
		values.Add("limit", fmt.Sprintf("%d", limit))
	}

	if len(invokeName) > 0 {
		values.Add("invokeName", invokeName)
	}

	if err = this.doRequest(kTemplateList, values, &result); err != nil {
		return nil, err
	}

	if result.Result == false {
		return nil, fmt.Errorf("%d-%s", result.StatusCode, result.Message)
	}

	return result, err
}

// GetTemplate 查询模板的详细信息 https://www.sendcloud.net/doc/email_v2/template_do/#_2
// invokeName   string   是    邮件模板调用名称
func (this *Client) GetTemplate(invokeName string) (result *GetTemplateRsp, err error) {
	var values = url.Values{}
	values.Add("invokeName", invokeName)

	if err = this.doRequest(kTemplateGet, values, &result); err != nil {
		return nil, err
	}

	if result.Result == false {
		return nil, fmt.Errorf("%d-%s", result.StatusCode, result.Message)
	}

	return result, err
}

// AddTemplate 添加模板 https://www.sendcloud.net/doc/email_v2/template_do/#_3
// invokeName     string    是   邮件模板调用名称
// name           string    是   邮件模板名称
// html           string    是   html格式内容
// subject        string    是   模板标题
// templateType   int       是   邮件模板类型: 0(触发), 1(批量)
// isSubmitAudit  int       否   是否提交审核: 0(不提交审核), 1(提交审核). 默认为 1
func (this *Client) AddTemplate(invokeName, name, html, subject string, templateType TemplateType) (result *AddTemplateRsp, err error) {
	var values = url.Values{}
	values.Add("invokeName", invokeName)
	values.Add("name", name)
	values.Add("html", html)
	values.Add("subject", subject)
	values.Add("templateType", fmt.Sprintf("%d", templateType))

	if err = this.doRequest(kTemplateAdd, values, &result); err != nil {
		return nil, err
	}

	if result.Result == false {
		return nil, fmt.Errorf("%d-%s", result.StatusCode, result.Message)
	}

	return result, err
}

// DeleteTemplate 删除模板 https://www.sendcloud.net/doc/email_v2/template_do/#_4
// invokeName     string    是   邮件模板调用名称
func (this *Client) DeleteTemplate(invokeName string) (result *DeleteTemplateRsp, err error) {
	var values = url.Values{}
	values.Add("invokeName", invokeName)

	if err = this.doRequest(kTemplateDelete, values, &result); err != nil {
		return nil, err
	}

	if result.Result == false {
		return nil, fmt.Errorf("%d-%s", result.StatusCode, result.Message)
	}

	return result, err
}

// UpdateTemplate 修改模板信息 https://www.sendcloud.net/doc/email_v2/template_do/#_5
// invokeName     string   是  邮件模板调用名称
// name           string   否  邮件模板名称
// html           string   否  html格式内容
// subject        string   否  模板标题
// templateType   int      否  邮件模板类型: 0(触发), 1(批量)
// isSubmitAudit  int      否  是否提交审核: 0(不提交审核), 1(提交审核). 默认为 1
func (this *Client) UpdateTemplate(invokeName, name, html, subject string, templateType TemplateType) (result *UpdateTemplateRsp, err error) {
	var values = url.Values{}
	values.Add("invokeName", invokeName)
	values.Add("name", name)
	values.Add("html", html)
	values.Add("subject", subject)
	values.Add("templateType", fmt.Sprintf("%d", templateType))

	if err = this.doRequest(kTemplateUpdate, values, &result); err != nil {
		return nil, err
	}

	if result.Result == false {
		return nil, fmt.Errorf("%d-%s", result.StatusCode, result.Message)
	}

	return result, err
}
