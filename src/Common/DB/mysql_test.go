package DB

import (
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

func TestSqlConnection(t *testing.T) {
	db, err := NewConn()
	if err != nil {
		t.Error("failure coon mysql ", err)
		return
	}
	row, err := db.Query("select count(*) from account")

	if err != nil {
		t.Error("failure get count ", err)
		return
	}
	for row.Next() {
		var count string
		if err := row.Scan(&count); err != nil {
			t.Error("failure scan ", err)
			return
		}
		t.Log("results", count)
	}

}
