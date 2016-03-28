package libs

import (
	"fmt"
	"submail"
)

//submail上发送短信
func SendSmsOnSubMail(phone string, proJectCode string, vars ...string) {
	fmt.Println(phone, proJectCode, vars)
	messageconfig := make(map[string]string)
	messageconfig["appid"] = "11051"
	messageconfig["appkey"] = "3e65ea76df71362d2cbe7f362e240d4a"
	messageconfig["signtype"] = "md5"
	//messagexsend
	messagexsend := submail.CreateMessageXSend()
	submail.MessageXSendAddTo(messagexsend, phone)
	submail.MessageXSendSetProject(messagexsend, proJectCode)
	if l := len(vars); l > 0 {
		for k := 0; k < l; k += 2 {
			submail.MessageXSendAddVar(messagexsend, vars[k], vars[k+1])
		}
	}
	fmt.Println("MessageXSend ", submail.MessageXSendRun(submail.MessageXSendBuildRequest(messagexsend), messageconfig))
}
