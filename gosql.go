package gosql

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func Create(db *sql.DB, schema string) {
	// create table
	stmt, err := db.Prepare(schema)
	checkErr(err)

	defer stmt.Close()

	_, err = stmt.Exec()
	checkErr(err)
}

func Drop(db *sql.DB, schema string) {
	// drop table
	stmt, err := db.Prepare(schema)
	checkErr(err)

	defer stmt.Close()

	_, err = stmt.Exec()
	checkErr(err)
}

func Open(driver string, source string) (*sql.DB, error) {
	// open db
	db, err := sql.Open(driver, source)
	checkErr(err)

	err = db.Ping()
	checkErr(err)

	return db, err
}

func Close(db *sql.DB) {
	// close db
	db.Close()
}

func Insert(db *sql.DB, sqlStatement string, args ...interface{}) (int64, error) {
	// insert record
	stmt, err := db.Prepare(sqlStatement)
	checkErr(err)

	defer stmt.Close()

	res, err := stmt.Exec(args...)
	checkErr(err)

	return res.LastInsertId()
}

func Query(db *sql.DB, sqlStatement string, args ...interface{}) (*[]map[string]string, error) {
	// query records
	stmtOut, err := db.Prepare(sqlStatement)
	if err != nil {
		panic(err.Error())
	}
	defer stmtOut.Close()

	rows, err := stmtOut.Query(args...)
	if err != nil {
		panic(err.Error())
	}

	columns, err := rows.Columns()
	if err != nil {
		panic(err.Error())
	}

	values := make([]sql.RawBytes, len(columns))
	scanArgs := make([]interface{}, len(values))

	ret := make([]map[string]string, 0)
	for i := range values {
		scanArgs[i] = &values[i]
	}

	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			panic(err.Error())
		}
		var value string
		vmap := make(map[string]string, len(scanArgs))
		for i, col := range values {
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			vmap[columns[i]] = value
		}
		ret = append(ret, vmap)
	}
	return &ret, nil
}

func Update(db *sql.DB, sqlStatement string, args ...interface{}) (int64, error) {
	// update columns
	stmt, err := db.Prepare(sqlStatement)
	checkErr(err)

	defer stmt.Close()

	res, err := stmt.Exec(args...)

	return res.RowsAffected()
}

func Delete(db *sql.DB, sqlStatement string, args ...interface{}) (int64, error) {
	// delete records
	stmt, err := db.Prepare(sqlStatement)
	checkErr(err)

	defer stmt.Close()

	res, err := stmt.Exec(args...)
	checkErr(err)
	return res.RowsAffected()
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
