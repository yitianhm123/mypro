package middleware

import (
	"lexs/x/logger"
	"net/http"
)

type authenticate struct {
	log *logger.Logger
}

func (self *authenticate) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	self.log.Debug("middleware authenticate...")
	next(w, r)
}

func Authenticate(log *logger.Logger) *authenticate {
	a := &authenticate{
		log: log,
	}
	return a
}
