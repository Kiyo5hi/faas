package function

import (
	"encoding/json"
	"net"
	"net/http"
	"strings"
	"time"
)

type You struct {
	Ip   string `json:"ip"`
	Time string `json:"time"`
}

func Handle(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	you := You{
		Ip:   getIPAddress(r),
		Time: time.Now().UTC().Format(time.RFC3339),
	}

	jsonStr, _ := json.Marshal(you)
	w.Write(jsonStr)
}

func getIPAddress(r *http.Request) string {
	forwarded := r.Header.Get("X-Forwarded-For")
	if forwarded != "" {
		ips := strings.Split(forwarded, ",")
		return strings.TrimSpace(ips[0])
	}

	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return ""
	}

	return ip
}
