package remoteip

import (
	"errors"
	"net"
	"net/http"
	"strings"
)

func GetRemoteIP(req *http.Request) string {
	var ip string
	header := req.Header
	realIp := header.Get("X-Real-Ip")
	forwardedFor := header.Get("X-Forwarded-For")
	if realIp == "" && forwardedFor == "" {
		ip, _ = parseIP(req.RemoteAddr)
	}
	if forwardedFor != "" {
		parts := strings.Split(forwardedFor, ",")
		for i, p := range parts {
			parts[i] = strings.TrimSpace(p)
		}
		ip = parts[0]
	}

	return ip
}

func parseIP(s string) (string, error) {
	ip, _, err := net.SplitHostPort(s)
	if err == nil {
		return ip, nil
	}

	ip2 := net.ParseIP(s)
	if ip2 == nil {
		return "", errors.New("invalid IP")
	}

	return ip2.String(), nil
}
