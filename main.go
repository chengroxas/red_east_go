package main

// import (
// 	"io"
// 	"log"

// 	"red-east/dao/database"
// 	"red-east/middleware"
// 	"red-east/router"

// 	"red-east/utils/logging"

// 	"github.com/gin-gonic/gin"
// )

import (
	// "strings"
	// "red-east/config"
	// "red-east/service"
	// . "red-east/utils"
	// "red-east/utils/external"
	"fmt"

	"github.com/beevik/etree"
)

var resultXml = `<?xml version="1.0" encoding="UTF-8"?>
<env:Envelope xmlns:env="http://www.w3.org/2003/05/soap-envelope" xmlns:ns1="/rpc/project/rpcserver" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:ns2="http://xml.apache.org/xml-soap" xmlns:enc="http://www.w3.org/2003/05/soap-encoding">
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
                    <value xsi:type="xsd:string"></value>
                </item>
                <item>
                    <key xsi:type="xsd:string">time</key>
                    <value xsi:type="xsd:int">1587088353</value>
                </item>
                <item>
                    <key xsi:type="xsd:string">data</key>
                    <value xsi:type="ns2:Map">
                        <item>
                            <key xsi:type="xsd:int">2</key>
                            <value xsi:type="ns2:Map">
                                <item>
                                    <key xsi:type="xsd:string">id</key>
                                    <value xsi:type="xsd:int">2</value>
                                </item>
                                <item>
                                    <key xsi:type="xsd:string">name</key>
                                    <value xsi:type="xsd:string">日子里</value>
                                </item>
                                <item>
                                    <key xsi:type="xsd:string">area</key>
                                    <value xsi:type="xsd:float">23222.2299999999996</value>
                                </item>
                                <item>
                                    <key xsi:type="xsd:string">intro</key>
                                    <value xsi:type="xsd:string"></value>
                                </item>
                                <item>
                                    <key xsi:type="xsd:string">address</key>
                                    <value xsi:type="xsd:string"></value>
                                </item>
                                <item>
                                    <key xsi:type="xsd:string">logo</key>
                                    <value xsi:type="xsd:string">https://yde.oss-cn-shenzhen.aliyuncs.com/20190409/2cc193f6536d46b9a189271251371594.png</value>
                                </item>
                            </value>
                        </item>
                        <item>
                            <key xsi:type="xsd:int">3</key>
                            <value xsi:type="ns2:Map">
                                <item>
                                    <key xsi:type="xsd:string">id</key>
                                    <value xsi:type="xsd:int">3</value>
                                </item>
                                <item>
                                    <key xsi:type="xsd:string">name</key>
                                    <value xsi:type="xsd:string">大族项目</value>
                                </item>
                                <item>
                                    <key xsi:type="xsd:string">area</key>
                                    <value xsi:type="xsd:float">23222.2299999999996</value>
                                </item>
                                <item>
                                    <key xsi:type="xsd:string">intro</key>
                                    <value xsi:type="xsd:string"></value>
                                </item>
                                <item>
                                    <key xsi:type="xsd:string">address</key>
                                    <value xsi:type="xsd:string">深圳</value>
                                </item>
                                <item>
                                    <key xsi:type="xsd:string">logo</key>
                                    <value xsi:type="xsd:string"></value>
                                </item>
                            </value>
                        </item>
                        <item>
                            <key xsi:type="xsd:int">1</key>
                            <value xsi:type="ns2:Map">
                                <item>
                                    <key xsi:type="xsd:string">id</key>
                                    <value xsi:type="xsd:int">1</value>
                                </item>
                                <item>
                                    <key xsi:type="xsd:string">name</key>
                                    <value xsi:type="xsd:string">自由里</value>
                                </item>
                                <item>
                                    <key xsi:type="xsd:string">area</key>
                                    <value xsi:type="xsd:float">25000</value>
                                </item>
                                <item>
                                    <key xsi:type="xsd:string">intro</key>
                                    <value xsi:type="xsd:string"></value>
                                </item>
                                <item>
                                    <key xsi:type="xsd:string">address</key>
                                    <value xsi:type="xsd:string">广州市花都区新雅街东镜村自由人花园一街24号首层101-103室、负一层自编101-102室</value>
                                </item>
                                <item>
                                    <key xsi:type="xsd:string">logo</key>
                                    <value xsi:type="xsd:string">https://yde.oss-cn-shenzhen.aliyuncs.com/20190409/7910f64b34a727ebe07569f1be19a176.png</value>
                                </item>
                            </value>
                        </item>
                    </value>
                </item>
            </return>
        </ns1:rpcinitResponse>
    </env:Body>
</env:Envelope>`

