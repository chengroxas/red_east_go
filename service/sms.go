package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	. "red-east/utils"

	"github.com/beevik/etree"
)

var signName = map[string]string{
	"0": "01酒店",
}

type MapItem map[string]interface{}

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

type Sms struct {
	Mobile      string
	CountryCode string
	CropId      string
}

/**
 * 发送短信验证码
 * @param code 验证码
 * @return result 解析xml后得到的map[string]interface{}
 * @return err error request会发生的错误
 **/
func (self *Sms) SendVerCodeMsg(code string) error {
	msg := "你的短信验证码是:123456"
	_, err := self.senMsg(msg, "1")
	return err
}

func (self *Sms) senMsg(msg string, smsType string) (result MapItem, err error) {
	params := make(map[string]interface{})
	params["smsArea"] = "1"
	params["smstype"] = smsType
	params["mobile"] = self.Mobile
	params["message"] = msg
	params["smsMerchantSignName"] = signName[self.CropId]
	params["project_id"] = self.CropId
	paramJson, _ := json.Marshal(params)
	paramJsonStr := fmt.Sprintf("%s", paramJson)
	result, err = smsSaopRequest(paramJsonStr)
	if err != nil {
		// Logger.Error("发送短信请求失败")
		return nil, errors.New("短信验证码请求失败")
	}
	//当参数不正确的时候返回的code是string，参数正确返回的是int
	var smsBools bool
	switch result["code"].(type) {
	case int:
		smsBools = result["code"].(int) == 0
	case string:
		smsBools = result["code"].(string) == "0"
	}
	if !smsBools {
		// Logger.Error("短信验证码发送失败:", result["message"].(string), "参数:", paramJsonStr)
		return nil, errors.New("短信验证码发送失败")
	}
	return result, nil
}

func smsSaopRequest(paramJsonStr string) (mapItem MapItem, err error) {
	smsConfig := Config.Msg
	secretHash := Sha256ToString(Md5ToString(paramJsonStr) + smsConfig.Secret)
	xmlTemplate = genRequestXml(smsConfig.Account, paramJsonStr, secretHash)
	var header = map[string]string{
		"Authorization": genAuthorization(smsConfig.Auth),
	}
	res := new(bytes.Buffer)
	err = Request.Post("http://msg-api.dadi01.net/rpc/sms/rpcserver", header, []byte(xmlTemplate), res)
	if err != nil {
		return nil, err
	}
	mapItem = parseResponseXml(res.String())
	return mapItem, nil
}

func genRequestXml(account string, paramJsonStr string, secretHash string) string {
	return fmt.Sprintf(xmlTemplate, account, paramJsonStr, secretHash)
}

func genAuthorization(auth string) string {
	//php soap的Authorization加密用的是 base64(auth+":")
	auth = auth + ":"
	return "Basic " + Base64ToString(auth)
}

func parseResponseXml(resultXml string) MapItem {
	mapItem := make(MapItem)
	arrayItem := []MapItem
	doc := etree.NewDocument()
	doc.ReadFromString(resultXml)
	Envelope := doc.SelectElement("Envelope")
	Body := Envelope.SelectElement("Body")
	rpcinitResponse := Body.SelectElement("rpcinitResponse")
	result := rpcinitResponse.SelectElement("return")
	recursiveGetElement(result, mapItem)
	return mapItem
}

func recursiveGetElement(element *etree.Element, mapItem MapItem) {
	items := element.SelectElements("item")
	for _, item := range items {
		value := item.SelectElement("value")
		valueItems := value.SelectElements("item")
		key := item.SelectElement("key").Text()

		if len(valueItems) > 0 {
			valueType := value.SelectAttrValue("xsi:type", "ns2:Map")
			if valueType == "enc:Array" {
				arrayMap := make(MapItem)
				arrayDatas := []map[string]interface{}{}
				for _, iitem := range valueItems {
					recursiveGetElement(iitem, arrayMap)
					arrayDatas = append(arrayDatas, arrayMap)
				}
				mapItem[key] = arrayDatas
			} else {
				valueItem := make(MapItem)
				recursiveGetElement(value, valueItem)
				mapItem[key] = valueItem
			}
		} else {
			mapItem[key] = value.Text()
		}
	}
}
