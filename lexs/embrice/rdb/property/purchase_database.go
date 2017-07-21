package property

import (
	"context"
	"errors"
	"fmt"
	entity "lexs/embrice/entity/property"
	"strconv"
	"strings"
	"time"
)

func (self *Database) AddPurchaseOrder(ctx context.Context, so *entity.PurchaseOrder) error {
	query := `INSERT INTO pp_purchase_order(porder_no, suplr_id,property_name, model,count, color, size,status,brand, 
	created_id,created_date,remark,deleted) 
	VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?)`
	self.log.Trace(so.CreatedDate)
	result, err := self.Exec(query, so.PorderNo, so.SuplrId, so.PropertyName, so.Model, so.Count, so.Color, so.Size, so.Status, so.Brand,
		so.CreatedId, so.CreatedDate, so.Remark, 0)
	self.log.Trace("model=%s", so.Model)
	if err == nil {
		so.Id, err = result.LastInsertId()
	}
	return err
}

func (self *Database) GetPurchaseOrder(ctx context.Context, so *entity.PurchaseOrder) error {
	query_storagno := `select max(SUBSTRING(porder_no,11,4)) as porder_no from pp_purchase_order  where SUBSTRING(storage_no,3,8)=?`
	t := time.Now().Format("2006-01-02")
	storageno := strings.Replace(t, "-", "", 2)
	fmt.Println(storageno)
	err := self.Get(so, query_storagno, storageno)
	return err
}

func (self *Database) GetPurchaseOrderS(ctx context.Context, sos *[]entity.PurchaseOrder, so *entity.PurchaseOrder) error {
	query_all := `select count(1) as total from pp_purchase_order where ""=?`
	err := self.Get(so, query_all, nil)
	if err != nil {
		fmt.Println(err)
	}

	query_storages := `select id, porder_no,property_name,created_id, created_date,status,remark from pp_purchase_order 
	 where 1=1 and (property_name=? or ""=?)
	 and (porder_no=? or ""=? )
	 and (remark=? or ""=?)
	 and (created_date >= ? or ""=?)
	limit ?,10
	`

	if so.Page*so.PageSize > so.Total {
		err = errors.New("page:" + strconv.Itoa(so.Page) + ",当前页数超过总页数")
	} else {
		fmt.Println(so.Page * so.PageSize)
		err = self.Select(sos, query_storages, so.PropertyName, so.PropertyName,
			so.PorderNo, so.PorderNo,
			so.Remark, so.Remark,
			so.CreatedDate, so.CreatedDate, so.Page*so.PageSize)
		if err != nil {
			fmt.Println(err)
		}
	}
	return err
}

func (self *Database) GetPurchaseOrderDetail(ctx context.Context, sos *[]entity.PurchaseOrder, so *entity.PurchaseOrder) error {
	query_storages := `select id, porder_no,property_name,model,color,brand,size,created_id, created_date,status,remark 
	from pp_purchase_order 
	 where 1=1 and (id = ? or ""=?)and (porder_no=? or ""=?)`

	err := self.Select(sos, query_storages, so.Id, so.Id, so.PorderNo, so.PorderNo)

	if err != nil {
		fmt.Println(err)
	}
	return err
}

func (self *Database) UpdatePurchaseOrderById(ctx context.Context, so *entity.PurchaseOrder) error {
	var err error
	update_porder := `update pp_porder_order set `
	var update_temp string
	var Id int64
	if len(so.PropertyName) > 0 {
		update_temp = update_porder + `property_name=? where id = ?`
		result, err := self.Exec(update_temp, so.PropertyName, so.Id)
		if err == nil {
			Id, err = result.RowsAffected()

		}
	}
	if len(so.Model) > 0 {
		update_temp = update_porder + `model=? where id= ?`
		result, err := self.Exec(update_temp, so.Model, so.Id)
		if err == nil {
			Id, err = result.RowsAffected()

		}
	}
	if len(so.Color) > 0 {
		update_temp = update_porder + `color=? where id = ?`
		result, err := self.Exec(update_temp, so.Color, so.Id)
		if err == nil {
			Id, err = result.RowsAffected()

		}
	}
	if len(so.Brand) > 0 {
		update_temp = update_porder + `brand=? where id = ?`
		result, err := self.Exec(update_temp, so.Brand, so.Id)
		if err == nil {
			Id, err = result.RowsAffected()

		}
	}
	if len(so.Size) > 0 {
		update_temp = update_porder + `size=? where id = ? `
		result, err := self.Exec(update_temp, so.Size, so.Id)
		if err == nil {
			Id, err = result.RowsAffected()

		}
	}
	if len(so.Status) > 0 {
		update_temp = update_porder + `status=? where id = ?`
		result, err := self.Exec(update_temp, so.Status, so.Id)
		if err == nil {
			Id, err = result.RowsAffected()

		}
	}
	if len(so.Remark) > 0 {
		update_temp = update_porder + `remark=? where id = ?`
		result, err := self.Exec(update_temp, so.Remark, so.Id)
		if err == nil {
			Id, err = result.RowsAffected()

		}
	}
	if so.Deleted > 0 {
		update_temp = update_porder + `deleted=? where id = ? `
		result, err := self.Exec(update_temp, so.Deleted, so.Id)
		if err == nil {
			Id, err = result.RowsAffected()

		}
	}
	if Id > 0 {
		so.Id = Id
	}
	return err
}
