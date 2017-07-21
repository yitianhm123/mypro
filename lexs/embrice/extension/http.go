package extension

import (
	"net"
	"net/http"
)

func GetRealIP(r *http.Request) string {
	// todo: return x-forwarded-for for cdn
	if ip, _, e := net.SplitHostPort(r.RemoteAddr); e == nil {
		return ip
	}
	return r.RemoteAddr
}
