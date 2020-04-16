package service

import (
	"encoding/json"
	"fmt"
	. "red-east/utils"
)

var signName = map[int]string{
	0: "01酒店",
}

var xmlTemplate = `<?xml version="1.0" encoding="UTF-8"?>
<env:Envelope
	xmlns:env="http://www.w3.org/2003/05/soap-envelope"
	xmlns:ns1="rpc/sms/rpcserver"
	xmlns:xsd="http://www.w3.org/2001/XMLSchema"
	xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
	xmlns:ns2="http://xml.apache.org/xml-soap"
	xmlns:enc="http://www.w3.org/2003/05/soap-encoding">
	<env:Body>
		<ns1:rpcinit env:encodingStyle="http://www.w3.org/2003/05/soap-encoding">
			<param0 xsi:type="xsd:string">send</param0>
			<param1 xsi:type="ns2:Map">
				<item>
					<key xsi:type="xsd:string">account</key>
					<value xsi:type="xsd:string">%s</value>
					</item>
				<item>
					<key xsi:type="xsd:string">params_json_str</key>
					<value xsi:type="xsd:string">%s</value>
				</item>
				<item>
					<key xsi:type="xsd:string">secret_hash</key>
					<value xsi:type="xsd:string">%s</value>
				</item>
			</param1>
			<param2 xsi:type="xsd:string">rpc/sms/send</param2>
		</ns1:rpcinit>
	</env:Body>
</env:Envelope>`

//发送短信
func SendMsg(mobile string, countryCode int, msg string, corpId int) {
	params := make(map[string]interface{})
	params["smsArea"] = 1
	params["smstype"] = 1
	params["mobile"] = mobile
	params["message"] = msg
	params["smsMerchantSignName"] = signName[corpId]
	params["project_id"] = corpId
	paramJson, _ := json.Marshal(params)
	paramJsonStr := fmt.Sprintf("%s", paramJson)
	smsSaopRequest(paramJsonStr)
}

func smsSaopRequest(paramJsonStr string) {
	smsConfig := Config.Msg
	secretHash := Sha256ToString(Md5ToString(paramJsonStr) + smsConfig.Secret)
	xmlTemplate = genXml(smsConfig.Account, paramJsonStr, secretHash)
	fmt.Println(xmlTemplate)
}

func genXml(account string, paramJsonStr string, secretHash string) string {
	return fmt.Sprintf(xmlTemplate, account, paramJsonStr, secretHash)
}
