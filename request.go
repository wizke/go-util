package util

import (
	"fmt"
	"net/http"
)

func ClientRemoteIp(r *http.Request) string {
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
