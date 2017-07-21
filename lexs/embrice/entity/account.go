package entity

type Account struct {
	ID       int64  `db:"id" json:"id"`
	Account  string `db:"account" json:"account"`
	Password string `db:"password" json:"password"`
	Mobile   string `db:"mobile" json:"mobile"`
	Status   int    `db:"status" json:"status"`
	Addr     string `db:"addr" json:"addr"`
	Email    string `db:"email" json:"email"`
	IP       string `db:"ip"`
}

type User struct {
	ID       int64   `db:"id" json:"id"`
	Account  string  `db:"account" json:"account"`
	Avatar   string  `db:"avatar" json:"avatar"`
	Nickname string  `db:"nickname" json:"nickname"`
	Money    float64 `db:"money" json:"money"`
	Coin     float64 `db:"coin" json:"coin"`
	Level    int     `db:"level" json:"level"`
	Rights   int     `db:"rights" json:"rights"`
	Vip      int     `db:"vip" json:"vip"`
	Point    int     `db:"point" json:"point"`
}
