package main

import (
	"encoding/json"
	"fmt"
	_ "lexs/embrice/constant"
	entity "lexs/embrice/entity/property"
	_ "lexs/embrice/extension"
	_ "lexs/embrice/validate"
	_ "lexs/x/crypt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// curl -X POST http://localhost:8001/property/addStorage -d '[{"property_name":"苹果","model":"{\"cpu\":\"i5\",\"memory\":\"8g\"}","count":2,"color":"白色","size":"13.5"}]'
func (self *server) addStorage(w http.ResponseWriter, r *http.Request) {
	var (
		err        error
		temp       int
		tempcount  int
		outStorage []entity.StorageOrder
		inStorage  []entity.StorageOrder
		storage    entity.StorageOrder
		t3         string
		response   entity.Response
		storageNo  string
	)

	defer r.Body.Close()
	if err = json.NewDecoder(r.Body).Decode(&inStorage); err != nil {
		self.log.Trace(err.Error())
	}
	t := time.Now().Format("2006-01-02 15:04:05")
	t1 := time.Now().Format("2006-01-02")
	t3 = strings.Replace(t1, "-", "", 2)
	err = self.db.GetStorageNo(self.ctx, &storage)
	if err != nil {
		goto FAILED
	}
	storageNo = storage.StorageNo
	self.log.Trace("storageno=%s", storageNo)
	if len(storageNo) == 0 {
		storageNo = strings.Repeat("0", 1) + "1"
	} else {
		temp, err = strconv.Atoi(storageNo)
		if err != nil {
			goto FAILED
		}
		temp = temp + 1
		storageNo = strings.Repeat("0", 2-len(strconv.Itoa(temp))) + strconv.Itoa(temp)
	}
	storageNo = "PI" + t3 + storageNo
	tempcount = 0
	for _, storage := range inStorage {
		tempcount = tempcount + storage.Count
		if storage.Count > 0 && len(storage.PropertyNo) == 0 {
			for i := tempcount - storage.Count; i < tempcount; i++ {
				I := strconv.Itoa(i)
				l := len(I)
				I = strings.Repeat("0", 4-l) + I
				storage.CreatedId = 123123123127
				storage.CreatedDate = t
				if len(storage.StorageDate) == 0 {
					storage.StorageDate = t
				}
				storage.Status = "未入库"
				storage.Deleted = 0
				storage.StorageNo = storageNo
				storage.PropertyNo = "PP-" + storageNo + "-" + I
				self.log.Trace(storage.PropertyNo)
				err = self.db.AddStorage(self.ctx, &storage)
				if err != nil {
					goto FAILED
				}
				outStorage = append(outStorage, storage)
			}
		} else {
			storage.CreatedId = 123123123127
			storage.CreatedDate = t
			if len(storage.StorageDate) == 0 {
				storage.StorageDate = t
			}
			storage.Status = "未入库"
			storage.Deleted = 0
			storage.StorageNo = storageNo
			err = self.db.AddStorage(self.ctx, &storage)
			if err != nil {
				goto FAILED
			}
			outStorage = append(outStorage, storage)
		}
	}

	if err == nil {
		response.Code = 0
		response.Data = outStorage
		response.Message = "success"
	}
	if err = json.NewEncoder(w).Encode(response); err != nil {
		goto FAILED
	}
	//	token = crypt.NewToken(fmt.Sprintf("%d", inProperty.Id), inProperty.PropertyNo, constant.TokenTimeout, constant.PrivKey)
	//  err = crypt.ValidateToken(fmt.Sprintf("%d", inProperty.PropertyNo), token, inProperty.PropertyName, constant.PrivKey)
	self.log.Trace("addStorage %s@%s%s successed", storage.PropertyName, storage.PropertyNo, storage.PropertyName)
	return

FAILED:
	w.Write([]byte(fmt.Sprintf(`{"code":1, "message":"%s"，"data":""}`, err)))
	self.log.Trace("addStorage %s@%s failed: %s", storage.PropertyName, storage.PropertyNo, err)
}

