package verify

import (
	"lexs/embrice/constant"
	"strconv"
)

func Account(ac string, pwd string, must bool) error {
	return nil
}

func Token(uid string, token string, must bool) error {
	if len(uid) < 3 || len(token) < 16 {
		return constant.AError
	}
	if _, err := strconv.Atoi(uid); err != nil {
		return constant.AError
	}
	return nil
}
