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
		panic(err)
	}
	log.Println("Inserted the coop into database!")

	return result, err
}

func ReadCooperativa(id int, db *sql.DB)   (cooperativa Entity.Cooperativa, err error){
	err = db.QueryRow("select ID, RazonSocial, MatriculaNacional from cooperativa where ID = ?", id).Scan(&cooperativa.ID, &cooperativa.RazonSocial, &cooperativa.MatriculaNacional)
	if err != nil {
		log.Fatal(err)
	}
	return cooperativa, err

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

	err:= rows.Close()
	if err != nil {
		//TODO handle error
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

func UpdateCooperativaBy(field string, value string, id int, db *sql.DB) (result sql.Result, err error){
	statement, _ := db.Prepare("update cooperativa set ?=? where id=?")
	result, err = statement.Exec(field,value,strconv.Itoa(id))
	if err != nil {
		//TODO handle error
	}
	log.Println("Successfully updated the book in database!")

	return result,err
}


func DeleteCooperativa(id int, db *sql.DB)  (result sql.Result, err error){
	statement, _ := db.Prepare("delete from cooperativa where id=?")
	result, err = statement.Exec(id)
	if err != nil {
		//TODO handle error
	}
	log.Println("Successfully deleted the book in database!")

	return result, err
}