type MapItem map[string]interface{}

func main() {
	// mapItem := make(MapItem)
	doc := etree.NewDocument()
	doc.ReadFromString(resultXml)
	Envelope := doc.SelectElement("Envelope")
	Body := Envelope.SelectElement("Body")
	rpcinitResponse := Body.SelectElement("rpcinitResponse")
	result := rpcinitResponse.SelectElement("return")
	items := result.SelectElements("item")
	for _, item := range items {
		// key := item.SelectElement("key").Text()

		value := item.SelectElement("value")
		// mapItem[key] = value.Text()
		data := recursiveGetChildElement(value)
		fmt.Println(data)
		// xsiType := value.SelectAttrValue("xsi:type", "string")
		// valueType := strings.Split(xsiType, ":")[1]
		// if valueType == "Map" {
		// 	valueItem := value.SelectElements("item")
		// 	for _, vitem := range valueItem {
		// 		vitemKey := vitem.SelectElement("key")
		// 		fmt.Println(vitemKey.Text())
		// 		vitemValue := vitem.SelectElement("value")
		// 		fmt.Println(vitemValue.Text())
		// 	}
		// }
	}
	// fmt.Println(mapItem)
	// var err error
	// Config, err = config.InitConfig()
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// op := external.Option{
	// 	Timeout:   10,
	// 	KeepAlive: 1,
	// 	MaxIdle:   1,
	// }
	// err = Request.Init(&op)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// service.SendMsg("15818359718", 86, "123456", 0)
}

func recursiveGetChildElement(element *etree.Element) (v interface{}) {
	//无限递归获得子元素里的值
	childElements := element.ChildElements()
	if len(childElements) == 0 {
		return
	}
	item := make(map[string]interface{})
	var data []map[string]interface{}
	for _, childElement := range childElements {
		value := childElement.SelectElement("value")
		if len(value.ChildElements()) != 0 {
			recursiveGetChildElement(value)
		}
		item[childElement.SelectElement("key").Text()] = value.Text()
		data = append(data, item)
		return data
	}
	return data
}

// func main() {
// 	var err error
// 	//初始化配置
// 	Config, err = config.InitConfig()
// 	if err != nil {
// 		log.Fatalln("init config fail", err.Error())
// 	}
// 	//初始化日志
// 	Logger, err = logging.InitLogger()
// 	if err != nil {
// 		log.Fatalln("init logger fail", err.Error())
// 	}

// 	//初始化数据库，使用mysql
// 	DB, err = mysql.InitMySql()
// 	if err != nil {
// 		Logger.Error("init database fail", err.Error())
// 	}
// 	defer DB.Close()
// 	// 初始化gin
// 	// gin输出到文件或者终端
// 	gin.DefaultWriter = io.MultiWriter(logging.Writers...)
// 	//去掉颜色
// 	gin.DisableConsoleColor()
// 	r := gin.New()
// 	//中间件要在路由注册之前
// 	r.Use(gin.Logger(), middleware.CheckCommonParam(), middleware.CheckSign())

// 	//注册路由
// 	router.RegisterRouter(r)
// 	//使用中间件，记录请求
// 	Logger.Info("start application....")
// 	r.Run()
// }
