package sendcloud

import (
	"fmt"
	"net/url"
)

const (
	kUserInfoGet = "http://api.sendcloud.net/apiv2/userinfo/get"
)

// GetUserInfo 查询当前用户的相关信息 https://www.sendcloud.net/doc/email_v2/user_info/#_1
func (this *Client) GetUserInfo() (result *UserInfoRsp, err error) {
	var values = url.Values{}

	if err = this.doRequest(kUserInfoGet, values, &result); err != nil {
		return nil, err
	}

	if result.Result == false {
		return nil, fmt.Errorf("%d-%s", result.StatusCode, result.Message)
	}

	return result, err
}
