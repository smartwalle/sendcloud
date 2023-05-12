package sendcloud

type VoInfo struct {
	Status        string `json:"status"`
	EmailId       string `json:"emailId"`
	ApiUser       string `json:"apiUser"`
	Recipients    string `json:"recipients"`
	RequestTime   string `json:"requestTime"`
	ModifiedTime  string `json:"modifiedTime"`
	SendLog       string `json:"sendLog"`
	TaskName      string `json:"taskName"`
	MailingStatus string `json:"mailingStatus"`
	SubStatus     int    `json:"subStatus"`
	SoftStatus    string `json:"softStatus"`
	TimeStr       string `json:"timeStr"`
	Event         string `json:"event"`
	Receiver      string `json:"receiver"`
	Message       string `json:"message"`
	Email         string `json:"email"`
	Name          string `json:"name"`
	Phone         string `json:"phone"`
	SubStatusDesc string `json:"subStatusDesc"`
}

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

type OpenAndClickInfo struct {
	TrackType int    `json:"trackType"`
	APIUser   string `json:"apiUser"`
	Email     string `json:"email"`
	URL       string `json:"url"`
	CurrTime  string `json:"currTime"`
}

type OpenAndClickRsp struct {
	Result     bool   `json:"result"`
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
	Info       struct {
		Total    int                 `json:"total"`
		Count    int                 `json:"count"`
		DataList []*OpenAndClickInfo `json:"dataList"`
	} `json:"info"`
}
