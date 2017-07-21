package middleware

import (
	"net/http"
)

type check struct {
}

func Check(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	// log.Debug("middleware validate...")
	next(w, r)
}
