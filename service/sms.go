package service

import (
	"bytes"
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

var responseXml = `<?xml version="1.0" encoding="UTF-8"?>
<env:Envelope 
	xmlns:env="http://www.w3.org/2003/05/soap-envelope" 
	xmlns:ns1="/rpc/sms/rpcserver" 
	xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" 
	xmlns:xsd="http://www.w3.org/2001/XMLSchema" 
	xmlns:ns2="http://xml.apache.org/xml-soap" 
	xmlns:enc="http://www.w3.org/2003/05/soap-encoding">
		<env:Body xmlns:rpc="http://www.w3.org/2003/05/soap-rpc">
			<ns1:rpcinitResponse env:encodingStyle="http://www.w3.org/2003/05/soap-encoding">
				<rpc:result>return</rpc:result>
				<return xsi:type="ns2:Map">
					<item>
						<key xsi:type="xsd:string">code</key>
						<value xsi:type="xsd:int">0</value>
					</item>
					<item>
						<key xsi:type="xsd:string">message</key>
						<value xsi:type="xsd:string">success</value>
					</item>
					<item>
						<key xsi:type="xsd:string">time</key>
						<value xsi:type="xsd:int">1587089482</value>
					</item>
					<item>
					<key xsi:type="xsd:string">data</key>
					<value xsi:type="ns2:Map">
						<item>
							<key xsi:type="xsd:string">order_no</key>
							<value xsi:type="xsd:string">20041710112204a36cb0384380</value>
						</item>
						<item>
							<key xsi:type="xsd:string">time</key>
							<value xsi:type="xsd:int">1587089482</value>
						</item>
					</value>
					</item>
				</return>
			</ns1:rpcinitResponse>
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
	var header = map[string]string{
		"Authorization": genAuthorization(smsConfig.Auth),
	}
	// fmt.Println(header)
	// fmt.Println(xmlTemplate)
	res := new(bytes.Buffer)
	err := Request.Post("http://msg-api.dadi01.net/rpc/sms/rpcserver", header, []byte(xmlTemplate), res)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(res)
}

func genXml(account string, paramJsonStr string, secretHash string) string {
	return fmt.Sprintf(xmlTemplate, account, paramJsonStr, secretHash)
}

func genAuthorization(auth string) string {
	//php soap的Authorization加密用的是 base64(auth+":")
	auth = auth + ":"
	return "Basic " + Base64ToString(auth)
}
