package main

import (
	"encoding/json"
	"fmt"
	entity "lexs/embrice/entity/property"
	"net/http"
)

// curl -X POST https://localhost:8002/property/getProperty -d '{"account":"rock"}'    "account=123"
func (self *server) getStockoutDetail(w http.ResponseWriter, r *http.Request) {
	var (
		err         error
		token       string
		instockout  entity.Stockout
		outstockout []entity.Stockout
		response    entity.Response
	)

	defer r.Body.Close()
	if err = json.NewDecoder(r.Body).Decode(&instockout); err != nil {
		goto FAILED
	}

	if err = self.db.GetStockoutDetail(self.ctx, &outstockout, &instockout); err != nil {
		goto FAILED
	}
	if err == nil {
		response.Code = 0
		response.Message = "success"
		response.Data = outstockout
	} else {
		goto FAILED
	}
	w.Write([]byte(fmt.Sprintf(`{"id":%d,"acc":"%s","token":"%s"}`, instockout.Id, instockout.PropertyNo, token)))
	self.log.Trace("getStockoutDetail %s,%,s%s successed", instockout.StockoutNo, instockout.PropertyNo, 1)
	return

FAILED:
	w.Write([]byte(fmt.Sprintf(`{"code":101, "msg":"%s"}`, err)))
	self.log.Trace("getStockoutDetail %s@%s failed: %s", instockout.StockoutNo, instockout.PropertyNo, err)
}

func (self *server) addStockout(w http.ResponseWriter, r *http.Request) {
	var (
		err        error
		instockout entity.Stockout
		response   entity.Response
	)

	defer r.Body.Close()
	if err = json.NewDecoder(r.Body).Decode(&instockout); err != nil {
		goto FAILED
	}

	if err = self.db.AddStockout(self.ctx, &instockout); err != nil {
		goto FAILED
	}
	if err == nil {
		response.Code = 0
		response.Message = "success"
		response.Data = instockout
	}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		goto FAILED
	}

	self.log.Trace("AddStockout %s,%s,%s successed", instockout.PropertyNo, instockout.StockoutNo, instockout.Status)
	return

FAILED:
	w.Write([]byte(fmt.Sprintf(`{"code":%d, "message":"%s"}`, 1, err)))
	self.log.Trace("AddStockout %s,%s failed: %s", instockout.StockoutNo, instockout.PropertyNo, err)
}

func (self *server) updateStockout(w http.ResponseWriter, r *http.Request) {
	var (
		err        error
		instockout entity.Stockout
		response   entity.Response
	)

	defer r.Body.Close()
	if err = json.NewDecoder(r.Body).Decode(&instockout); err != nil {
		goto FAILED
	}

	if err = self.db.UpdateStockoutByPpNo(self.ctx, &instockout); err != nil {
		goto FAILED
	}
	if err == nil {
		response.Code = 0
		response.Message = "success"
		response.Data = instockout
	}
	if err = json.NewEncoder(w).Encode(response); err != nil {
		goto FAILED
	}
	self.log.Trace("UpdateStockoutByPpNo PropertyNo=%s,StockoutNo=%s,status = %s successed", instockout.PropertyNo, instockout.StockoutNo, instockout.Status)
	return

FAILED:
	w.Write([]byte(fmt.Sprintf(`{"code":%d, "message":"%s","data":""}`, 1, err)))
	self.log.Trace("UpdateStockoutByPpNo StockoutNo=%s,PropertyNo=%s failed: %s", instockout.StockoutNo, instockout.PropertyNo, err)
}
