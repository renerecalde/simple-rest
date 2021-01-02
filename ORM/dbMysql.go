package ORM

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/renerecalde/simple-rest/Entity"
	"log"
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

func CreateSchmema(db *sql.DB) (result bool, err error)  {
	statement, err := db.Prepare(Entity.TableCooperativa)
	if err != nil {
		log.Println(err)
		result = false

	}
	_, err = statement.Exec()
	if err != nil {
		log.Println("Table already exists!")
	}
	statement, err = db.Prepare(Entity.TableUser)
	if err != nil {
		log.Println(err)
	}
	_, err = statement.Exec()
	if err != nil {
		log.Println("Table already exists!")
	}

	log.Println("All tables created/initialized successfully!")
	result = true

	return result, err
}

