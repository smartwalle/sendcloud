package main

import (
	"github.com/smartwalle/log4go"
	"github.com/smartwalle/sendcloud"
)

func main() {
	var sc = sendcloud.New("dm.bdvapeus.com", "ENEC2JVSvRk2Qa89")

	rsp, err := sc.SendTemplateMailToAddressList("vip0003@maillist.sendcloud.org____", "tpl20190420", "vip@dm.bdvapeus.com", "BD VAPE", "vip@bdvapeus.com", "Welcome to join BD VAPE FAM- 100 Early Birds Wanted", nil)
	if err != nil {
		log4go.Errorln(err)
		return
	}
	log4go.Infoln("EmailId", rsp.Info.EmailIdList)
	log4go.Infoln("TaskId", rsp.Info.MailListTaskId)
}