// curl -X POST http://localhost:8001/property/get_storages -d '{"property_name":"苹果","storage_no":"PI2017071901"}'
func (self *server) getStorages(w http.ResponseWriter, r *http.Request) {
	var (
		err        error
		outStorage []entity.StorageOrder
		storage    entity.StorageOrder
		response   entity.Response
	)

	defer r.Body.Close()
	if err = json.NewDecoder(r.Body).Decode(&storage); err != nil {
		goto FAILED
	}

	err = self.db.GetStorageS(self.ctx, &outStorage, &storage)
	if err == nil {
		response.Code = 0
		response.Data = outStorage
		response.Message = "success"
	} else {
		goto FAILED
	}
	if err = json.NewEncoder(w).Encode(response); err != nil {
		goto FAILED
	}
	self.log.Trace("getstorages success: the number of the getorder is  %d ", len(outStorage))
	return

FAILED:
	w.Write([]byte(fmt.Sprintf(`{"code":1, "message":"%s","data":""}`, err)))
	self.log.Trace("getstorages %s,%s failed: %s", storage.PropertyName, storage.PropertyNo, err)
}

// curl -X POST http://localhost:8001/property/get_storage_detail -d '{storage_no":"PI2017071901"}'
func (self *server) getStorageDetail(w http.ResponseWriter, r *http.Request) {
	var (
		err        error
		response   entity.Response
		outStorage []entity.StorageOrder
		storage    entity.StorageOrder
	)

	defer r.Body.Close()
	if err = json.NewDecoder(r.Body).Decode(&storage); err != nil {
		goto FAILED
	}

	err = self.db.GetStorageDetail(self.ctx, &outStorage, &storage)
	if err == nil {
		response.Code = 0
		response.Message = "success"
		response.Data = outStorage
	} else {
		goto FAILED
	}
	if err = json.NewEncoder(w).Encode(response); err != nil {
		goto FAILED
	}
	self.log.Trace("GetStorageDetail the count of the getorder is  %d success", len(outStorage))
	return

FAILED:
	w.Write([]byte(fmt.Sprintf(`{"code":1, "message":"%s","data":""}`, err)))
	self.log.Trace("GetStorageDetail PropertyName=%s,PropertyNo=%s failed: %s", storage.PropertyName, storage.PropertyNo, err)
}

// curl -X POST http://localhost:8001/property/updateStorageByPpno -d '{"property_name":"苹果","storage_no":"PI201707130001"}'
func (self *server) updateStorageByPpno(w http.ResponseWriter, r *http.Request) {
	var (
		err       error
		response  entity.Response
		inStorage []entity.StorageOrder
		storage   entity.StorageOrder
	)

	defer r.Body.Close()
	if err = json.NewDecoder(r.Body).Decode(&inStorage); err != nil {
		goto FAILED
	}
	for _, storage := range inStorage {
		err = self.db.UpdateStorageByPpNo(self.ctx, &storage)
	}

	if err == nil {
		response.Code = 0
		response.Message = "success"
		response.Data = storage
	} else {
		goto FAILED
	}
	if err = json.NewEncoder(w).Encode(response); err != nil {
		goto FAILED
	}
	self.log.Trace("updateStorageByPpNO success ")
	return

FAILED:
	w.Write([]byte(fmt.Sprintf(`{"code":%d, "message":"%s","data":""}`, 1, err)))
	self.log.Trace("UpdateStorageByPpNo %s@%s failed: %s", storage.PropertyName, storage.PropertyNo, err)
}

func (self *server) hello(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	var (
		propertyNo string
	)
	//	self.db.GetStorageNo(self.ctx, propertyNo)

	self.log.Trace(propertyNo)
	w.Write([]byte("hello world!\n"))

	self.log.Trace("propertyNo = %s", propertyNo)

}
