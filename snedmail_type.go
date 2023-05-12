package sendcloud

type TaskInfo struct {
	GMTCreated     string `json:"gmtCreated"`
	GMTtUpdated    string `json:"gmtUpdated"`
	MailListTaskId int    `json:"maillistTaskId"`
	ApiUser        string `json:"apiUser"`
	AddressList    string `json:"addressList"`
	MemberCount    int    `json:"memberCount"`
	Status         string `json:"status"`
}

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

type GetTaskInfoRsp struct {
	Result     bool   `json:"result"`
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
	Info       struct {
		Data *TaskInfo `json:"data"`
	} `json:"info"`
}
