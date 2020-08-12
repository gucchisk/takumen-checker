package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	// "github.com/PuerkitoBio/goquery"
)

func main() {
	// searchUrl := "https://www.takumen.com/search"
	// res, err := http.Get(searchUrl)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// defer res.Body.Close()
	// cookies := res.Cookies()
	// cookieStr := ""
	// for i, cookie := range(cookies) {
	// 	if i > 0 {
	// 		cookieStr += ";"
	// 	}
	// 	cookieStr += cookie.Name + "=" + cookie.Value
	// }
	// fmt.Printf("cookie: %s\n", cookieStr)
	// doc, err := goquery.NewDocumentFromReader(res.Body)
	// var token string
	// doc.Find("meta").Each(func(_ int, s *goquery.Selection) {
	// 	name, _ := s.Attr("name")
	// 	if name == "csrf-token" {
	// 		token, _ = s.Attr("content")
	// 	}
	// })
	// fmt.Printf("token: %v\n", token)


	productsUrl := "https://www.takumen.com/search/index.js"
	values := url.Values{}
	values.Set("stock", "true")
	values.Add("order_by", "total")
	values.Add("all", "true")
	// values.Add("authenticity_token", token)
	
	req, err := http.NewRequest(
		"POST",
		productsUrl,
		strings.NewReader(values.Encode()),
	)
	if err != nil {
		fmt.Println(err)
		return
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	// req.Header.Set("Cookie", cookieStr)
	// req.Header.Set("X-CSRF-Token", token)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s\n", string(b))
}
