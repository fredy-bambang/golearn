package main

import (
	"time"

	"github.com/fredy-bambang/golearn/http/routes"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import your used driver
)

func init() {
	// register model
	orm.RegisterModel(new(Profile))

	// set default database
	maxIdle := 30
	maxConn := 30
	orm.DefaultTimeLoc = time.UTC
	orm.RegisterDataBase("default", "mysql", "root:@/laravel6?charset=utf8", maxIdle, maxConn)

	// create table
	// orm.RunSyncdb("default", false, true)
}

// Profile .
type Profile struct {
	ID      int    `orm:"auto;column(id)"`
	Name    string `orm:"size(100);unique"`
	Address string `orm:"size(100)"`
}

func main() {
	routes.HandleRequests()
	// o := orm.NewOrm()

	// user := Profile{Name: "donal"}

	// // insert
	// id, err := o.Insert(&user)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(id)

	// // update
	// // user := Profile{}
	// user.Name = "astaxie"
	// _, err = o.Update(&user)

	// // read one
	// u := Profile{ID: user.ID}
	// err = o.Read(&u)

	// // delete
	// _, err = o.Delete(&u)
	// fmt.Println("done")

}
