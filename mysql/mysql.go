package mysql

import (
	//	"database/sql"
	"fmt"
)

//"database/sql"
//"fmt"

//	"_github.com/go-sql-driver/mysql"

type Conn struct {
	Driver string
	Port   int
	Ip     string
	User   string
	Pass   string
	Db     string
}

func init() {

}

func MysqlConn() (*Conn, error) {
	//	var Conn = fmt.Sprint("%s:%s@%s/?charset=utf8", this.User, this.Pass, this.Db)
	//db, err := sql.Open("mysql", "root:123456@yoyojie_user/?charset=utf8")
	db := &Conn{Port: 6339}
	return db, nil
}

func (this *Conn) GetPort() int {
	return this.Port
}

/*
查询方法
*/
func (this *Conn) Select() {
	fmt.Println("this is Select")
}

/*
删除方法
*/
func (this *Conn) Delete() {
	fmt.Println("this is Delete")
}

/*
插入方法
*/
func (this *Conn) Insert() {
	fmt.Println("this is Insert")
}

/*
更新方法
*/
func (this *Conn) Update() {
	fmt.Println("this is Update")
}
