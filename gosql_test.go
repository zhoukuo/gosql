package gosql

import (
	// "fmt"
	"os"
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

	id, _ := Insert(db, "INSERT INTO userinfo(username, departname, created) values(?,?,?)", "zhangsan", "dev", "2017-1-1")
	if id != 1 {
		t.Errorf("expected: id=1, actually: id=%d", id)
	}

	id, _ = Insert(db, "INSERT INTO userinfo(username, departname, created) values(?,?,?)", "lisi", "dev", "2017-2-1")
	if id != 2 {
		t.Errorf("expected: id=1, actually: id=%d", id)
	}
}

func TestUpdate(t *testing.T) {
	db, _ := Open("sqlite3", "test.db")
	defer Close(db)

	affected, _ := Update(db, "UPDATE userinfo SET departname=? WHERE username=?", "dev2", "lisi")
	if affected != 1 {
		t.Errorf("expected: affected=1, actually: affected=%d", affected)
	}
}

func TestQuery(t *testing.T) {
	db, _ := Open("sqlite3", "test.db")
	defer Close(db)

	rows, _ := Query(db, "SELECT * FROM userinfo")
	var count int
	for i, _ := range *rows {
		// fmt.Printf("%d: ", i)
		// fmt.Println(row)
		count = i + 1
	}

	if count != 2 {
		t.Errorf("expected: rows.length=2, actually: rows.length=%d", count)
	}

	rows, _ = Query(db, "SELECT * FROM userinfo WHERE USERNAME='lisi'")
	for i, _ := range *rows {
		// fmt.Println(row)
		count = i + 1
	}
	if count != 1 {
		t.Errorf("expected: rows.length=1, actually: rows.length=%d", count)
	}
}

func TestDelete(t *testing.T) {
	db, _ := Open("sqlite3", "test.db")
	defer Close(db)

	affected, _ := Delete(db, "DELETE FROM userinfo")

	if affected != 2 {
		t.Errorf("expected: affected=2, actually: affected=%d", affected)
	}
}

func TestDrop(t *testing.T) {
	db, _ := Open("sqlite3", "test.db")
	defer Close(db)

	Drop(db, "DROP TABLE userinfo")
	os.Remove("test.db")
}
