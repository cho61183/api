package mysql

import (
	"database/sql"
	"encoding/json"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	//	"reflect"
)

//"database/sql"
//"fmt"

//	"_github.com/go-sql-driver/mysql"

type DbInfo struct {
	Host   string `json:"host"`
	Port   int    `json:"port"`
	User   string `json:"user"`
	Pass   string `json:"pass"`
	Dbname string `json:"dbname"`
	Num    int    `json:"num"`
}

type DB struct {
	Db_conn *sql.DB
}

type User struct {
	user_id          int
	user_name, phone string
}

func init() {

}

func MysqlConn() (*DB, error) {

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
       "num" : 2
   	}]`
	Conn_str := []DbInfo{}
	err := json.Unmarshal([]byte(str), &Conn_str)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(Conn_str[0].User)
	Conns := fmt.Sprintf("%s:%s@/%s?charset=utf8", Conn_str[0].User, Conn_str[0].Pass, Conn_str[0].Dbname)
	dbs, _ := sql.Open("mysql", Conns)
	db_drive := &DB{Db_conn: dbs}
	return db_drive, nil
}

func (this *DB) GetPort() int {
	return 633
}

/*
查询方法
*/
func (this *DB) Select() {
	var u User
	//查询数据
	rows, err := this.Db_conn.Query("SELECT user_id,user_name,phone FROM t_user where user_id=66")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(rows.Columns())
	userinfo := make(map[interface{}]interface{})
	for rows.Next() {
		err := rows.Scan(&u.user_id, &u.user_name, &u.phone)
		if err != nil {
			fmt.Println(err)
		}
		userinfo[u.user_id] = u
	}
	fmt.Println(userinfo)
}

/*
删除方法
*/
func (this *DB) Delete() {
	fmt.Println("this is Delete")
}

/*
插入方法
*/
func (this *DB) Insert() {
	fmt.Println("this is Insert")
}

/*
更新方法
*/
func (this *DB) Update() {
	fmt.Println("this is Update")
}
