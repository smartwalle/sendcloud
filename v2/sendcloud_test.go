package v2

import (
	"testing"
	"fmt"
)

func Test_SendMailWithTemplate(t *testing.T) {
	UpdateApiInfo("", "")

	var to = make([]map[string]string, 1)
	to[0] = map[string]string{"to":"917996695@qq.com", "%url%": "http://www.baidu.com"}

	var ok, err, result = SendTemplateMail("mail_verify_cn", "service@smoktech.com", "SMOK", "service@smoktech.com", "SMOK", to, "")

	fmt.Println(ok, err, result)
}