package Controller

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/renerecalde/simple-rest/Entity"
	"github.com/renerecalde/simple-rest/ORM"
	"github.com/renerecalde/simple-rest/Repository"
	"io/ioutil"
	"net/http"
	"strconv"
)


func OptionsCooperativaHandler(w http.ResponseWriter, r *http.Request)  {
	if r.Method != http.MethodOptions {
		w.WriteHeader(http.StatusMethodNotAllowed)
		_, err := fmt.Fprintf(w, "405 - Method Not Allowed")
		if err != nil {
			http.Error(w, err.Error(), 500)
		}
		return
	}


	var methodsList = []string {
		http.MethodOptions,
		http.MethodPatch,
		http.MethodPut,
		http.MethodGet,
		http.MethodDelete,
		http.MethodPost,
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	err:= json.NewEncoder(w).Encode(methodsList)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
}


func UpdatePatchCooperativaHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPatch {
		w.WriteHeader(http.StatusMethodNotAllowed)
		_,err:=fmt.Fprintf(w, "405 - Method Not Allowed")
		if err != nil {
			http.Error(w, err.Error(), 500)
		}
		return
	}

	vars := mux.Vars(r)
	coopertivaId, err := strconv.Atoi(vars["id"])

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_,err:= fmt.Fprintf(w, "Id inválido")
		if err != nil {
			http.Error(w, err.Error(), 500)
		}
		return
	}

	var updateCooperativa Entity.Cooperativa
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_,err:=fmt.Fprintf(w, "Inserte datos válidos")
		if err != nil {
			http.Error(w, err.Error(), 500)
		}
		return
	}

	updateCooperativa.ID = coopertivaId

	err = json.Unmarshal(reqBody, &updateCooperativa)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}


	DB := ORM.OpenDb()

	if updateCooperativa.MatriculaNacional != 0 {
		_, err = Repository.UpdateCooperativaBy(
			"MatriculaNacional",
			strconv.Itoa(int(updateCooperativa.MatriculaNacional)),
			updateCooperativa.ID,
			DB)
		if err != nil {
			//TODO handle error
		}
	}

	if updateCooperativa.RazonSocial != "" {
		_, err = Repository.UpdateCooperativaBy(
			"RazonSocial",
			updateCooperativa.RazonSocial,
			updateCooperativa.ID,
			DB)
		if err != nil {
			//TODO handle error
		}
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	err= json.NewEncoder(w).Encode(updateCooperativa)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}


	//for i, cooperativa := range cooperativasList {
	//	if cooperativa.ID == coopertivaId {
	//
	//		if updateCooperativa.MatriculaNacional != 0 {
	//			cooperativasList[i].MatriculaNacional = updateCooperativa.MatriculaNacional
	//		}
	//
	//		if updateCooperativa.RazonSocial != "" {
	//			cooperativasList[i].RazonSocial = updateCooperativa.RazonSocial
	//		}
	//
	//		w.Header().Set("Content-type", "application/json")
	//		w.WriteHeader(http.StatusOK)
	//		err:= json.NewEncoder(w).Encode(cooperativasList[i])
	//		if err != nil {
	//			http.Error(w, err.Error(), 500)
	//		}
	//	}
	//}
}


func UpdateCooperativaHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		w.WriteHeader(http.StatusMethodNotAllowed)
		_,err:=fmt.Fprintf(w, "405 - Method Not Allowed")
		if err != nil {
			http.Error(w, err.Error(), 500)
		}
		return
	}
	vars := mux.Vars(r)
	coopertivaId, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_,err:=fmt.Fprintf(w, "Id inválido")
		if err != nil {
			http.Error(w, err.Error(), 500)
		}
		return
	}

	var updateCooperativa Entity.Cooperativa
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_,err:= fmt.Fprintf(w, "Inserte datos válidos")
		if err != nil {
			http.Error(w, err.Error(), 500)
		}
		return
	}

	err = json.Unmarshal(reqBody, &updateCooperativa)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	updateCooperativa.ID = coopertivaId

	DB := ORM.OpenDb()
	_, err = Repository.UpdateCooperativa(updateCooperativa, DB)
	if err != nil {
		//TODO handle error
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	err=json.NewEncoder(w).Encode(updateCooperativa)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}

}

func DeleteCooperativaHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		w.WriteHeader(http.StatusMethodNotAllowed)
		_,err:=fmt.Fprintf(w, "405 - Method Not Allowed")
		if err != nil {
			http.Error(w, err.Error(), 500)
		}
		return
	}


	vars := mux.Vars(r)
	coopertivaId, err := strconv.Atoi(vars["id"])


	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_,err:=fmt.Fprintf(w, "Id inválido")
		if err != nil {

			http.Error(w, err.Error(), 500)
		}
		return
	}

	DB := ORM.OpenDb()
	_, err = Repository.DeleteCooperativa(coopertivaId,DB)
	if err != nil {
		//TODO handle error
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	err=json.NewEncoder(w).Encode(coopertivaId)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}

}

func GetCooperativaHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		_, err := fmt.Fprintf(w, "405 - Method Not Allowed")
		if err != nil {

			http.Error(w, err.Error(), 500)
		}
		return
	}

	vars := mux.Vars(r)
	coopertivaId, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, err := fmt.Fprintf(w, "Id inválido")
		if err != nil {
			http.Error(w, err.Error(), 500)
		}
		return
	}

	DB := ORM.OpenDb()
	cooperativa, err := Repository.ReadCooperativa(coopertivaId,DB)
	if err != nil {
		//TODO handle error
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(cooperativa)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
}

func CreateCooperativaHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		_,err:=fmt.Fprintf(w, "405 - Method Not Allowed")
		if err != nil {
			http.Error(w, err.Error(), 500)
		}
		return
	}

	var newCooperativa Entity.Cooperativa
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_,err:=fmt.Fprintf(w, "Inserte datos válidos")
		if err != nil {

			http.Error(w, err.Error(), 500)
		}
		return
	}

	err=json.Unmarshal(reqBody, &newCooperativa)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	DB := ORM.OpenDb()
	_, err = Repository.CreateCooperativa(newCooperativa,DB)
	if err != nil {
		http.Error(w, err.Error(), 409)
	}


	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)

	//Acá debería devolver el id de la entidad que cree recientemente
	err=json.NewEncoder(w).Encode(newCooperativa)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
}



func GetCooperativasHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		_,err:=fmt.Fprintf(w, "405 - Method Not Allowed")
		if err != nil {
			http.Error(w, err.Error(), 500)
		}
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)

	DB := ORM.OpenDb()
	cooperativasList := Repository.ReadCooperativas(DB)

	err:=json.NewEncoder(w).Encode(cooperativasList)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
}
