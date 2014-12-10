package sharecount

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/ikeikeikeike/gopkg/convert"
)

type Result struct {
	Service string
	Count   int
}

var (
	context http.ResponseWriter
)

func Fetch(url string) (results []Result) {
	c := make(chan Result)

	go func() { c <- hatena(url) }()
	go func() { c <- twitter(url) }()
	go func() { c <- facebook(url) }()
	go func() { c <- linkedin(url) }()

	timeout := time.After(5000 * time.Millisecond)
	for i := 0; i < 4; i++ {
		select {
		case result := <-c:
			results = append(results, result)
		case <-timeout:
			log.Println("timed out")
			return
		}
	}
	return
}

func fetchUrl(url string) (body []byte, err error) {
	response, err := http.Get(url)
	if err != nil {
		return
	}

	defer response.Body.Close()
	body, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}

	return
}

func hatena(url string) (result Result) {
	result.Service = "hatena"
	result.Count = -1

	endpoint := fmt.Sprintf("http://api.b.st-hatena.com/entry.count?url=%s", url)
	body, err := fetchUrl(endpoint)
	if err != nil {
		log.Printf("Error in getting response from Hatena: %s", err.Error())
		return
	}

	i, _ := convert.StrTo(string(body)).Int()
	result.Count = i
	return
}

func twitter(url string) (result Result) {
	result.Service = "twitter"
	result.Count = -1

	endpoint := fmt.Sprintf("http://urls.api.twitter.com/1/urls/count.json?url=%s", url)
	body, err := fetchUrl(endpoint)
	if err != nil {
		log.Printf("Error in getting response from Twitter: %s", err.Error())
		return
	}

	type twitterResult struct {
		Count int
	}
	var t twitterResult

	err = json.Unmarshal(body, &t)
	if err != nil {
		log.Printf("Error in parsing response from Twitter: %s", err.Error())
		return
	}

	result.Count = t.Count
	return
}

func facebook(url string) (result Result) {
	result.Service = "facebook"
	result.Count = -1

	endpoint := fmt.Sprintf("http://api.ak.facebook.com/restserver.php?v=1.0&method=links.getStats&format=json&urls=%s", url)
	body, err := fetchUrl(endpoint)
	if err != nil {
		log.Printf("Error in getting response from Facebook: %s", err.Error())
		return
	}

	type facebookResult struct {
		Total_count int
	}
	var f []facebookResult

	err = json.Unmarshal(body, &f)
	if err != nil {
		log.Printf("Error in parsing response from Facebook: %s", err.Error())
		return
	}

	result.Count = f[0].Total_count
	return
}

func linkedin(url string) (result Result) {
	result.Service = "linkedin"
	result.Count = -1
	endpoint := fmt.Sprintf("http://www.linkedin.com/countserv/count/share?format=json&url=%s", url)
	body, err := fetchUrl(endpoint)
	if err != nil {
		log.Printf("Error in getting response from LinkedIn: %s", err.Error())
		return
	}

	type linkedinResult struct {
		Count int
	}
	var l linkedinResult

	err = json.Unmarshal(body, &l)
	if err != nil {
		log.Printf("Error in parsing response from LinkedIn: %s", err.Error())
		return
	}

	result.Count = l.Count
	return
}
