package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"runtime"
	"strings"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello~"))
	})
	http.HandleFunc("/healthz/", healthz)

	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func healthz(w http.ResponseWriter, r *http.Request) {
	ri := HTTPReqInfo{
		method: r.Method,
		uri:    r.URL.String(),
		code:   http.StatusText(200),
	}
	ri.ipaddr = RemoteIP(r)

	fmt.Println("entering healthz handler")

	w.Header().Set("Header-Accept", r.Header.Get("Accept"))
	w.Header().Set("Header-User-Agent", r.Header.Get("User-Agent"))
	w.Header().Set("Version", runtime.Version())

	log.Printf("请求IP为 " + ri.ipaddr + "返回码：200")

	w.Write([]byte(ri.code))

}

// RemoteIP 通过 RemoteAddr 获取 IP 地址， 只是一个快速解析方法。
func RemoteIP(r *http.Request) string {
	if ip, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr)); err == nil {
		return ip
	}

	return ""
}

