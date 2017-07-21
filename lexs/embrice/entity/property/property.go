package entity

type Purchase struct {
}

type BrokenOrder struct {
	Id            int64  `db:"id"               json:"id"`
	PropertyNo    string `db:"property_no"      json:"property_no"`
	BrokenOrderNo string `db:"broken_order_no"  json:"broken_order_no"`
	CreatedId     int64  `db:"created_id"       json:"created_id"`
	CreatedDate   string `db:"created_date"     json:"created_date"`
	Status        string `db:"status"           json:"status"`
	Remark        string `db:"remark"           json:"remark"`
	UpdateId      int64  `db:"updated_id"       json:"updated_id"`
	UpdateDate    string `db:"update_Date"      json:"update_Date"`
	Deleted       int    `db:"deleted"          json:"deleted"`
	Page          int    `json:"page"`
	PageSize      int    `json:"pagesize"`
	Total         int    `json:"total"`
}

type Response struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message"`
}

type Pagination struct {
	Page     int64 `json:"page"`
	PageSize int64 `json:"pagesize"`
	Total    int64 `json:"total"`
}

type Property struct {
	Id           int64   `db:"id"             json:"id"            `
	PropertyNo   string  `db:"property_no"    json:"property_no"   `
	PropertyName string  `db:"property_name"  json:"property_name" `
	Model        string  `db:"model"          json:"model"         `
	Size         string  `db:"size"           json:"size"          `
	Color        string  `db:"color"          json:"color"         `
	Brand        string  `db:"brand"          json:"brand"         `
	CreatedDate  string  `db:"created_date"   json:"created_date"  `
	CreatedId    int64   `db:"created_id"     json:"created_id"    `
	PropUnit     string  `db:"prop_unit"      json:"prop_unit"     `
	UntPrice     float32 `db:"unt_price"      json:"unt_price"     `
	Status       string  `db:"status"         json:"status"        `
	Remark       string  `db:"remark"         json:"remark"`
	SuplrId      int64   `db:"suplr_id"       json:"suplr_id"      `
	UpdatedId    int64   `db:"updated_id"     json:"updated_id"    `
	UpdatedDate  string  `db:"updated_date"   json:"updated_date"  `
	Deleted      int     `db:"deleted"        json:"deleted"       `
	Page         int     `json:"page"`
	PageSize     int     `json:"pagesize"`
	Total        int     `json:"total"`
}

type StorageOrder struct {
	Id           int64  `db:"id"           json:"id"          `
	StorageNo    string `db:"storage_no"   json:"storage_no"  `
	PropertyNo   string `db:"property_no"  json:"property_no" `
	PropertyName string `db:"property_name" json:"property_name"`
	Model        string `db:"model"        json:"model"       `
	Brand        string `db:"brand"        json:"brand"`
	Count        int    `db:"pcount"       json:"count"`
	Status       string `db:"status"       json:"status"`
	Color        string `db:"color"        json:"color"       `
	Size         string `db:"size"         json:"size"        `
	CreatedId    int64  `db:"created_id"   json:"created_id"  `
	CreatedDate  string `db:"created_date" json:"created_date"`
	Source       string `db:"source"       json:"source"      `
	Remark       string `db:"remark"       json:"remark"      `
	StorageDate  string `db:"storage_date" json:"storage_date"`
	UpdatedId    int64  `db:"updated_id"   json:"updated_id"  `
	UpdatedDate  string `db:"updated_date" json:"updated_date"`
	Deleted      int32  `db:"deleted"      json:"deleted"     `
	Page         int    `json:"page"`
	PageSize     int    `json:"pagesize"`
	Total        int    `json:"total"`
}

type Stockout struct {
	Id          int64  `db:"id"            json:"id"`
	PropertyNo  string `db:"property_no"   json:"property_no"`
	StockoutNo  string `db:"stockout_no"   json:"stockout_no"`
	CreatedId   int64  `db:"created_id"    json:"created_id"`
	CreatedDate string `db:"created_date"  json:"created_date"`
	StokoutDate string `db:"stokout_date"  json:"stokout_Date"`
	Function    string `db:"function"      json:"function"`
	Remark      string `db:"remark"        json:"remark"`
	Status      string `db:"status"        json:"status"`
	UpdateId    string `db:"updated_id"    json:"updated_id"`
	UpdateDate  string `db:"updated_date"  json:"updated_date"`
	Deleted     int    `db:"deleted"       json:"Deleted"`
	Page        int    `json:"page"`
	PageSize    int    `json:"pagesize"`
	Total       int    `json:"total"`
}

type PurchaseOrder struct {
	Id           int64  `db:"id"              json:"id"            `
	PorderNo     string `db:"porder_no"       json:"porder_no"     `
	SuplrId      int64  `db:"suplr_id"        json:"suplr_id"      `
	PropertyName string `db:"property_name"   json:"property_name" `
	Model        string `db:"model"           json:"model"         `
	Count        int    `db:"count"           json:"count"         `
	Status       string `db:"status"          json:"status"        `
	Remark       string `db:"remark"          json:"remark"`
	CreatedId    int64  `db:"created_id"      json:"created_id"    `
	CreatedDate  string `db:"created_date"    json:"created_date"  `
	Color        string `db:"color"           json:"color"         `
	Brand        string `db:"brand"           json:"brand"         `
	Size         string `db:"size"            json:"size"          `
	UpdatedId    int64  `db:"update_id"       json:"update_id"     `
	UpdatedDate  string `db:"updated_date"    json:"updated_date"  `
	Deleted      int    `db:"deleted"         json:"deleted"       `
	Page         int    `json:"page"`
	PageSize     int    `json:"pagesize"`
	Total        int    `json:"total"`
}
