package ORM

import (
	"database/sql"
	"fmt"
)

func CloseDb(db *sql.DB) {
	err := db.Close()
	if err != nil {
		panic(err)
	}
}

func OpenDb() (db *sql.DB)  {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3307)/cooperativa")
	if err != nil {
		panic(err)
	}
	_,err= fmt.Println("Conexión exitosa")
	if err != nil {
		panic(err)
	}
	return db
}

