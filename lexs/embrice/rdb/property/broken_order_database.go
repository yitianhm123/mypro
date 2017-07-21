package property

import (
	"context"
	entity "lexs/embrice/entity/property"
	"lexs/x/logger"
	_ "lexs/x/mysql"
	"lexs/x/sqlx"
)

type Database struct {
	*sqlx.DB
	log *logger.Logger
}

func NewDB(log *logger.Logger, maxConns, maxIdles int) *Database {
	// todo: config mysql connect
	db := &Database{
		log: log,
		DB:  sqlx.MustConnect("mysql", "devsql:devsql123.com@tcp(10.10.10.121:3306)/Rocket"),
		// DB: sqlx.MustConnect("mysql", "root:root@tcp(127.0.0.1:3306)/Rocket"),
		// DB: sqlx.MustConnect("mysql", "root:123456@tcp(10.188.10.2:3306)/aaa_4"),
	}
	db.SetMaxOpenConns(maxConns)
	db.SetMaxIdleConns(maxIdles)
	return db
}

func (self *Database) AddBrokenOrder(ctx context.Context, bro *entity.BrokenOrder) error {
	query := `INSERT INTO pp_broken_order(broken_order_no, property_no, status,remark,
	deleted,created_id,created_date) 
	VALUES (?,?,?,?,?,?,?)`
	result, err := self.Exec(query, bro.BrokenOrderNo, bro.PropertyNo, bro.Status,
		bro.Remark, bro.Deleted,
		bro.CreatedId, bro.CreatedDate)

	if err == nil {
		bro.Id, err = result.LastInsertId()
	}
	return err
}

func (self *Database) GetBrokenOrder(ctx context.Context, bos *[]entity.BrokenOrder, bo *entity.BrokenOrder) error {
	query_stock := `select id, broken_order_no, property_no, status,remark,deleted
	created_id, created_date
	 from pp_broken_order 
	 where 1=1 and (property_no=? or ""=?)
	 and (remark=? or ""=?)
	 and (status =? or ""=?)
	 and (created_date >= ? or ""=?)
	limit ?,?`

	err := self.Select(bos, query_stock, bo.PropertyNo, bo.PropertyNo,
		bo.Remark, bo.Remark, bo.Status, bo.Status,
		bo.CreatedDate, bo.CreatedDate, bo.Page*bo.PageSize, bo.PageSize)

	return err
}

func (self *Database) UpdateBrokenOrderByPpno(ctx context.Context, so *entity.BrokenOrder) error {
	var err error
	update_storage := `update pp_broken_order set `
	var update_temp string
	var Id int64
	if len(so.Status) > 0 {
		update_temp = update_storage + `status=? where id = ?`
		result, err := self.Exec(update_temp, so.Status, so.Id)
		if err == nil {
			Id, err = result.RowsAffected()

		}
	}
	if len(so.Remark) > 0 {
		update_temp = update_storage + `remark=? where id = ?`
		result, err := self.Exec(update_temp, so.Remark, so.Id)
		if err == nil {
			Id, err = result.RowsAffected()

		}
	}
	if so.Deleted > 0 {
		update_temp = update_storage + `deleted=? where id = ?`
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

func (self *Database) Close() {
	self.DB.Close()
}
