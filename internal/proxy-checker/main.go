package proxy_checker

import (
	"h12.io/socks"
	"net"
	"net/http"
	"strconv"
)

type ProxyInfo struct {
	addr string
	result bool
}

const (
	SOCKS5 int = 0
	SOCKS4     = 1
	HTTP       = 2
)

func CheckSocksProxy(ProxyAddr string, ProxyType int, Timeout int) (*ProxyInfo, error){

	var TempInfoProxy 	ProxyInfo

	var dialSocksProxy func(string, string) (net.Conn, error)

	switch ProxyType{
	case SOCKS5:
		dialSocksProxy = socks.Dial("socks5://" + ProxyAddr + "?timeout=" + strconv.Itoa(Timeout) + "s")
	case SOCKS4:
		dialSocksProxy  = socks.Dial("socks4://"  + ProxyAddr +  "?timeout=" + strconv.Itoa(Timeout) + "s")
	}

	tr := &http.Transport{Dial: dialSocksProxy}
	httpClient := &http.Client{Transport: tr}
	resp, err := httpClient.Get("http://www.google.com")
	if err != nil {
		TempInfoProxy.addr = ProxyAddr
		TempInfoProxy.result = false
		return &TempInfoProxy, err
	}


	if ( resp.Body == nil) {
		TempInfoProxy.addr = ProxyAddr
		TempInfoProxy.result = false
		return &TempInfoProxy, err
	}



		TempInfoProxy.addr = ProxyAddr
		TempInfoProxy.result = true
		return &TempInfoProxy, err
}

func CheckProxyList ( ProxyList []string, TypeProxy int, Timeout int) (  GoodProxy []string, BadProxy[]string){
	for i := 0; i < len(ProxyList); i++ {
		TempInfo, _ := CheckSocksProxy(ProxyList[i], TypeProxy, Timeout)
		if TempInfo.result == true {
			print ("Founded good proxy\n")
			GoodProxy = append( GoodProxy, TempInfo.addr )
		}else{
			print ("Founded bad proxy\n")
			BadProxy = append( BadProxy, TempInfo.addr  )
		}
	}
	print( len(GoodProxy))
	return
}