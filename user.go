package sendcloud

import "net/url"

const (
	kUserInfoGet = "http://api.sendcloud.net/apiv2/userinfo/get"
)

////////////////////////////////////////////////////////////////////////////////
// GetUserInfo 查询当前用户的相关信息
func (this *Client) GetUserInfo() (bool, error, string) {
	params := url.Values{}
	return this.doRequest(kUserInfoGet, params)
}
