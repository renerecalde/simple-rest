package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type user struct {
	Id string `json:"user_id"`
	Name string `json:"name"`
}

func closeDb(db *sql.DB) {
	err := db.Close()
	if err != nil {
		panic(err)
	}
}

func openDb() (db *sql.DB)  {
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

func insertDb(usuario user, db *sql.DB)  {
	consulta := "INSERT INTO user (name) VALUES ('"+usuario.Name+"')"
	insert,err:= db.Query(consulta)
	if err != nil {
		panic(err)
	}

	defer insert.Close()

	_,err = fmt.Println("Insert exitoso")
	if err != nil {
		panic(err)
	}

}

func main()  {
	db:= openDb()
	defer closeDb(db)

	usuario:= user{
		Name: "rene1",
	}

	insertDb(usuario,db)




}
