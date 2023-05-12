package sendcloud

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

type UserInfoRsp struct {
	Result     bool      `json:"result"`
	StatusCode int       `json:"statusCode"`
	Message    string    `json:"message"`
	Info       *UserInfo `json:"info"`
}
