package Repository

import (
	"database/sql"
	"github.com/renerecalde/simple-rest/Entity"
	"log"
	"strconv"
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

func ReadCooperativas(db *sql.DB) (cooperativaList Entity.Cooperativas) {
	rows, _ := db.Query("SELECT ID, RazonSocial, MatriculaNacional FROM cooperativa")
	var tempCooperativa Entity.Cooperativa
	for rows.Next() {
		err:= rows.Scan(&tempCooperativa.ID, &tempCooperativa.RazonSocial, &tempCooperativa.MatriculaNacional)
		if err != nil {
			//TODO handle error
		}
		log.Printf("ID:%d, RazonSocial:%s, MatriculaNacional:%s\n", tempCooperativa.ID, tempCooperativa.RazonSocial, strconv.Itoa(int(tempCooperativa.MatriculaNacional)))
		cooperativaList = append(cooperativaList, tempCooperativa)
	}
	return cooperativaList
}

func UpdateCooperativa(cooperativa Entity.Cooperativa, db *sql.DB) (result sql.Result, err error){
	statement, _ := db.Prepare("update cooperativa set RazonSocial=?, MatriculaNacional=? where id=?")
	result, err = statement.Exec(cooperativa.RazonSocial, cooperativa.MatriculaNacional, cooperativa.ID)
	if err != nil {
		//TODO handle error
	}
	log.Println("Successfully updated the book in database!")

	return result,err
}

func DeleteCooperativa(cooperativa Entity.Cooperativa, db *sql.DB)  (result sql.Result, err error){
	statement, _ := db.Prepare("delete from cooperativa where id=?")
	result, err = statement.Exec(cooperativa.ID)
	if err != nil {
		//TODO handle error
	}
	log.Println("Successfully deleted the book in database!")

	return result, err
}