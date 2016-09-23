package main

import (
	"github.com/cho61183/api/mysql"
	"github.com/cho61183/api/pool"
)

type DbInfo struct {
	Host   string `json:"host"`
	Port   int    `json:"port"`
	User   string `json:"user"`
	Pass   string `json:"pass"`
	Dbname string `json:"dbname"`
	Num    int    `json:"num"`
}

func main() {
	var mysqlconn = newMysqlConn()
	mysqlconn.Select()
}

func newMysqlConn() *mysql.DB {
	var pools = &pool.Pool{
		MaxActive: 10,
		Dial: func() (interface{}, error) {
			conn, err := mysql.MysqlConn()
			return conn, err
		},
	}
	return pools.Get().(*mysql.DB)
}
