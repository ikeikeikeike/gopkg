package http

import (
	"net/http"
	"net/http/cookiejar"
	urlparse "net/url"
)

var DefaultUserAgent = "Mozilla/5.0 UserBehavior"

type UserBehavior struct {
	request *http.Request
	client  *http.Client
}

func NewUserBehavior() *UserBehavior {
	var request http.Request
	request.Method = "GET"
	request.Header = http.Header{}
	request.Header.Set("User-Agent", DefaultUserAgent)

	jar, _ := cookiejar.New(nil)

	return &UserBehavior{
		request: &request,
		client:  &http.Client{Jar: jar},
	}
}

func (u *UserBehavior) Header(key, value string) {
	u.request.Header.Set(key, value)
}

func (u *UserBehavior) Behave(urlStr string) (*http.Response, error) {
	u.request.URL, _ = urlparse.Parse(urlStr)
	return u.client.Do(u.request)
}

func (u *UserBehavior) Get(urlStr string) (*http.Response, error) {
	return u.Behave(urlStr)
}

func (u *UserBehavior) Post(urlStr string) (*http.Response, error) {
	u.request.URL, _ = urlparse.Parse(urlStr)
	u.request.Method = "POST"
	return u.client.Do(u.request)
}
