package main

import (
	 "checkerx/internal/proxy-checker"
	"checkerx/internal/proxy-grabber"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

func main(){

	print("proxy-grabber/checker v 0.1 by OmniRF\n")
	print("Enter url:")

	var url string

	fmt.Scanln(&url)

	print("Enter amount of proxy parse:")
	var AmountProxy int
	fmt.Scanln(&AmountProxy)


	print("Enter type of proxy ( socks5 - 0; socks 4 - 1)  :")
	var ProxyType int
	fmt.Scanln(&ProxyType)

	ProxyList, _ := proxy_grabber.ParseProxy(url, AmountProxy)

	print ( "Proxy grabbed successfully | Amount:"  + strconv.Itoa(len ( ProxyList) ) + "\n" )

	GoodProxy, BadProxy := proxy_checker.CheckProxyList(ProxyList, ProxyType, 5000)

	file, _ := os.Create("BadProxy.json")
	JsonData, _ := json.Marshal(BadProxy)
	file.Write(JsonData)

	file, _ = os.Create("GoodProxy.json")
	JsonData, _ = json.Marshal(GoodProxy)
	file.Write(JsonData)


}