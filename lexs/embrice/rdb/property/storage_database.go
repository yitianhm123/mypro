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

func (self *Database) AddStorage(ctx context.Context, so *entity.StorageOrder) error {
	query := `INSERT INTO pp_storage_order(storage_no, property_no,property_name, model,pcount, color, size,status,brand, 
	created_id,created_date,source,remark,storage_date,deleted) 
	VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`
	result, err := self.Exec(query, so.StorageNo, so.PropertyNo, so.PropertyName, so.Model, so.Count, so.Color, so.Size, so.Status, so.Brand,
		so.CreatedId, so.CreatedDate, so.Source, so.Remark, so.StorageDate, so.Deleted)
	self.log.Trace("model=%s", so.Model)
	if err == nil {
		so.Id, err = result.LastInsertId()
	}
	return err
}

func (self *Database) GetStorageNo(ctx context.Context, so *entity.StorageOrder) error {
	query_storagno := `select case when max(SUBSTRING(storage_no,11,4)) is null then "" else max(SUBSTRING(storage_no,11,4)) end as storage_no from pp_storage_order  where SUBSTRING(storage_no,3,8)=?`
	t := time.Now().Format("2006-01-02")
	storageno := strings.Replace(t, "-", "", 2)
	fmt.Println(storageno)
	err := self.Get(so, query_storagno, storageno)
	return err
}

func (self *Database) GetStorageS(ctx context.Context, sos *[]entity.StorageOrder, so *entity.StorageOrder) error {
	query_all := `select count(1) as total from pp_storage_order where deleted = 0 and ""=?`
	err := self.Get(so, query_all, nil)
	if err != nil {
		fmt.Println(err)
	}

	query_storages := `select count(1) as pcount, storage_no,property_name,model,size,color,created_id, created_date,status,source,remark,storage_date from pp_storage_order 
	 where deleted=0 and (property_no=? or ""=?)
	 and (storage_no=? or ""=? )
	 and (remark=? or ""=?)
	 and (property_name= ? or ""=?)
	 and (source =? or ""=?)
	 and (storage_date >= ? or ""=?)
	 and (created_date >= ? or ""=?)
	 group by property_name,model,color,size,created_id,created_date,status,source,remark,storage_date,storage_no
	limit ?,?
	`

	if so.Page*so.PageSize > so.Total {
		err = errors.New("pageNo:" + strconv.Itoa(so.Page) + ",当前页数超过总页数")
	} else {
		fmt.Println(so.Page * so.PageSize)
		err = self.Select(sos, query_storages, so.PropertyNo, so.PropertyNo,
			so.StorageNo, so.StorageNo,
			so.Remark, so.Remark,
			so.PropertyName, so.PropertyName,
			so.Source, so.Source,
			so.StorageDate, so.StorageDate,
			so.CreatedDate, so.CreatedDate, so.Page*so.PageSize, so.PageSize)
		if err != nil {
			fmt.Println(err)
		}
	}
	return err
}

func (self *Database) GetStorageDetail(ctx context.Context, sos *[]entity.StorageOrder, so *entity.StorageOrder) error {
	query_storages := `select id,1 as pcount, storage_no,property_no,property_name,model,color,brand,size,created_id, created_date,status,source,remark,storage_date 
	from pp_storage_order 
	 where deleted=0 and (id = ? or ""=?)and (storage_no=? or ""=?) limit ?,?`

	err := self.Select(sos, query_storages, so.Id, so.Id, so.StorageNo, so.StorageNo, so.PageSize*so.Page, so.PageSize)

	if err != nil {
		fmt.Println(err)
	}
	return err
}

func (self *Database) UpdateStorageByPpNo(ctx context.Context, so *entity.StorageOrder) error {
	var err error
	update_storage := `update pp_storage_order set `
	var update_temp string
	var Id int64
	if len(so.PropertyName) > 0 {
		update_temp = update_storage + `property_name=? where id = ?`
		result, err := self.Exec(update_temp, so.PropertyName, so.Id)
		if err == nil {
			Id, err = result.RowsAffected()
		}

	}
	if len(so.Model) > 0 {
		update_temp = update_storage + `model=? where id= ?`
		fmt.Println(update_temp)
		fmt.Println(so.Model + "\n" + so.PropertyNo + "\n" + so.StorageNo)
		result, err := self.Exec(update_temp, so.Model, so.Id)
		if err == nil {
			Id, err = result.RowsAffected()
		}
	}
	if len(so.Color) > 0 {
		update_temp = update_storage + `color=? where id = ?`
		result, err := self.Exec(update_temp, so.Color, so.Id)
		if err == nil {
			Id, err = result.RowsAffected()

		}
	}
	if len(so.Brand) > 0 {
		update_temp = update_storage + `brand=? where id = ?`
		result, err := self.Exec(update_temp, so.Brand, so.Id)
		if err == nil {
			Id, err = result.RowsAffected()

		}
	}
	if len(so.Size) > 0 {
		update_temp = update_storage + `size=? where id = ? `
		result, err := self.Exec(update_temp, so.Size, so.Id)
		if err == nil {
			Id, err = result.RowsAffected()

		}
	}
	//如果更新状态则该单品未入库
	if strings.Contains(so.Status, "已入") {
		po := entity.Property{}
		po.Brand = so.Brand
		po.Color = so.Color
		po.Model = so.Model
		po.PropertyName = so.PropertyName
		po.Size = so.Size
		po.SuplrId = 123
		po.PropertyNo = so.PropertyNo
		po.Status = "1" //入库闲置
		po.CreatedDate = so.CreatedDate
		po.CreatedId = so.CreatedId
		tempproperty := po
		self.log.Trace(tempproperty.PropertyNo)
		if err = self.GetPropertyByNo(ctx, &tempproperty); err != nil {
			return err
		}
		if len(tempproperty.PropertyNo) > 0 {
			tx := self.MustBegin()
			update_temp = update_storage + `status=?, deleted=1 where id = ?`
			self.log.Trace(update_temp)
			fmt.Println(so.Id)
			tx.MustExec(update_temp, so.Status, so.Id)
			update_property := `update pp_property set status = 1 where property_no = ?`

			tx.MustExec(update_property, tempproperty.PropertyNo)
			err = tx.Commit()
			if err != nil {
				return err
			}
		} else {

			tx := self.MustBegin()

			update_temp = update_storage + `status=?, deleted=1 where id = ?`
			self.log.Trace(update_temp)
			fmt.Println(so.Id)
			tx.MustExec(update_temp, so.Status, so.Id)

			query_property := `INSERT INTO pp_property(property_no, property_name, model,size,
			color,brand,prop_unit,unt_price,status,suplr_id,deleted, created_id,created_date)
			 VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?)`
			tx.MustExec(query_property, po.PropertyNo, po.PropertyName, po.Model, po.Size,
				po.Color, po.Brand, po.PropUnit, po.UntPrice, po.Status, po.SuplrId, po.Deleted,
				po.CreatedId, po.CreatedDate)

			err = tx.Commit()
			if err != nil {
				return err
			}
		}
	}
	if len(so.Remark) > 0 {
		update_temp = update_storage + `remark=? where id = ?`
		result, err := self.Exec(update_temp, so.Remark, so.Id)
		if err == nil {
			Id, err = result.RowsAffected()

		}
	}
	if len(so.StorageDate) > 0 {
		update_temp = update_storage + `storage_date=? where id = ?`
		result, err := self.Exec(update_temp, so.StorageDate, so.Id)
		if err == nil {
			Id, err = result.RowsAffected()

		}
	}
	if so.Deleted > 0 {
		update_temp = update_storage + `deleted=? where id = ? `
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
