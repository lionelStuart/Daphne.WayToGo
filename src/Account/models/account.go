package models

import (
	"Common/DB"
	"database/sql"
	"fmt"
	"github.com/astaxie/beego/logs"
	"time"
)

type Account struct {
	Uid        int16     `db:"uid"`
	Uname      string    `db:"uname"`
	Email      string    `db:"email"`
	Password   string    `db:"password"`
	Phone      string    `db:"phone"`
	Auth       int16     `db:"auth"`
	CreateTime time.Time `db:"create_time"`
	UpdateTime time.Time `db:"update_time"`
}

var (
	db  *sql.DB
	err error
)

func init() {
	db, err = DB.NewConn()
	if err != nil {
		logs.Error("failure init db")
		panic(err)
	}
}

func (account *Account) Create() {
	stmt, err := db.Prepare("insert into account(uname,email,`password`,phone,auth) VALUES(?,?,?,?,?)")
	if err != nil {
		panic(err)
	}
	res, err := stmt.Exec(account.Uname, account.Email, account.Password, account.Phone, account.Auth)
	if err != nil {
		panic(err)
	}
	id, err := res.LastInsertId()
	if err != nil {
		panic(err)
	} else {
		fmt.Println("auto increment id ", id)
	}

}

func (account *Account) QueryByName() {
	stmt, err := db.Prepare("select uid,email,'password',phone,auth from account where uname=?")
	if err != nil {
		panic(err)
	}
	row, err := stmt.Query(account.Uname)
	if err != nil {
		panic(err)
	}

	for row.Next() {
		err := row.Scan(&account.Uid, &account.Email, &account.Password, &account.Phone, &account.Auth)
		if err != nil {
			panic(err)
		} else {
			fmt.Print("success get account id ", account.Uid)
			break
		}
	}
}

func (account *Account) Update() {
	stmt, err := db.Prepare("update account set email=?,`password`=?,phone=?,auth=? where uid=?")
	if err != nil {
		panic(err)
	}
	res, err := stmt.Exec(account.Email, account.Password, account.Phone, account.Auth, account.Uid)
	if err != nil {
		panic(err)
	} else {
		affect, _ := res.RowsAffected()
		fmt.Println("success execute ", affect)
	}

}
