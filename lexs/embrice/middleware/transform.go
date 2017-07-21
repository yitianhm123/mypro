package middleware

import (
	"lexs/x/logger"
	"net/http"
)

type transform struct {
	log *logger.Logger
}

func (self *transform) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	self.log.Debug("middleware transform...")
	next(w, r)
}

func Transform(log *logger.Logger) *transform {
	t := &transform{
		log: log,
	}
	return t
}
