package sendcloud

import (
	"fmt"
	"net/url"
)

const (
	kEmailStatus      = "http://api.sendcloud.net/apiv2/data/emailStatus"
	kOpenAndClickList = "http://api.sendcloud.net/apiv2/openandclick/list"
)

// GetEmailStatusList 投递回应
// 目前只提供了几个常用参数，其它参数信息请参考 http://www.sendcloud.net/doc/email_v2/deliver_response/#_1
// start    int  否  查询起始位置, 取值区间 [0-], 默认为 0
// limit    int  否  查询个数, 取值区间 [0-100], 默认为 100
func (this *Client) GetEmailStatusList(days int, startDate, endDate string, start, limit int) (result *EmailStatusListRsp, err error) {
	var values = url.Values{}

	if days > 0 {
		values.Add("days", fmt.Sprintf("%d", days))
	}
	if startDate != "" {
		values.Set("startDate", startDate)
	}
	if endDate != "" {
		values.Set("endDate", endDate)
	}
	if start >= 0 {
		values.Add("start", fmt.Sprintf("%d", start))
	}
	if limit >= 1 {
		values.Add("limit", fmt.Sprintf("%d", limit))
	}

	if err = this.doRequest(kEmailStatus, values, &result); err != nil {
		return nil, err
	}

	if result.Result == false {
		return nil, fmt.Errorf("%d-%s", result.StatusCode, result.Message)
	}

	return result, err
}

// GetOpenAndClickList 获取打开点击数据
// 目前只提供了几个常用参数，其它参数信息请参考 http://www.sendcloud.net/doc/email_v2/open_or_click/#_1
func (this *Client) GetOpenAndClickList(days int, startDate, endDate string, start, limit int) (result *OpenAndClickRsp, err error) {
	var values = url.Values{}

	if days > 0 {
		values.Add("days", fmt.Sprintf("%d", days))
	}
	if startDate != "" {
		values.Set("startDate", startDate)
	}
	if endDate != "" {
		values.Set("endDate", endDate)
	}
	if start >= 0 {
		values.Add("start", fmt.Sprintf("%d", start))
	}
	if limit >= 1 {
		values.Add("limit", fmt.Sprintf("%d", limit))
	}

	if err = this.doRequest(kOpenAndClickList, values, &result); err != nil {
		return nil, err
	}

	if result.Result == false {
		return nil, fmt.Errorf("%d-%s", result.StatusCode, result.Message)
	}

	return result, err
}
