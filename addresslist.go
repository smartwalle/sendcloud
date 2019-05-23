package sendcloud

import (
	"errors"
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

type Address struct {
	GMTCreated  string `json:"gmtCreated"`
	GMTtUpdated string `json:"gmtUpdated"`
	Address     string `json:"address"`
	MemberCount int    `json:"memberCount"`
	Description string `json:"description"`
	Name        string `json:"name"`
	ListType    int    `json:"listType"`
}

type Member struct {
	GMTCreated  string `json:"gmtCreated"`
	GMTtUpdated string `json:"gmtUpdated"`
	Member      string `json:"member"`
	Name        string `json:"name"`
	Address     string `json:"address"`
}

// --------------------------------------------------------------------------------
type AddressListRsp struct {
	Result     bool   `json:"result"`
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
	Info       struct {
		DataList []*Address `json:"dataList"`
		Total    int        `json:"total"`
		Count    int        `json:"count"`
	} `json:"info"`
}

// GetAddressList 查询地址列表(批量查询)
//address  list 否  别名地址的列表, 多个用 ; 分隔
//start    int  否  查询起始位置, 取值区间 [0-], 默认为 0
//limit    int  否  查询个数, 取值区间 [0-100], 默认为 100
func (this *Client) GetAddressList(address string, start, limit int) (result *AddressListRsp, err error) {
	param := url.Values{}
	if len(address) > 0 {
		param.Add("address", address)
	}
	if start >= 0 {
		param.Add("start", fmt.Sprintf("%d", start))
	}
	if limit >= 1 {
		param.Add("limit", fmt.Sprintf("%d", limit))
	}

	if err = this.doRequest(kAddressListList, param, &result); err != nil {
		return nil, err
	}

	if result.Result == false {
		return nil, errors.New(fmt.Sprintf("%d-%s", result.StatusCode, result.Message))
	}

	return result, err
}

// --------------------------------------------------------------------------------
type AddAddressListRsp struct {
	Result     bool   `json:"result"`
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
	Info       struct {
		Data *Address `json:"data"`
	} `json:"info"`
}

// AddAddressList 添加地址列表
//address  string   是    别称地址, 使用该别称地址进行调用, 格式为xxx@maillist.sendcloud.org
//name     string   是    列表名称
//desc     string   否    对列表的描述信息
//listType int      否    列表的类型. 0: 普通地址列表, 1: 高级地址列表(需要开通权限才能使用). 默认为0
func (this *Client) AddAddressList(address, name, desc string /*, listType int*/) (result *AddAddressListRsp, err error) {
	param := url.Values{}
	param.Add("address", address)
	param.Add("name", name)
	param.Add("desc", desc)
	//param.Add("listType", fmt.Sprintf("%d", listType))

	if err = this.doRequest(kAddressListAdd, param, &result); err != nil {
		return nil, err
	}

	if result.Result == false {
		return nil, errors.New(fmt.Sprintf("%d-%s", result.StatusCode, result.Message))
	}

	return result, err
}

// --------------------------------------------------------------------------------
type DeleteAddressListRsp struct {
	Result     bool   `json:"result"`
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
	Info       struct {
		Count int `json:"count"`
	} `json:"info"`
}

// DeleteAddressList 删除地址列表
// address  string  是   别称地址, 使用该别称地址进行调用, 格式为xxx@maillist.sendcloud.org
func (this *Client) DeleteAddressList(address string) (result *DeleteAddressListRsp, err error) {
	param := url.Values{}
	param.Add("address", address)

	if err = this.doRequest(kAddressListDelete, param, &result); err != nil {
		return nil, err
	}

	if result.Result == false {
		return nil, errors.New(fmt.Sprintf("%d-%s", result.StatusCode, result.Message))
	}

	return result, err
}

// --------------------------------------------------------------------------------
type UpdateAddressListRsp struct {
	Result     bool   `json:"result"`
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
	Info       struct {
		Count int `json:"count"`
	} `json:"info"`
}

// UpdateAddressList 更新地址列表
// address     string   是    别称地址, 使用该别称地址进行调用, 格式为xxx@maillist.sendcloud.org
// newAddress  string   否    修改后的别称地址
// name        string   否    修改后的列表名称
// desc        string   否    修改后的列表描述信息
func (this *Client) UpdateAddressList(address, newAddress, name, desc string) (result *UpdateAddressListRsp, err error) {
	param := url.Values{}
	param.Add("address", address)
	param.Add("newAddress", newAddress)
	param.Add("name", name)
	param.Add("desc", desc)

	if err = this.doRequest(kAddressListUpdate, param, &result); err != nil {
		return nil, err
	}

	if result.Result == false {
		return nil, errors.New(fmt.Sprintf("%d-%s", result.StatusCode, result.Message))
	}

	return result, err
}

// --------------------------------------------------------------------------------
type AddressMemberListRsp struct {
	Result     bool   `json:"result"`
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
	Info       struct {
		DataList []*Member `json:"dataList"`
		Total    int       `json:"total"`
		Count    int       `json:"count"`
	} `json:"info"`
}

// GetAddressMemberList 获取邮件列表的成员列表
// address string 是 地址列表的别称地址
// start   int    否 查询起始位置, 取值区间 [0-], 默认为 0
// limit   int    否 查询个数, 取值区间 [0-100], 默认为 100
func (this *Client) GetAddressMemberList(address string, start, limit int) (result *AddressMemberListRsp, err error) {
	param := url.Values{}
	param.Add("address", address)
	if start >= 0 {
		param.Add("start", fmt.Sprintf("%d", start))
	}
	if limit >= 1 {
		param.Add("limit", fmt.Sprintf("%d", limit))
	}

	if err = this.doRequest(kAddressMemberList, param, &result); err != nil {
		return nil, err
	}

	if result.Result == false {
		return nil, errors.New(fmt.Sprintf("%d-%s", result.StatusCode, result.Message))
	}

	return result, err
}

// --------------------------------------------------------------------------------
type AddAddressMemberRsp struct {
	Result     bool   `json:"result"`
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
	Info       struct {
		Count int `json:"count"`
	} `json:"info"`
}

// AddAddressMember 向地址列表添加成员
// address string  是   地址列表的别称地址
// members list    是   需要添加成员的地址, 多个地址用 ; 分隔
// names   list	   否   地址成员姓名, 多个地址用 ; 分隔
func (this *Client) AddAddressMember(address string, members, names []string) (result *AddAddressMemberRsp, err error) {
	param := url.Values{}
	param.Add("address", address)
	param.Add("members", strings.Join(members, ";"))
	if len(names) > 0 {
		param.Add("names", strings.Join(names, ";"))
	}

	if err = this.doRequest(kAddressMemberAdd, param, &result); err != nil {
		return nil, err
	}

	if result.Result == false {
		return nil, errors.New(fmt.Sprintf("%d-%s", result.StatusCode, result.Message))
	}

	return result, err
}

// --------------------------------------------------------------------------------
type UpdateAddressMemberRsp struct {
	Result     bool   `json:"result"`
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
	Info       struct {
		Count int `json:"count"`
	} `json:"info"`
}

// UpdateAddressMember 修改邮件列表的成员
// address  string 是  地址列表的别称地址
// members  list   是  需要添加成员的地址, 多个地址用 ; 分隔
// names    list   否  地址成员姓名, 多个地址用 ; 分隔
func (this *Client) UpdateAddressMember(address string, members, names []string) (result *UpdateAddressMemberRsp, err error) {
	param := url.Values{}
	param.Add("address", address)
	param.Add("members", strings.Join(members, ";"))
	if len(names) > 0 {
		param.Add("names", strings.Join(names, ";"))
	}

	if err = this.doRequest(kAddressMemberUpdate, param, &result); err != nil {
		return nil, err
	}

	if result.Result == false {
		return nil, errors.New(fmt.Sprintf("%d-%s", result.StatusCode, result.Message))
	}

	return result, err
}

// --------------------------------------------------------------------------------
type DeleteAddressMemberRsp struct {
	Result     bool   `json:"result"`
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
	Info       struct {
		Count int `json:"count"`
	} `json:"info"`
}

// DeleteAddressMember 从邮件列表删除成员
// address string  是  地址列表的别称地址
// members list    是  需要删除成员的地址, 多个地址用 ; 分隔
func (this *Client) DeleteAddressMember(address string, members []string) (result *DeleteAddressMemberRsp, err error) {
	param := url.Values{}
	param.Add("address", address)
	param.Add("members", strings.Join(members, ";"))

	if err = this.doRequest(kAddressMemberDelete, param, &result); err != nil {
		return nil, err
	}

	if result.Result == false {
		return nil, errors.New(fmt.Sprintf("%d-%s", result.StatusCode, result.Message))
	}

	return result, err
}
