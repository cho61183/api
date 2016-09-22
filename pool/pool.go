/*
连接池的操作
*/
package pool

import (
	"fmt"
	"reflect"
	"time"
)

var nowFun = time.Now

type Pool struct {
	//这里保存一个操作redis/mysql操作的类
	Dial func() (interface{}, error)

	MaxActive int

	MaxIdle int

	idle chan interface{}
}

/*
定义一个操作连接
*/
type Conn struct {
	t time.Time
	c interface{}
}

/*
初始化连接池
*/
func (this *Pool) InitPool() error {
	this.idle = make(chan interface{}, this.MaxActive)
	for i := 0; i < this.MaxActive; i++ {
		conn, err := this.Dial()
		fmt.Println(reflect.TypeOf(conn))
		if err != nil {
			return err
		}
		this.idle <- Conn{t: nowFun(), c: conn}
	}
	return nil
}

/*
获得连接
*/
func (this *Pool) Get() interface{} {
	if this.idle == nil {
		this.InitPool()
	}
	c := <-this.idle
	conn := c.(Conn).c
	defer this.Release(conn)
	return conn
}

/*
连接池回收
*/
func (this *Pool) Release(conn interface{}) {
	this.idle <- Conn{t: nowFun(), c: conn}
}
