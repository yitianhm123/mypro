package main

import (
	"encoding/json"
	"fmt"
	"lexs/embrice/constant"
	"lexs/embrice/entity"
	"lexs/embrice/extension"
	"lexs/embrice/validate"
	"lexs/x/crypt"
	"net/http"
)

// curl -X POST https://api.lexs.com/account/login -d "account=rock&password=123456"
func (self *server) login(w http.ResponseWriter, r *http.Request) {
	var (
		err        error
		token      string
		inAccount  entity.Account
		outAccount entity.Account
	)

	inAccount.IP = extension.GetRealIP(r)

	defer r.Body.Close()
	if err = json.NewDecoder(r.Body).Decode(&inAccount); err != nil {
		goto FAILED
	}
	if err = validate.Account(inAccount.Account, true); err != nil {
		goto FAILED
	}
	if err = validate.Password(inAccount.Password, true); err != nil {
		goto FAILED
	}

	// todo: from cache
	outAccount.Account = inAccount.Account
	if err = self.db.Login(self.ctx, &outAccount); err != nil {
		goto FAILED
	}

	if err = crypt.ValidatePwd(inAccount.Password, outAccount.Password); err != nil {
		goto FAILED
	}

	token = crypt.NewToken(fmt.Sprintf("%d", inAccount.ID), inAccount.IP, constant.TokenTimeout, constant.PrivKey)
	// err = crypt.ValidateToken(fmt.Sprintf("%d", inAccount.Account), token, inAccount.IP, constant.PrivKey)
	w.Write([]byte(fmt.Sprintf(`{"id":%d,"acc":"%s","token":"%s"}`, inAccount.ID, inAccount.Account, token)))
	self.log.Trace("login %s@%s successed", inAccount.Account, inAccount.IP)
	return

FAILED:
	w.Write([]byte(fmt.Sprintf(`{"code":101, "msg":"%s"}`, err)))
	self.log.Trace("login %s@%s failed: %s", inAccount.Account, inAccount.IP, err)
}
