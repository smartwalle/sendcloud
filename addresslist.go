package sendcloud

import (
	"fmt"
	"net/url"
	"strings"
)

const (
	kAddressListList     = "http://api.sendcloud.net/apiv2/addresslist/list"
	kAddressListAdd      = "http://api.sendcloud.net/apiv2/addresslist/add"
	kAddressListDelete   = "http://api.sendcloud.net/apiv2/addresslist/delete"
	kAddressListUpdate   = "http://api.sendcloud.net/apiv2/addresslist/update"
	kAddressMemberList   = "http://api.sendcloud.net/apiv2/addressmember/list"
	kAddressMemberAdd    = "http://api.sendcloud.net/apiv2/addressmember/add"
	kAddressMemberUpdate = "http://api.sendcloud.net/apiv2/addressmember/update"
	kAddressMemberDelete = "http://api.sendcloud.net/apiv2/addressmember/delete"
)

// GetAddressList 查询地址列表(批量查询) https://www.sendcloud.net/doc/email_v2/list_do/#_1
// address  list 否  别名地址的列表, 多个用 ; 分隔
// start    int  否  查询起始位置, 取值区间 [0-], 默认为 0
// limit    int  否  查询个数, 取值区间 [0-100], 默认为 100
func (this *Client) GetAddressList(address string, start, limit int) (result *AddressListRsp, err error) {
	var values = url.Values{}
	if len(address) > 0 {
		values.Add("address", address)
	}
	if start >= 0 {
		values.Add("start", fmt.Sprintf("%d", start))
	}
	if limit >= 1 {
		values.Add("limit", fmt.Sprintf("%d", limit))
	}

	if err = this.doRequest(kAddressListList, values, &result); err != nil {
		return nil, err
	}

	if result.Result == false {
		return nil, fmt.Errorf("%d-%s", result.StatusCode, result.Message)
	}

	return result, err
}

// AddAddressList 添加地址列表 https://www.sendcloud.net/doc/email_v2/list_do/#_2
// address  string   是    别称地址, 使用该别称地址进行调用, 格式为xxx@maillist.sendcloud.org
// name     string   是    列表名称
// desc     string   否    对列表的描述信息
func (this *Client) AddAddressList(address, name, desc string) (result *AddAddressListRsp, err error) {
	var values = url.Values{}
	values.Add("address", address)
	values.Add("name", name)
	values.Add("desc", desc)

	if err = this.doRequest(kAddressListAdd, values, &result); err != nil {
		return nil, err
	}

	if result.Result == false {
		return nil, fmt.Errorf("%d-%s", result.StatusCode, result.Message)
	}

	return result, err
}

// DeleteAddressList 删除地址列表 https://www.sendcloud.net/doc/email_v2/list_do/#_3
// address  string  是   别称地址, 使用该别称地址进行调用, 格式为xxx@maillist.sendcloud.org
func (this *Client) DeleteAddressList(address string) (result *DeleteAddressListRsp, err error) {
	var values = url.Values{}
	values.Add("address", address)

	if err = this.doRequest(kAddressListDelete, values, &result); err != nil {
		return nil, err
	}

	if result.Result == false {
		return nil, fmt.Errorf("%d-%s", result.StatusCode, result.Message)
	}

	return result, err
}

// UpdateAddressList 更新地址列表 https://www.sendcloud.net/doc/email_v2/list_do/#_4
// address     string   是    别称地址, 使用该别称地址进行调用, 格式为xxx@maillist.sendcloud.org
// newAddress  string   否    修改后的别称地址
// name        string   否    修改后的列表名称
// desc        string   否    修改后的列表描述信息
func (this *Client) UpdateAddressList(address, newAddress, name, desc string) (result *UpdateAddressListRsp, err error) {
	var values = url.Values{}
	values.Add("address", address)
	values.Add("newAddress", newAddress)
	values.Add("name", name)
	values.Add("desc", desc)

	if err = this.doRequest(kAddressListUpdate, values, &result); err != nil {
		return nil, err
	}

	if result.Result == false {
		return nil, fmt.Errorf("%d-%s", result.StatusCode, result.Message)
	}

	return result, err
}

// GetAddressMemberList 获取邮件列表的成员列表 https://www.sendcloud.net/doc/email_v2/list_do/#_5
// address string 是 地址列表的别称地址
// start   int    否 查询起始位置, 取值区间 [0-], 默认为 0
// limit   int    否 查询个数, 取值区间 [0-100], 默认为 100
func (this *Client) GetAddressMemberList(address string, start, limit int) (result *AddressMemberListRsp, err error) {
	var values = url.Values{}
	values.Add("address", address)
	if start >= 0 {
		values.Add("start", fmt.Sprintf("%d", start))
	}
	if limit >= 1 {
		values.Add("limit", fmt.Sprintf("%d", limit))
	}

	if err = this.doRequest(kAddressMemberList, values, &result); err != nil {
		return nil, err
	}

	if result.Result == false {
		return nil, fmt.Errorf("%d-%s", result.StatusCode, result.Message)
	}

	return result, err
}

// AddAddressMember 向地址列表添加成员 https://www.sendcloud.net/doc/email_v2/list_do/#_7
// address string  是   地址列表的别称地址
// members list    是   需要添加成员的地址, 多个地址用 ; 分隔
// names   list	   否   地址成员姓名, 多个地址用 ; 分隔
func (this *Client) AddAddressMember(address string, members, names []string) (result *AddAddressMemberRsp, err error) {
	var values = url.Values{}
	values.Add("address", address)
	values.Add("members", strings.Join(members, ";"))
	if len(names) > 0 {
		values.Add("names", strings.Join(names, ";"))
	}

	if err = this.doRequest(kAddressMemberAdd, values, &result); err != nil {
		return nil, err
	}

	if result.Result == false {
		return nil, fmt.Errorf("%d-%s", result.StatusCode, result.Message)
	}

	return result, err
}

// UpdateAddressMember 修改邮件列表的成员 https://www.sendcloud.net/doc/email_v2/list_do/#_8
// address  string 是  地址列表的别称地址
// members  list   是  需要添加成员的地址, 多个地址用 ; 分隔
// names    list   否  地址成员姓名, 多个地址用 ; 分隔
func (this *Client) UpdateAddressMember(address string, members, names []string) (result *UpdateAddressMemberRsp, err error) {
	var values = url.Values{}
	values.Add("address", address)
	values.Add("members", strings.Join(members, ";"))
	if len(names) > 0 {
		values.Add("names", strings.Join(names, ";"))
	}

	if err = this.doRequest(kAddressMemberUpdate, values, &result); err != nil {
		return nil, err
	}

	if result.Result == false {
		return nil, fmt.Errorf("%d-%s", result.StatusCode, result.Message)
	}

	return result, err
}

// DeleteAddressMember 从邮件列表删除成员 https://www.sendcloud.net/doc/email_v2/list_do/#_9
// address string  是  地址列表的别称地址
// members list    是  需要删除成员的地址, 多个地址用 ; 分隔
func (this *Client) DeleteAddressMember(address string, members []string) (result *DeleteAddressMemberRsp, err error) {
	var values = url.Values{}
	values.Add("address", address)
	values.Add("members", strings.Join(members, ";"))

	if err = this.doRequest(kAddressMemberDelete, values, &result); err != nil {
		return nil, err
	}

	if result.Result == false {
		return nil, fmt.Errorf("%d-%s", result.StatusCode, result.Message)
	}

	return result, err
}
