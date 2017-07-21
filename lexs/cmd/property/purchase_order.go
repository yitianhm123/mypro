package main

import (
	"encoding/json"
	"fmt"
	entity "lexs/embrice/entity/property"
	"net/http"
)

// curl -X POST https://localhost:8002/property/getProperty -d '{"account":"rock"}'    "account=123"
func (self *server) getPurchaseOrderDetail(w http.ResponseWriter, r *http.Request) {
	var (
		err               error
		inPurchaseOrder   entity.PurchaseOrder
		outPurchaseOrders []entity.PurchaseOrder
		response          entity.Response
	)

	defer r.Body.Close()
	if err = json.NewDecoder(r.Body).Decode(&inPurchaseOrder); err != nil {
		goto FAILED
	}

	if err = self.db.GetPurchaseOrderDetail(self.ctx, &outPurchaseOrders, &inPurchaseOrder); err != nil {
		goto FAILED
	}
	if err == nil {
		response.Code = 0
		response.Message = "success"
		response.Data = outPurchaseOrders
	}
	if err = json.NewEncoder(w).Encode(response); err != nil {
		goto FAILED
	}
	self.log.Trace("getPurchaseOrderDetail %s successed", inPurchaseOrder.PorderNo)
	return

FAILED:
	w.Write([]byte(fmt.Sprintf(`{"code":%d, "message":"%s","data":""}`, 1, err)))
	self.log.Trace("getPurchaseOrderDetail %s failed: %s", inPurchaseOrder.PorderNo, err)
}

// curl -X POST https://localhost:8002/property/createPurchaseOrder -d '{"account":"rock"}'
func (self *server) createPurchaseOrder(w http.ResponseWriter, r *http.Request) {
	var (
		err             error
		inPurchaseOrder entity.PurchaseOrder
		response        entity.Response
	)

	defer r.Body.Close()
	if err = json.NewDecoder(r.Body).Decode(&inPurchaseOrder); err != nil {
		goto FAILED
	}

	if err = self.db.AddPurchaseOrder(self.ctx, &inPurchaseOrder); err != nil {
		goto FAILED
	}
	if err == nil {
		response.Code = 0
		response.Message = "success"
		response.Data = inPurchaseOrder
	}
	if err = json.NewEncoder(w).Encode(inPurchaseOrder); err != nil {
		goto FAILED
	}
	self.log.Trace("createPurchaseOrder %s,%s successed", inPurchaseOrder.PorderNo, inPurchaseOrder.Status)
	return

FAILED:
	w.Write([]byte(fmt.Sprintf(`{"code":%d, "message":"%s","data":""}`, 1, err)))
	self.log.Trace("createPurchaseOrder %s failed: %s", inPurchaseOrder.PorderNo, err)
}

func (self *server) updatePurchaseOrderById(w http.ResponseWriter, r *http.Request) {
	var (
		err             error
		inPurchaseOrder entity.PurchaseOrder
		response        entity.Response
	)

	defer r.Body.Close()
	if err = json.NewDecoder(r.Body).Decode(&inPurchaseOrder); err != nil {
		goto FAILED
	}

	if err = self.db.UpdatePurchaseOrderById(self.ctx, &inPurchaseOrder); err != nil {
		goto FAILED
	}
	if err == nil {
		response.Code = 0
		response.Message = "success"
		response.Data = inPurchaseOrder.PorderNo
	}
	if err = json.NewEncoder(w).Encode(response); err != nil {
		goto FAILED
	}

	self.log.Trace("updatePurchaseOrderById %s,%s,%s successed", inPurchaseOrder.PorderNo, inPurchaseOrder.Status)
	return

FAILED:
	w.Write([]byte(fmt.Sprintf(`{"code":%d, "message":"%s","data":""}`, 1, err)))
	self.log.Trace("updatePurchaseOrderById %s,%s failed: %s", inPurchaseOrder.PorderNo, inPurchaseOrder.Status, err)
}
