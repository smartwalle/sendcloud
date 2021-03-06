package sendcloud

import (
	"errors"
	"fmt"
	"net/url"
)

const (
	kUserInfoGet = "http://api.sendcloud.net/apiv2/userinfo/get"
)

type UserInfo struct {
	AvailableBalance   float64 `json:"avaliableBalance"`
	Balance            float64 `json:"balance"`
	Phone              string  `json:"phone"`
	Quota              int     `json:"quota"`
	RegTime            string  `json:"regTime"`
	Email              string  `json:"email"`
	Reputation         float64 `json:"reputation"`
	WebsiteAuthStatus  bool    `json:"websiteAuthStatus"`
	AccountType        string  `json:"accountType"`
	UserName           string  `json:"userName"`
	BusinessAuthStatus bool    `json:"businessAuthStatus"`
	TodayUsedQuota     int     `json:"todayUsedQuota"`
}

// --------------------------------------------------------------------------------
type UserInfoRsp struct {
	Result     bool      `json:"result"`
	StatusCode int       `json:"statusCode"`
	Message    string    `json:"message"`
	Info       *UserInfo `json:"info"`
}

// GetUserInfo 查询当前用户的相关信息
func (this *Client) GetUserInfo() (result *UserInfoRsp, err error) {
	param := url.Values{}

	if err = this.doRequest(kUserInfoGet, param, &result); err != nil {
		return nil, err
	}

	if result.Result == false {
		return nil, errors.New(fmt.Sprintf("%d-%s", result.StatusCode, result.Message))
	}

	return result, err
}
