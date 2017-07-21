package middleware

import (
	"lexs/x/logger"
	"net/http"
)

type rights struct {
	log *logger.Logger
}

func (self *rights) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	self.log.Debug("middleware rights...")
	next(w, r)
}

func Rights(log *logger.Logger) *rights {
	r := &rights{
		log: log,
	}
	return r
}
