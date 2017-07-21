package main

import (
	"encoding/json"
	"fmt"
	entity "lexs/embrice/entity/property"
	"net/http"
)

// curl -X POST https://localhost:8002/property/getProperty -d '{"account":"rock"}'    "account=123"
func (self *server) getPropertyByNo(w http.ResponseWriter, r *http.Request) {
	var (
		err        error
		inProperty entity.Property
		response   entity.Response
	)

	defer r.Body.Close()
	if err = json.NewDecoder(r.Body).Decode(&inProperty); err != nil {
		goto FAILED
	}

	if err = self.db.GetPropertyByNo(self.ctx, &inProperty); err != nil {
		goto FAILED
	}
	if err == nil {
		response.Code = 0
		response.Message = "success"
		response.Data = inProperty
	} else {
		goto FAILED
	}
	if err = json.NewEncoder(w).Encode(response); err != nil {
		goto FAILED
	}
	self.log.Trace("getPropertyByNo %s successed", inProperty.PropertyNo)
	return

FAILED:
	w.Write([]byte(fmt.Sprintf(`{"code":%d, "message":"%s","data":""}`, 1, err)))
	self.log.Trace("getPropertyByNo %s failed: %s", inProperty.PropertyNo, err)
}

func (self *server) getPropertyByName(w http.ResponseWriter, r *http.Request) {
	var (
		err         error
		inProperty  entity.Property
		outProperty []entity.Property
		response    entity.Response
	)

	defer r.Body.Close()
	if err = json.NewDecoder(r.Body).Decode(&inProperty); err != nil {
		goto FAILED
	}

	if err = self.db.GetPropertyByName(self.ctx, &outProperty, &inProperty); err != nil {
		goto FAILED
	}
	if err == nil {
		response.Code = 0
		response.Message = "success"
		response.Data = inProperty
	}
	if err = json.NewEncoder(w).Encode(response); err != nil {
		goto FAILED
	}
	self.log.Trace("getPropertyByName %s ,%s successed", inProperty.PropertyName, inProperty.PropertyNo)
	return

FAILED:
	w.Write([]byte(fmt.Sprintf(`{"code":1, "message":"%s","data":""}`, err)))
	self.log.Trace("getPropertyByName %s@%s failed: %s", inProperty.PropertyName, inProperty.PropertyNo, err)
}

func (self *server) addProperty(w http.ResponseWriter, r *http.Request) {
	var (
		err        error
		inProperty entity.Property
	)

	defer r.Body.Close()
	if err = json.NewDecoder(r.Body).Decode(&inProperty); err != nil {
		goto FAILED
	}

	if err = self.db.AddProperty(self.ctx, &inProperty); err != nil {
		goto FAILED
	}

	self.log.Trace("addProperty %s success: ", inProperty.PropertyNo)
	return
FAILED:
	w.Write([]byte(fmt.Sprintf(`{"code":%d, "message":"%s","data":""}`, 1, err)))
	self.log.Trace("addProperty %s,%s failed: %s", inProperty.PropertyName, inProperty.PropertyNo, err)

}

func (self *server) updatePropertyByNo(w http.ResponseWriter, r *http.Request) {
	var (
		err        error
		response   entity.Response
		inProperty entity.Property
	)

	defer r.Body.Close()
	if err = json.NewDecoder(r.Body).Decode(&inProperty); err != nil {
		goto FAILED
	}

	if err = self.db.UpdatePropertyByNo(self.ctx, &inProperty); err != nil {
		goto FAILED
	}
	response.Data = inProperty.Id
	response.Code = 0
	response.Message = "success"
	if err = json.NewEncoder(w).Encode(response); err != nil {
		goto FAILED
	}
	self.log.Trace("updatePropertyByNo success: %s ", inProperty.PropertyNo)
	return
FAILED:
	w.Write([]byte(fmt.Sprintf(`{"code":%d, "message":"%s","data":""}`, 1, err)))
	self.log.Trace("updatePropertyByNo %s@%s failed: %s", inProperty.PropertyName, inProperty.PropertyNo, err)

}
