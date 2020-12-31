package Controller

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/renerecalde/simple-rest/Entity"
	"io/ioutil"
	"net/http"
	"strconv"
)


var cooperativa = Entity.Cooperativa{
	ID:                1,
	RazonSocial:       "Cooperativa 1",
	MatriculaNacional: 1,
}
var cooperativasList = Entity.Cooperativas{
	cooperativa,
}
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

	err = json.Unmarshal(reqBody, &updateCooperativa)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	for i, cooperativa := range cooperativasList {
		if cooperativa.ID == coopertivaId {

			if updateCooperativa.MatriculaNacional != 0 {
				cooperativasList[i].MatriculaNacional = updateCooperativa.MatriculaNacional
			}

			if updateCooperativa.RazonSocial != "" {
				cooperativasList[i].RazonSocial = updateCooperativa.RazonSocial
			}

			w.Header().Set("Content-type", "application/json")
			w.WriteHeader(http.StatusOK)
			err:= json.NewEncoder(w).Encode(cooperativasList[i])
			if err != nil {
				http.Error(w, err.Error(), 500)
			}
		}
	}
}


func UpdateCooperativaHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
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

	for i, cooperativa := range cooperativasList {
		if cooperativa.ID == coopertivaId {

			cooperativasList = append(cooperativasList[:i], cooperativasList[i+1:]...)
			updateCooperativa.ID = coopertivaId

			cooperativasList = append(cooperativasList, updateCooperativa)

			w.Header().Set("Content-type", "application/json")
			w.WriteHeader(http.StatusOK)
			err:=json.NewEncoder(w).Encode(cooperativasList[i])
			if err != nil {
				http.Error(w, err.Error(), 500)
			}
		}
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

	for i, cooperativa := range cooperativasList {
		if cooperativa.ID == coopertivaId {

			cooperativasList = append(
				cooperativasList[:i],
				cooperativasList[i+1:]...,
			)

			w.Header().Set("Content-type", "application/json")
			w.WriteHeader(http.StatusOK)
			err:=json.NewEncoder(w).Encode(cooperativa)
			if err != nil {

				http.Error(w, err.Error(), 500)

			}
		}
	}
}

func GetCooperativaHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
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

	for _, cooperativa := range cooperativasList {
		if cooperativa.ID == coopertivaId {
			w.Header().Set("Content-type", "application/json")
			w.WriteHeader(http.StatusOK)
			err:=json.NewEncoder(w).Encode(cooperativa)
			if err != nil {

				http.Error(w, err.Error(), 500)
			}
		}
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
	newCooperativa.ID = len(cooperativasList) + 1
	cooperativasList = append(
		cooperativasList,
		newCooperativa,
	)


	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
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
	err:=json.NewEncoder(w).Encode(cooperativasList)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
}
