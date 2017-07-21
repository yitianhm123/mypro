package property

import (
	"context"
	entity "lexs/embrice/entity/property"
)

func (self *Database) AddStockout(ctx context.Context, sto *entity.Stockout) error {
	query := `INSERT INTO pp_stockout_order(stockout_no, property_no, stokout_date,function,
	remark,status,deleted,
	created_id,created_date) 
	VALUES (?,?,?,?,?,?,?,?,?,)`
	result, err := self.Exec(query, sto.StockoutNo, sto.PropertyNo, sto.StokoutDate,
		sto.Function, sto.Remark, sto.Status, sto.Deleted,
		sto.CreatedId, sto.CreatedDate)

	if err == nil {
		sto.Id, err = result.LastInsertId()
	}
	return err
}

func (self *Database) GetStockoutDetail(ctx context.Context, stks *[]entity.Stockout, sto *entity.Stockout) error {
	query_stock := `select id, stockout_no,property_no,stokout_date,function,remark,status,deleted,
	created_id, created_date
	 from pp_stockout_order 
	 where deleted=0 and (property_no=? or ""=?)
	 and (stokout_date=? or ""=? )
	 and (remark=? or ""=?)
	 and (function= ? or ""=?)
	 and (status =? or ""=?)
	 and (created_date >= ? or ""=?)
	limit ?,?
	`
	err := self.Select(stks, query_stock, sto.PropertyNo, sto.PropertyNo, sto.StokoutDate, sto.StokoutDate,
		sto.Remark, sto.Remark, sto.Function, sto.Function, sto.Status, sto.Status,
		sto.CreatedDate, sto.CreatedDate, sto.Page*sto.PageSize, sto.PageSize)

	return err
}

func (self *Database) UpdateStockoutByPpNo(ctx context.Context, so *entity.Stockout) error {
	var err error
	update_storage := `update pp_stockout_order set `
	var update_temp string
	var Id int64
	if len(so.Function) > 0 {
		update_temp = update_storage + `function =? where id = ?`
		result, err := self.Exec(update_temp, so.Function, so.Id)
		if err == nil {
			Id, err = result.RowsAffected()

		}
	}

	if len(so.Status) > 0 {
		update_temp = update_storage + `status=? where id = ?`
		result, err := self.Exec(update_temp, so.Status, so.Id)
		if err == nil {
			Id, err = result.RowsAffected()

		}
	}
	if len(so.Remark) > 0 {
		update_temp = update_storage + `remark=? where  id = ?`
		result, err := self.Exec(update_temp, so.Remark, so.Id)
		if err == nil {
			Id, err = result.RowsAffected()

		}
	}
	if len(so.StokoutDate) > 0 {
		update_temp = update_storage + `stokout_date=? where id = ?`
		result, err := self.Exec(update_temp, so.StokoutDate, so.Id)
		if err == nil {
			so.Id, err = result.RowsAffected()

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
