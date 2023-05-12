package sendcloud

type TemplateType int

const (
	TemplateTypeAll     = -1
	TemplateTypeTrigger = 0
	TemplateTypeBatch   = 1
)

type TemplateState int

const (
	TemplateStateAll       = -3
	TemplateStateNotReview = -2 // 未提交审核
	TemplateStateNotPass   = -1 // 审核不通过
	TemplateStateReviewing = 0  // 待审核
	TemplateStatePassed    = 1  // 审核通过
)

type Template struct {
	GMTCreated     string       `json:"gmtCreated"`
	GMTtUpdated    string       `json:"gmtUpdated"`
	Name           string       `json:"name"`
	InvokeName     string       `json:"invokeName"`
	TemplateType   TemplateType `json:"templateType"`
	HTML           string       `json:"html"`
	Subject        string       `json:"subject"`
	ContentSummary string       `json:"contentSummary"`
}

type TemplateListRsp struct {
	Result     bool   `json:"result"`
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
	Info       struct {
		DataList []*Template `json:"dataList"`
		Total    int         `json:"total"`
		Count    int         `json:"count"`
	} `json:"info"`
}

type GetTemplateRsp struct {
	Result     bool   `json:"result"`
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
	Info       struct {
		Data *Template `json:"data"`
	} `json:"info"`
}

type AddTemplateRsp struct {
	Result     bool   `json:"result"`
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
	Info       struct {
		Data *Template `json:"data"`
	} `json:"info"`
}

type DeleteTemplateRsp struct {
	Result     bool   `json:"result"`
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
	Info       struct {
		Count int `json:"count"`
	} `json:"info"`
}

type UpdateTemplateRsp struct {
	Result     bool   `json:"result"`
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
	Info       struct {
		Count int `json:"count"`
	} `json:"info"`
}
