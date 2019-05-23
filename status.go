package sendcloud

import (
	"errors"
	"fmt"
	"net/url"
)

const (
	kEmailStatus = "http://api.sendcloud.net/apiv2/data/emailStatus"
)

type VoInfo struct {
	Status        string `json:status`
	EmailId       string `json:emailId`
	ApiUser       string `json:apiUser`
	Recipients    string `json:recipients`
	RequestTime   string `json:requestTime`
	ModifiedTime  string `json:modifiedTime`
	SendLog       string `json:sendLog`
	TaskName      string `json:taskName`
	MailingStatus string `json:mailingStatus`
	SubStatus     string `json:subStatus`
	SoftStatus    string `json:softStatus`
	TimeStr       string `json:timeStr`
	Event         string `json:event`
	Receiver      string `json:receiver`
	Message       string `json:message`
	Email         string `json:email`
	Name          string `json:name`
	Phone         string `json:phone`
	SubStatusDesc string `json:subStatusDesc`
}

// --------------------------------------------------------------------------------
type EmailStatusListRsp struct {
	Result     bool   `json:"result"`
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
	Info       struct {
		Total      int       `json:"total,string"`
		VoListSize int       `json:"voListSize"`
		VoList     []*VoInfo `json:"voList"`
	} `json:"info"`
}

// GetEmailStatusList 投递回应
// 目前只提供了几个常用参数，其它参数信息请参考 http://www.sendcloud.net/doc/email_v2/deliver_response/#_1
//start    int  否  查询起始位置, 取值区间 [0-], 默认为 0
//limit    int  否  查询个数, 取值区间 [0-100], 默认为 100
func (this *Client) GetEmailStatusList(days int, startDate, endDate string, start, limit int) (result *EmailStatusListRsp, err error) {
	params := url.Values{}

	if days > 0 {
		params.Add("days", fmt.Sprintf("%d", days))
	}
	if startDate != "" {
		params.Set("startDate", startDate)
	}
	if endDate != "" {
		params.Set("endDate", endDate)
	}
	if start >= 0 {
		params.Add("start", fmt.Sprintf("%d", start))
	}
	if limit >= 1 {
		params.Add("limit", fmt.Sprintf("%d", limit))
	}

	if err = this.doRequest(kEmailStatus, params, &result); err != nil {
		return nil, err
	}

	if result.Result == false {
		return nil, errors.New(fmt.Sprintf("%d-%s", result.StatusCode, result.Message))
	}

	return result, err
}
