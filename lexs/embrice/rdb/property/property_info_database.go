package property

import (
	"context"
	"fmt"
	entity "lexs/embrice/entity/property"
)

func (self *Database) AddProperty(ctx context.Context, po *entity.Property) error {
	query := `INSERT INTO pp_property(property_no, property_name, model,size,
	color,brand,prop_unit,unt_price,status,suplr_id,deleted, created_id,created_date)
	 VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?)`
	result, err := self.Exec(query, po.PropertyNo, po.PropertyName, po.Model, po.Size,
		po.Color, po.Brand, po.PropUnit, po.UntPrice, po.Status, po.SuplrId, po.Deleted,
		po.CreatedId, po.CreatedDate)
	if err == nil {
		po.Id, err = result.LastInsertId()
	}
	return err
}

func (self *Database) UpdatePropertyByNo(ctx context.Context, po *entity.Property) error {

	var err error
	update_property := `update pp_property set  `
	var update_temp string
	var Id int64
	if len(po.PropertyName) > 0 {
		update_temp = update_property + `property_name=? where property_no=? `
		result, err := self.Exec(update_temp, po.PropertyName, po.PropertyNo)
		if err == nil {
			Id, err = result.RowsAffected()

		}
	}
	if len(po.Model) > 0 {
		update_temp = update_property + `model=? where property_no=? `

		result, err := self.Exec(update_temp, po.Model, po.PropertyNo)
		if err == nil {
			Id, err = result.RowsAffected()

		}
	}
	if len(po.Color) > 0 {
		update_temp = update_property + `color=? where property_no=? `
		result, err := self.Exec(update_temp, po.Color, po.PropertyNo)
		if err == nil {
			Id, err = result.RowsAffected()

		}
	}
	if len(po.Brand) > 0 {
		update_temp = update_property + `brand=? where property_no=? `
		result, err := self.Exec(update_temp, po.Brand, po.PropertyNo)
		if err == nil {
			Id, err = result.RowsAffected()

		}
	}
	if len(po.Size) > 0 {
		update_temp = update_property + `size=? where property_no=? `
		result, err := self.Exec(update_temp, po.Size, po.PropertyNo)
		if err == nil {
			Id, err = result.RowsAffected()

		}
	}
	if len(po.Status) > 0 {
		update_temp = update_property + `status=? where property_no=? `
		result, err := self.Exec(update_temp, po.Status, po.PropertyNo)
		if err == nil {
			Id, err = result.RowsAffected()

		}
	}
	if len(po.PropUnit) > 0 {
		update_temp = update_property + `prop_unit=? where property_no=? `
		result, err := self.Exec(update_temp, po.PropUnit, po.PropertyNo)
		if err == nil {
			Id, err = result.RowsAffected()

		}
	}
	if po.UntPrice > 0 {
		update_temp = update_property + `unt_price=? where property_no=? `
		result, err := self.Exec(update_temp, po.UntPrice, po.PropertyNo)
		if err == nil {
			Id, err = result.RowsAffected()

		}
	}
	if len(po.Remark) > 0 {
		update_temp = update_property + `remark=? where property_no=? `
		result, err := self.Exec(update_temp, po.Remark, po.PropertyNo)
		if err == nil {
			Id, err = result.RowsAffected()

		}
	}

	if po.Deleted > 0 {
		update_temp = update_property + `deleted=? where property_no=?`
		result, err := self.Exec(update_temp, po.Deleted, po.PropertyNo)
		if err == nil {
			Id, err = result.RowsAffected()

		}
	}
	if Id > 0 {
		po.Id = Id
	}
	return err

}

func (self *Database) GetPropertyByNo(ctx context.Context, po *entity.Property) error {
	query_property := `select property_no,property_name,model,brand,
	color,size,prop_unit,status,deleted, created_id,created_date from pp_property where deleted= 0 and property_no=? `
	err := self.Get(po, query_property, po.PropertyNo)
	if err != nil {
		fmt.Println(err)
	}
	return err
}

func (self *Database) GetPropertyByName(ctx context.Context, pos *[]entity.Property, po *entity.Property) error {
	query_property := `select property_no,property_name,model,brand,
	color,size,prop_unit,status,deleted, created_id,created_date from pp_property 
	where deleted= 0 and property_name =?"`
	err := self.Get(pos, query_property, po.PropertyName)
	if err != nil {
		fmt.Println(err)
	}
	return err
}
