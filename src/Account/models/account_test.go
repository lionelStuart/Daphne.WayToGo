package models

import (
	"Common/DB"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

func TestAccount(t *testing.T) {
	db, err := DB.NewConn()
	defer func() {
		db.Close()
	}()

	if err != nil {
		t.Error("failure coon mysql ", err)
		return
	}

	rows, err := db.Query("select uid,uname,update_time from account")

	for rows.Next() {
		account := Account{}
		err = rows.Scan(&account.Uid, &account.Uname, &account.UpdateTime)

		if err != nil {
			t.Error("failure scan row ", err)
			return
		}
		fmt.Printf("receive account: %d %s %s\n", account.Uid, account.Uname, account.UpdateTime.Format("2006-01-02 15:04:05"))

	}

}

func TestInsert(t *testing.T) {
	account := Account{
		Uname:    "jim",
		Email:    "jim@123.com",
		Password: "123456",
		Phone:    "12312345678",
		Auth:     1,
	}
	account.Create()

}

func TestAccount_QueryByName(t *testing.T) {
	accout := Account{
		Uname: "jim",
	}
	accout.QueryByName()
	t.Log("get account :", accout)
}

func TestAccount_Update(t *testing.T) {
	account := Account{
		Uname: "jim",
	}
	account.QueryByName()
	account.Email = "jim@1234.com"
	account.Update()
}
