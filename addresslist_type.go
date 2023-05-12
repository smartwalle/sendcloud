package sendcloud

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

type AddAddressListRsp struct {
	Result     bool   `json:"result"`
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
	Info       struct {
		Data *Address `json:"data"`
	} `json:"info"`
}

type DeleteAddressListRsp struct {
	Result     bool   `json:"result"`
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
	Info       struct {
		Count int `json:"count"`
	} `json:"info"`
}

type UpdateAddressListRsp struct {
	Result     bool   `json:"result"`
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
	Info       struct {
		Count int `json:"count"`
	} `json:"info"`
}

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

type AddAddressMemberRsp struct {
	Result     bool   `json:"result"`
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
	Info       struct {
		Count int `json:"count"`
	} `json:"info"`
}

type UpdateAddressMemberRsp struct {
	Result     bool   `json:"result"`
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
	Info       struct {
		Count int `json:"count"`
	} `json:"info"`
}

type DeleteAddressMemberRsp struct {
	Result     bool   `json:"result"`
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
	Info       struct {
		Count int `json:"count"`
	} `json:"info"`
}
