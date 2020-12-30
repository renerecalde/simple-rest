package Repository

import (
	"database/sql"
	"github.com/renerecalde/simple-rest/Entity"
	"log"
)

func CreateCooperativa(cooperativa Entity.Cooperativa, db *sql.DB) (result sql.Result, err error)  {

	statement, _ := db.Prepare("INSERT INTO cooperativa (RazonSocial, MatriculaNacional) VALUES (?,?)")
	result, err = statement.Exec(cooperativa.RazonSocial, cooperativa.MatriculaNacional)
	if err != nil {
		//TODO handle error

	}
	log.Println("Inserted the coop into database!")

	return result, err
}
