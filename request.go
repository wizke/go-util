package util

import (
	"fmt"
	"net/http"
)

func ClientRemoteIpAll(r *http.Request) string {
	var clientIP string
	for k, v := range r.Header {
		if k == "X-Real-Ip" {
			if len(v) == 1 {
				clientIP = v[0]
			} else {
				clientIP = fmt.Sprintf("%s", v)
			}
			break
		}
	}
	if clientIP != "" {
		return clientIP
	}
	return r.RemoteAddr
}

func ClientRemoteIpOne(r *http.Request) string {
	var clientIP string
	clientIP = r.Header.Get("X-Real-Ip")
	if clientIP != "" {
		return clientIP
	}
	return r.RemoteAddr
}
