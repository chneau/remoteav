package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

func proxy(w http.ResponseWriter, r *http.Request) {
	url, _ := url.Parse("http://localhost:5173")

	proxy := httputil.ReverseProxy{Director: func(r *http.Request) {
		r.URL.Scheme = url.Scheme
		r.URL.Host = url.Host
		r.URL.Path = url.Path + r.URL.Path
		r.Host = url.Host
	}}
	proxy.ServeHTTP(w, r)
}
