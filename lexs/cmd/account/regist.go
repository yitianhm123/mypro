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

// curl -X POST http://127.0.0.1:8000/account/regist -d "account=rock&password=123456&mobile=15900520751"
func (self *server) regist(w http.ResponseWriter, r *http.Request) {
	var (
		err   error
		token string
		ac    entity.Account
	)

	ac.IP = extension.GetRealIP(r)

	defer r.Body.Close()
	if err = json.NewDecoder(r.Body).Decode(&ac); err != nil {
		goto FAILED
	}

	if err = validate.Account(ac.Account, true); err != nil {
		goto FAILED
	}
	if err = validate.Password(ac.Password, true); err != nil {
		goto FAILED
	}
	if err = validate.Mobile(ac.Mobile, false); err != nil {
		goto FAILED
	}
	if err = validate.Email(ac.Email, false); err != nil {
		goto FAILED
	}

	// todo: check account conflict in redis

	// todo: check ip risk in redis

	ac.Password = crypt.EncryptPwd(ac.Password)
	if err = self.db.Reg(self.ctx, &ac); err != nil {
		goto FAILED
	}

	token = crypt.NewToken(fmt.Sprintf("%d", ac.ID), ac.IP, constant.TokenTimeout, constant.PrivKey)
	// err = crypt.ValidateToken(fmt.Sprintf("%d", a.ID), token, ip, constant.PrivKey)
	w.Write([]byte(fmt.Sprintf(`{"id":%d,"account":"%s","token":"%s"}`, ac.ID, ac.Account, token)))
	self.log.Trace("regist %s@%s successed", ac.Account, ac.IP)
	return

FAILED:
	w.Write([]byte(fmt.Sprintf(`{"code":100, "msg":"%s"}`, err)))
	self.log.Trace("regist %s@%s failed: %s", ac.Account, ac.IP, err)
}
