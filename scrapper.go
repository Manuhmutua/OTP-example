//package main
//
//import (
//	"fmt"
//	"github.com/xlzd/gotp"
//)

//import (
//	"encoding/json"
//	"fmt"
//	"math/rand"
//	"net/http"
//	"net/url"
//	"strings"
//	"time"
//)

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

//package main
//
//import (
//	"encoding/json"
//	"fmt"
//	"math/rand"
//	"net/http"
//	"net/url"
//	"strings"
//	"time"
//)
//


package main

import (
"fmt"
	"time"

	"github.com/xlzd/gotp"
)


func main() {
	//fmt.Println("Random secret:", gotp.RandomSecret(16))
	defaultTOTPUsage()
}

func defaultTOTPUsage() {
	otpS := gotp.NewDefaultTOTP("4S62BZNFXXSZLCRO")
	fmt.Println(otpS.ProvisioningUri("demoAccountName", "issuerName"))
	otp := otpS.Now()
	fmt.Println("current one-time password is:", otp)
	defaultHOTPUsage(otp)

	//fmt.Println(otp.Verify("179394", 1524485781))
}

func defaultHOTPUsage(otps string) {
	otp := gotp.NewDefaultTOTP("4S62BZNFXXSZLCRO")
	fmt.Println(otp.ProvisioningUri("demoAccountName", "issuerName"))
	now := time.Now()
	fmt.Println(otp.Verify(otps, int(now.Unix())))
}