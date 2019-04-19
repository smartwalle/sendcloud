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

////////////////////////////////////////////////////////////////////////////////
// GetTemplateList 查询(批量查询)邮件模板
// invokeNam    string  否   邮件模板调用名称
// templateType int     否   邮件模板类型: 0(触发), 1(批量)
// templateStat int     否   邮件模板状态: -2(未提交审核), -1(审核不通过), 0(待审核), 1(审核通过)
// start        int     否   查询起始位置, 取值区间 [0-], 默认为 0
// limit        int     否   查询个数, 取值区间 [0-100], 默认为 100

const (
	TEMPLATE_TYPE_OF_ALL     = -1
	TEMPLATE_TYPE_OF_TRIGGER = 0
	TEMPLATE_TYPE_OF_BATCH   = 1
)

const (
	TEMPLATE_STATE_OF_ALL        = -3
	TEMPLATE_STATE_OF_NOT_REVIEW = -2 // 未提交审核
	TEMPLATE_STATE_OF_NOT_PASS   = -1 // 审核不通过
	TEMPLATE_STATE_OF_REVIEWING  = 0  // 待审核
	TEMPLATE_STATE_OF_PASSED     = 1  // 审核通过
)

func (this *Client) GetTemplateList(invokeName string, templateType, templateStat, start, limit int) (bool, error, string) {
	params := url.Values{}
	if templateType != TEMPLATE_TYPE_OF_ALL {
		params.Add("templateType", fmt.Sprintf("%d", templateType))
	}
	if templateStat != TEMPLATE_STATE_OF_ALL {
		params.Add("templateStat", fmt.Sprintf("%d", templateStat))
	}
	if start >= 0 {
		params.Add("start", fmt.Sprintf("%d", start))
	}
	if limit >= 1 {
		params.Add("limit", fmt.Sprintf("%d", limit))
	}

	if len(invokeName) > 0 {
		params.Add("invokeName", invokeName)
	}

	return this.doRequest(kTemplateList, params)
}

////////////////////////////////////////////////////////////////////////////////
// GetTemplate 查询模板的详细信息
// invokeName   string   是    邮件模板调用名称
func (this *Client) GetTemplate(invokeName string) (bool, error, string) {
	params := url.Values{}
	params.Add("invokeName", invokeName)
	return this.doRequest(kTemplateGet, params)
}

////////////////////////////////////////////////////////////////////////////////
// AddTemplate 添加模板
// invokeName     string    是   邮件模板调用名称
// name           string    是   邮件模板名称
// html           string    是   html格式内容
// subject        string    是   模板标题
// templateType   int       是   邮件模板类型: 0(触发), 1(批量)
// isSubmitAudit  int       否   是否提交审核: 0(不提交审核), 1(提交审核). 默认为 1
func (this *Client) AddTemplate(invokeName, name, html, subject string, templateType int, isSubmitAudit bool) (bool, error, string) {
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
	return this.doRequest(kTemplateAdd, params)
}

////////////////////////////////////////////////////////////////////////////////
// DeleteTemplate 删除模板
// invokeName     string    是   邮件模板调用名称
func (this *Client) DeleteTemplate(invokeName string) (bool, error, string) {
	params := url.Values{}
	params.Add("invokeName", invokeName)
	return this.doRequest(kTemplateDelete, params)
}

////////////////////////////////////////////////////////////////////////////////
// UpdateTemplate 修改模板信息
// invokeName     string   是  邮件模板调用名称
// name           string   否  邮件模板名称
// html           string   否  html格式内容
// subject        string   否  模板标题
// templateType   int      否  邮件模板类型: 0(触发), 1(批量)
// isSubmitAudit  int      否  是否提交审核: 0(不提交审核), 1(提交审核). 默认为 1
func (this *Client) UpdateTemplate(invokeName, name, html, subject string, templateType int, isSubmitAudit bool) (bool, error, string) {
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
	return this.doRequest(kTemplateUpdate, params)
}
