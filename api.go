package main

import (
	"encoding/json"
	"fmt"

	"github.com/cho61183/api/mysql"
	"github.com/cho61183/api/pool"
)

type Response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {
	var mysqlconn = newMysqlConn()
	mysqlconn.Select()

	// 这里是将切片和字典编码为JSON数组或对象
	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := &Response2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0])
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
