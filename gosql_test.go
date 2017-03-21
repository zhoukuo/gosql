package gosql

import (
	"fmt"
	"testing"
)

var schema = `
	CREATE TABLE "userinfo" (
	    "uid" INTEGER PRIMARY KEY AUTOINCREMENT,
	    "username" VARCHAR(64) NULL,
	    "departname" VARCHAR(64) NULL,
	    "created" VARCHAR(64) NULL
	);
`

func TestOpenAndClose(t *testing.T) {
	db, _ := Open("sqlite3", "test.db")
	defer Close(db)
}

func TestCreate(t *testing.T) {
	db, _ := Open("sqlite3", "test.db")
	defer Close(db)

	Create(db, schema)
}

func TestInsert(t *testing.T) {
	db, _ := Open("sqlite3", "test.db")
	defer Close(db)

	Insert(db, "INSERT INTO userinfo(username, departname, created) values(?,?,?)", "zhangsan", "dev", "2017-1-1")
	Insert(db, "INSERT INTO userinfo(username, departname, created) values(?,?,?)", "lisi", "dev", "2017-2-1")
}

func TestUpdate(t *testing.T) {
	db, _ := Open("sqlite3", "test.db")
	defer Close(db)

	Update(db, "UPDATE userinfo SET departname=? WHERE username=?", "dev2", "lisi")
}

func TestQuery(t *testing.T) {
	db, _ := Open("sqlite3", "test.db")
	defer Close(db)

	rows, _ := Query(db, "SELECT * FROM userinfo")
	for _, row := range *rows {
		fmt.Println(row)
	}

	rows, _ = Query(db, "SELECT * FROM userinfo WHERE USERNAME='lisi'")
	for _, row := range *rows {
		fmt.Println(row)
	}
}

func TestDelete(t *testing.T) {
	db, _ := Open("sqlite3", "test.db")
	defer Close(db)

	Delete(db, "DELETE FROM userinfo")
}

func TestDrop(t *testing.T) {
	db, _ := Open("sqlite3", "test.db")
	defer Close(db)

	Drop(db, "DROP TABLE userinfo")
}
