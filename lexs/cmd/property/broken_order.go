package main

import (
	"encoding/json"
	"fmt"
	entity "lexs/embrice/entity/property"
	"net/http"
)

// curl -X POST https://localhost:8002/property/get_broken_order -d '{"account":"rock"}'    "account=123"
func (self *server) getBrokenorder(w http.ResponseWriter, r *http.Request) {
	var (
		err            error
		inBrokenorder  entity.BrokenOrder
		outBrokenorder []entity.BrokenOrder
		response       entity.Response
	)

	defer r.Body.Close()
	if err = json.NewDecoder(r.Body).Decode(&inBrokenorder); err != nil {
		goto FAILED
	}

	if err = self.db.GetBrokenOrder(self.ctx, &outBrokenorder, &inBrokenorder); err != nil {
		goto FAILED
	}
	if err == nil {
		response.Code = 0
		response.Message = "success"
		response.Data = outBrokenorder
	}
	if err = json.NewEncoder(w).Encode(response); err != nil {
		goto FAILED
	}
	self.log.Trace("GetBrokenOrder %s,%s successed", inBrokenorder.BrokenOrderNo, inBrokenorder.PropertyNo)
	return

FAILED:
	w.Write([]byte(fmt.Sprintf(`{"code":%d, "message":"%s","data":""}`, 1, err)))
	self.log.Trace("GetBrokenOrder %s,%s failed: %s", inBrokenorder.BrokenOrderNo, inBrokenorder.PropertyNo, err)
}

func (self *server) addBrokenorder(w http.ResponseWriter, r *http.Request) {
	var (
		err           error
		inBrokenorder entity.BrokenOrder
		response      entity.Response
	)

	defer r.Body.Close()
	if err = json.NewDecoder(r.Body).Decode(&inBrokenorder); err != nil {
		goto FAILED
	}

	if err = self.db.AddBrokenOrder(self.ctx, &inBrokenorder); err != nil {
		goto FAILED
	}
	if err == nil {
		response.Code = 0
		response.Message = "success"
		response.Data = inBrokenorder
	}
	if err = json.NewEncoder(w).Encode(inBrokenorder); err != nil {
		goto FAILED
	}
	self.log.Trace("AddBrokenOrder %s,%s,%s successed", inBrokenorder.PropertyNo, inBrokenorder.BrokenOrderNo, inBrokenorder.Status)
	return

FAILED:
	w.Write([]byte(fmt.Sprintf(`{"code":%d, "message":"%s","data":""}`, 1, err)))
	self.log.Trace("AddBrokenOrder %s@%s failed: %s", inBrokenorder.BrokenOrderNo, inBrokenorder.PropertyNo, err)
}

func (self *server) updateBrokenOrderByPpno(w http.ResponseWriter, r *http.Request) {
	var (
		err           error
		inBrokenorder entity.BrokenOrder
		response      entity.Response
	)

	defer r.Body.Close()
	if err = json.NewDecoder(r.Body).Decode(&inBrokenorder); err != nil {
		goto FAILED
	}

	if err = self.db.UpdateBrokenOrderByPpno(self.ctx, &inBrokenorder); err != nil {
		goto FAILED
	}
	if err == nil {
		response.Code = 0
		response.Message = "success"
		response.Data = inBrokenorder.BrokenOrderNo
	}
	if err = json.NewEncoder(w).Encode(response); err != nil {
		goto FAILED
	}

	self.log.Trace("UpdateBrokenOrderByPpno %s,%s,%s successed", inBrokenorder.PropertyNo, inBrokenorder.BrokenOrderNo, inBrokenorder.Status)
	return

FAILED:
	w.Write([]byte(fmt.Sprintf(`{"code":%d, "message":"%s","data":""}`, 1, err)))
	self.log.Trace("UpdateBrokenOrderByPpno %s,%s failed: %s", inBrokenorder.BrokenOrderNo, inBrokenorder.PropertyNo, err)
}
