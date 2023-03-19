package main

import (
	"fmt"
	"time"

	"github.com/beego/beego/v2/client/orm"
	_ "github.com/go-sql-driver/mysql"
)

// User model for users table
type User struct {
	Id    int    // Column id
	Name  string // Column name
	Email string // Column email
}

func (u *User) TableName() string {
	return "users"
}

// Order model for orders table
type Order struct {
	Id        int       // Column id
	UserId    int       // Column user_id
	OrderTime time.Time // Column order_time
	Price     int       // Column price
}

func (u *Order) TableName() string {
	return "orders"
}

type QueryResp struct {
	Name      string
	OrderTime time.Time
	Price     int
}

func main() {
	orm.RegisterDataBase("default", "mysql", "beego_qb:tmp_pwd@tcp(127.0.0.1:3306)/beego_qb")
	orm.RegisterModel(new(User), new(Order))

	var rows []QueryResp

	// Get a QueryBuilder object.
	qb, _ := orm.NewQueryBuilder("mysql")

	// Construct query object
	qb.Select("users.name",
		"orders.order_time", "orders.price").
		From("users").
		InnerJoin("orders").On("users.id = orders.user_id").
		Where("orders.price").
		OrderBy("orders.price").Desc().
		Limit(10)

	// export raw query string from QueryBuilder
	sql := qb.String()

	fmt.Println(sql)

	// execute the raw query
	o := orm.NewOrm()
	o.Raw(sql).QueryRows(&rows)

	fmt.Printf("%v\n\n", rows)
}
