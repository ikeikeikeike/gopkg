package http

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/ikeikeikeike/gopkg/convert"
)

func BuildRequestUrl(req *http.Request, uri string) string {
	if uri == "" {
		uri = req.RequestURI
	}
	return fmt.Sprintf("%s:%s%s", Site(req), convert.ToStr(Port(req)), uri)
}

func Port(req *http.Request) int {
	parts := strings.Split(req.Host, ":")
	if len(parts) == 2 {
		port, _ := strconv.Atoi(parts[1])
		return port
	}
	return 80
}

func Site(req *http.Request) string {
	return Scheme(req) + "://" + Domain(req)
}

func Scheme(req *http.Request) string {
	if req.URL.Scheme != "" {
		return req.URL.Scheme
	}
	if req.TLS == nil {
		return "http"
	}
	return "https"
}

func Domain(req *http.Request) string {
	return Host(req)
}

func Host(req *http.Request) string {
	if req.Host != "" {
		hostParts := strings.Split(req.Host, ":")
		if len(hostParts) > 0 {
			return hostParts[0]
		}
		return req.Host
	}
	return "localhost"
}

func UserAgent(req *http.Request) string {
	return req.Header.Get("User-Agent")
}
