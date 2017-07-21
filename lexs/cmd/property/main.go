package main

import (
	"context"
	"lexs/embrice/rdb/property"
	"lexs/x/logger"
	"lexs/x/web"
)

type server struct {
	*web.Server
	db  *property.Database
	log *logger.Logger
	ctx context.Context
}

// start the server
func (self *server) start() {
	self.Server = web.NewServer(self.log)
	router := self.PathPrefix("/property/").SubRouter()
	//self.Filter(middleware.Authenticate).FilterFunc(middleware.Check)
	router.HandleFunc("/hello", self.hello).Methods("POST")
	//添加资产信息
	router.HandleFunc("/add_property", self.addProperty).Methods("POST")
	//通过资产编号获取资产信息
	router.HandleFunc("/get_property_by_no", self.getPropertyByNo).Methods("POST")
	//通过资产编号更新资产信息
	router.HandleFunc("/update_property_by_no", self.updatePropertyByNo).Methods("POST")
	//通过资产名称获取资产信息
	router.HandleFunc("/get_property_by_name", self.getPropertyByName).Methods("POST")
	//获取出库单
	router.HandleFunc("/get_stockout", self.getStockoutDetail).Methods("POST")
	//创建出库单
	router.HandleFunc("/add_stockout", self.addStockout).Methods("POST")
	//更新出库单
	router.HandleFunc("/update_stockout_by_id", self.updateStockout).Methods("POST")
	//获取报损单
	router.HandleFunc("/get_broken_order", self.getBrokenorder).Methods("POST")
	//创建报损单
	router.HandleFunc("/add_broken_order", self.addBrokenorder).Methods("POST")
	//更新报损单
	router.HandleFunc("/update_broken_order", self.updateBrokenOrderByPpno).Methods("POST")
	//创建入库单
	router.HandleFunc("/add_storage", self.addStorage).Methods("POST")
	//获取入库单列表
	router.HandleFunc("/get_storages", self.getStorages).Methods("POST")
	//获取入库单详情
	router.HandleFunc("/get_storage_detail", self.getStorageDetail).Methods("POST")
	//根据资产编号和入库单号更新入库单信息
	router.HandleFunc("/update_storage_by_ppno", self.updateStorageByPpno).Methods("POST")
	//创建采购单
	router.HandleFunc("/created_purchase_order", self.createPurchaseOrder).Methods("POST")
	//查询采购单详情
	router.HandleFunc("/get_purchase_order_detail", self.getPurchaseOrderDetail).Methods("POST")
	//更新采购单
	router.HandleFunc("/update_purchase_order_by_id", self.updatePurchaseOrderById).Methods("POST")
	self.Serve(":8001")
}

// stop the server
func (self *server) stop() {
	self.db.Close()
}

func main() {
	// todo: load conf later
	log := logger.NewStdLogger(true, true, true, true, true)
	s := &server{
		db:  property.NewDB(log, 10, 10),
		ctx: context.TODO(),
		log: log,
	}

	s.start()
	defer s.stop()
}
