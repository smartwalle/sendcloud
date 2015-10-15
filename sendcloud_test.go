package sendcloud

import (
	"testing"
	"fmt"
)

func Test_SendMailWithTemplate(t *testing.T) {
	UpdateApiInfo("", "")

	var to = make([]map[string]string, 1)
	to[0] = map[string]string{"to":"917996695@qq.com", "%url%": "http://www.baidu.com"}

	var ok, err, result = SendMailWithTemplate("template name", "from mail", "from name", "replay to email address", "subject", to)

	fmt.Println(ok, err, result)
}
