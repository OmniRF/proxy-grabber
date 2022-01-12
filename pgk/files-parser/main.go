package files_parser

import (
	"bufio"
	"os"
)



const (
	HTTP int = 0
	Socks   = 1
)


func ParseUrlsFromFile( path string ) ([]string, error) {
	var Links []string
	file, err := os.Open( path )
	if(err != nil){
		return Links, err
	}

	defer file.Close()

	FileScaner := bufio.NewScanner(file)

	for FileScaner.Scan(){
		Links = append(Links,  FileScaner.Text())
	}

	return Links, nil
}

func ParseProxyFromFile( path string,  ProxyType int) ([]string, error) {

	var Proxies []string

	file, err := os.Open(path)
	if( err != nil ){
		return Proxies, err
	}

	defer file.Close()

	FileScaner := bufio.NewScanner(file)

	for FileScaner.Scan() {

		switch ProxyType {
		case HTTP:
			Proxies = append(Proxies, "http://"+FileScaner.Text())
		case Socks:
			Proxies = append(Proxies, FileScaner.Text())
		}
		
	}

	return Proxies, nil
}