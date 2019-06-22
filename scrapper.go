package main
//
//import (
//	"bytes"
//	"fmt"
//	"github.com/PuerkitoBio/goquery"
//	"log"
//	"net/http"
//	"net/url"
//	"strings"
//)
//
//func ScrapeHTML() {
//	request_url := "https://flix.co.ke"
//
//	form := url.Values{
//		"AJAX":   {"search"},
//		"cinema": {"0"},
//		"date":   {"0"},
//		"movie":  {"0"},
//	}
//
//	body := bytes.NewBufferString(form.Encode())
//	resp, err := http.Post(request_url, "application/x-www-form-urlencoded", body)
//
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer resp.Body.Close()
//
//	if resp.StatusCode != 200 {
//		log.Fatalf("status code error: %d %s", resp.StatusCode, resp.Status)
//	}
//
//	doc, err := goquery.NewDocumentFromReader(resp.Body)
//	if err != nil{
//		log.Fatal(err)
//	}
//
//	doc.Find("a").Each(func(i int, s *goquery.Selection){
//		fmt.Println(strings.TrimSpace(s.Find("h3").Text()))
//	})
//
//	//io.Copy(os.Stdout, resp.Body)
//}
//
//func main() {
//	ScrapeHTML()
//}
