package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

// ReverseProxy struct
type ReverseProxy struct {
	target *url.URL
	proxy  *httputil.ReverseProxy
}

// StringSliceIncludes checks if a string slice contains a specific value
func StringSliceIncludes(slice []string, val string) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

// NewReverseProxy creates a new instance of ReverseProxy
func NewReverseProxy(target string) *ReverseProxy {
	url, _ := url.Parse(target)
	return &ReverseProxy{
		target: url,
		proxy:  httputil.NewSingleHostReverseProxy(url),
	}
}

// HandleRequest processes the incoming request
func (p *ReverseProxy) HandleRequest(res http.ResponseWriter, req *http.Request) {
	allowOrigins := strings.Split(os.Getenv("ALLOWED_ORIGINS"), ",") // get allowed origins from .env
	origin := req.Header.Get("Origin")                               // get the origin header
	if StringSliceIncludes(allowOrigins, origin) {
		res.Header().Set("Access-Control-Allow-Origin", origin)
	} else {
		res.Header().Set("Access-Control-Allow-Origin", "")
		log.Println("Origin not allowed")
		log.Println("Origin: ", origin)
	}
	req.URL.Host = p.target.Host
	req.URL.Scheme = p.target.Scheme
	req.Header.Set("X-Forwarded-Host", req.Header.Get("Host"))
	req.Host = p.target.Host
	req.Header.Del("Origin") // delete the Origin header
	res.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	res.Header().Set("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With")

	if req.Method == "OPTIONS" {
		res.WriteHeader(http.StatusOK)
		return
	}
	p.proxy.ServeHTTP(res, req) // serve the request
}
func main() {
	log.Println("Starting reverse proxy server")
	log.Println("Version: 0.0.1")
	log.Println("Loading .env file")
	// Load .env file
	err := godotenv.Load("config/.env")
	if err != nil {
		log.Fatal("Error loading config/.env file")
	}
	proxy := NewReverseProxy(os.Getenv("UPSTREAM_DOMAIN"))
	http.HandleFunc("/", proxy.HandleRequest)
	log.Println("Listening on port 3000")
	http.ListenAndServe(":3000", nil)
}
