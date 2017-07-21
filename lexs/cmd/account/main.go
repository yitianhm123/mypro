package main

import (
	"context"
	"lexs/embrice/rdb/account"
	"lexs/x/logger"
	"lexs/x/web"
)

type server struct {
	*web.Server
	db  *account.Database
	log *logger.Logger
	ctx context.Context
}

// start the server
func (self *server) start() {
	self.Server = web.NewServer(self.log)
	router := self.PathPrefix("/account/").SubRouter()
	//self.Filter(middleware.Authenticate).FilterFunc(middleware.Check)
	router.HandleFunc("/regist", self.regist).Methods("POST")
	router.HandleFunc("/login", self.login).Methods("POST")
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
		db:  account.NewDB(log, 10, 10),
		ctx: context.TODO(),
		log: log,
	}

	s.start()
	defer s.stop()
}
