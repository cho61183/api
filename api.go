package main

import (
	"encoding/json"
	"fmt"

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

	// 这里是将切片和字典编码为JSON数组或对象
	str := `[{
       "host": "127.0.0.1",
       "port": 3306,
       "user": "root",
       "pass": "123456",
       "dbname": "yoyojie_user",
       "num" : 1
   	},
	{
       "host": "127.0.0.1",
       "port": 3306,
       "user": "root",
       "pass": "123456",
       "dbname": "yoyojie_user",
       "num" : 1
   	}]`
	res := []DbInfo{}
	json.Unmarshal([]byte(str), &res)

	for _, v := range res {
		fmt.Println(v)
	}
	//fmt.Println(res.Dbname)
}

func newMysqlConn() *mysql.Conn {
	var pools = &pool.Pool{
		MaxActive: 10,
		Dial: func() (interface{}, error) {
			conn, err := mysql.MysqlConn()
			return conn, err
		},
	}
	return pools.Get().(*mysql.Conn)
}
