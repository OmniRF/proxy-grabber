package proxy_grabber

import (
	"checkerx/pgk/utils"
	"strings"
)

func ParseProxy( url string, MaxProxyParsed int) ( []string , error){
	Document := utils.GetDocument(url)

	return strings.SplitN(Document.Text(), "\n", MaxProxyParsed), nil
}