package utils

import (
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

func GetDocument(url string) *goquery.Document {

	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Print(err)
		return nil
	}

	req.Header.Set("User-Agent", "telegram:@omnirf")
	resp, err := client.Do(req)

	if err != nil {
		log.Print(err)
	} else {

		defer resp.Body.Close()

		doc, err := goquery.NewDocumentFromReader(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		//log.Printf("GET %s successfully\n", url)
		return doc
	}

	return nil
}

func ClearIP( DirtIP string) string{

	r1, _ := regexp.Compile(`^.......|..$`)
	ClearedIP := r1.ReplaceAllString(DirtIP, "")
	return ClearedIP
}

func GetUserIP( url string) (string, error){
	res, err := http.Get(url)
	if err != nil {

		return "", err
	}

	body1, _ := ioutil.ReadAll(res.Body)

	UserIP := ClearIP(string(body1))

	return UserIP, err
}
