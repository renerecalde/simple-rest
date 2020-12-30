package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/renerecalde/simple-rest/Entity"
	"github.com/renerecalde/simple-rest/ORM"
	"github.com/renerecalde/simple-rest/Repository"
)


func insertDb(usuario Entity.User, db *sql.DB)  {
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
	db:= ORM.OpenDb()
	defer ORM.CloseDb(db)

	usuario:= Entity.User{
		Name: "rene1",
	}

	cooperativa:= Entity.Cooperativa{
		RazonSocial: "Cooperativa 2",
		MatriculaNacional: 2,
	}

	insertDb(usuario,db)

	_,err:=Repository.CreateCooperativa(cooperativa, db)
	if err != nil {
		panic(err)
	}




}
