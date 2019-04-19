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

////////////////////////////////////////////////////////////////////////////////
// GetAddressList 查询地址列表(批量查询)
//address  list 否  别名地址的列表, 多个用 ; 分隔
//start    int  否  查询起始位置, 取值区间 [0-], 默认为 0
//limit    int  否  查询个数, 取值区间 [0-100], 默认为 100
func (this *Client) GetAddressList(address string, start, limit int) (bool, error, string) {
	params := url.Values{}
	if len(address) > 0 {
		params.Add("address", address)
	}
	if start >= 0 {
		params.Add("start", fmt.Sprintf("%d", start))
	}
	if limit >= 1 {
		params.Add("limit", fmt.Sprintf("%d", limit))
	}
	return this.doRequest(kAddressListList, params)
}

////////////////////////////////////////////////////////////////////////////////
// AddAddressList 添加地址列表
//address  string   是    别称地址, 使用该别称地址进行调用, 格式为xxx@maillist.sendcloud.org
//name     string   是    列表名称
//desc     string   否    对列表的描述信息
//listType int      否    列表的类型. 0: 普通地址列表, 1: 高级地址列表(需要开通权限才能使用). 默认为0
func (this *Client) AddAddressList(address, name, desc string /*, listType int*/) (bool, error, string) {
	params := url.Values{}
	params.Add("address", address)
	params.Add("name", name)
	params.Add("desc", desc)
	//params.Add("listType", fmt.Sprintf("%d", listType))
	return this.doRequest(kAddressListAdd, params)
}

////////////////////////////////////////////////////////////////////////////////
// DeleteAddressList 删除地址列表
// address  string  是   别称地址, 使用该别称地址进行调用, 格式为xxx@maillist.sendcloud.org
func (this *Client) DeleteAddressList(address string) (bool, error, string) {
	params := url.Values{}
	params.Add("address", address)
	return this.doRequest(kAddressListDelete, params)
}

////////////////////////////////////////////////////////////////////////////////
// UpdateAddressList 更新地址列表
// address     string   是    别称地址, 使用该别称地址进行调用, 格式为xxx@maillist.sendcloud.org
// newAddress  string   否    修改后的别称地址
// name        string   否    修改后的列表名称
// desc        string   否    修改后的列表描述信息
func (this *Client) UpdateAddressList(address, newAddress, name, desc string) (bool, error, string) {
	params := url.Values{}
	params.Add("address", address)
	params.Add("newAddress", newAddress)
	params.Add("name", name)
	params.Add("desc", desc)
	return this.doRequest(kAddressListUpdate, params)
}

////////////////////////////////////////////////////////////////////////////////
// GetAddressMemberList 获取邮件列表的成员列表
// address string 是 地址列表的别称地址
// start   int    否 查询起始位置, 取值区间 [0-], 默认为 0
// limit   int    否 查询个数, 取值区间 [0-100], 默认为 100
func (this *Client) GetAddressMemberList(address string, start, limit int) (bool, error, string) {
	params := url.Values{}
	params.Add("address", address)
	if start >= 0 {
		params.Add("start", fmt.Sprintf("%d", start))
	}
	if limit >= 1 {
		params.Add("limit", fmt.Sprintf("%d", limit))
	}
	return this.doRequest(kAddressMemberList, params)
}

////////////////////////////////////////////////////////////////////////////////
// AddAddressMember 向地址列表添加成员
// address string  是   地址列表的别称地址
// members list    是   需要添加成员的地址, 多个地址用 ; 分隔
// names   list	   否   地址成员姓名, 多个地址用 ; 分隔
func (this *Client) AddAddressMember(address string, members, names []string) (bool, error, string) {
	params := url.Values{}
	params.Add("address", address)
	params.Add("members", strings.Join(members, ";"))
	if names != nil && len(names) > 0 {
		params.Add("names", strings.Join(names, ";"))
	}
	return this.doRequest(kAddressMemberAdd, params)
}

////////////////////////////////////////////////////////////////////////////////
// UpdateAddressMember 修改邮件列表的成员
// address  string 是  地址列表的别称地址
// members  list   是  需要添加成员的地址, 多个地址用 ; 分隔
// names    list   否  地址成员姓名, 多个地址用 ; 分隔
func (this *Client) UpdateAddressMember(address string, members, names []string) (bool, error, string) {
	params := url.Values{}
	params.Add("address", address)
	params.Add("members", strings.Join(members, ";"))
	if names != nil && len(names) > 0 {
		params.Add("names", strings.Join(names, ";"))
	}
	return this.doRequest(kAddressMemberUpdate, params)
}

////////////////////////////////////////////////////////////////////////////////
// DeleteAddressMember 从邮件列表删除成员
// address string  是  地址列表的别称地址
// members list    是  需要删除成员的地址, 多个地址用 ; 分隔
func (this *Client) DeleteAddressMember(address string, members []string) (bool, error, string) {
	params := url.Values{}
	params.Add("address", address)
	params.Add("members", strings.Join(members, ";"))
	return this.doRequest(kAddressMemberDelete, params)
}
