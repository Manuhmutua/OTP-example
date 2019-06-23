package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
	"time"
)

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
func main() {
	// Set account keys & information
	accountSid := "ACfd36d2574047c8202d744f55dabd409b"
	authToken := "5036288a47ca6ea8bde8254417a3b35f"
	urlStr := "https://api.twilio.com/2010-04-01/Accounts/" + accountSid + "/Messages.json"

	// Create possible message bodies
	quotes := [7]string{"I urge you to please notice when you are happy, and exclaim or murmur or think at some point, 'If this isn't nice, I don't know what is.'",
		"Peculiar travel suggestions are dancing lessons from God.",
		"There's only one rule that I know of, babies—God damn it, you've got to be kind.",
		"Many people need desperately to receive this message: 'I feel and think much as you do, care about many of the things you care about, although most people do not care about them. You are not alone.'",
		"That is my principal objection to life, I think: It's too easy, when alive, to make perfectly horrible mistakes.",
		"So it goes.",
		"We must be careful about what we pretend to be."}

	// Set up rand
	rand.Seed(time.Now().Unix())

	msgData := url.Values{}
	msgData.Set("To", "+254714353160")
	msgData.Set("From", "+18635327586")
	msgData.Set("Body", quotes[0])
	msgDataReader := *strings.NewReader(msgData.Encode())

	client := &http.Client{}
	req, _ := http.NewRequest("POST", urlStr, &msgDataReader)
	req.SetBasicAuth(accountSid, authToken)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, _ := client.Do(req)
	if (resp.StatusCode >= 200 && resp.StatusCode < 300) {
		var data map[string]interface{}
		decoder := json.NewDecoder(resp.Body)
		err := decoder.Decode(&data)
		if (err == nil) {
			fmt.Println(data["sid"])
		}
	} else {
		fmt.Println(resp.Status);
	}
}
