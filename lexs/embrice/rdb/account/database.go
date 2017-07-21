package account

import (
	"context"
	"lexs/embrice/entity"
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
		//DB:  sqlx.MustConnect("mysql", "devsql:devsql123.com@tcp(10.10.10.121:3306)/Rocket"),
		DB: sqlx.MustConnect("mysql", "root:123456@tcp(10.188.10.2:3306)/aaa_4"),
	}
	db.SetMaxOpenConns(maxConns)
	db.SetMaxIdleConns(maxIdles)
	return db
}

func (self *Database) Reg(ctx context.Context, ac *entity.Account) error {
	query := `INSERT INTO account(account, password, mobile, ip) VALUES (?,?,?,?)`
	result, err := self.Exec(query, ac.Account, ac.Password, ac.Mobile, ac.IP)
	if err == nil {
		ac.ID, err = result.LastInsertId()
	}
	return err
}

func (self *Database) Login(ctx context.Context, ac *entity.Account) error {
	return self.Get(ac, "SELECT id, password,mobile FROM account WHERE account =?", ac.Account)
}

func (self *Database) Close() {
	self.DB.Close()
}
