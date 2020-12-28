package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

type cooperativa struct {
	ID                int    `json:"ID"`
	RazonSocial       string `json:"RazonSocial"`
	MatriculaNacional int8   `json:"MatriculaNacional"`
}

type cooperativas []cooperativa

var cooperativasList = cooperativas{
	{
		ID:                1,
		RazonSocial:       "Cooperativa 1",
		MatriculaNacional: 1,
	},
}


func optionsCooperativa(w http.ResponseWriter, r *http.Request)  {

	w.Header().Set("Access-Control-Allow-Origin", "*")
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

func updatePatchCooperativa(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
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

	var updateCooperativa cooperativa
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

func updateCooperativa(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
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

	var updateCooperativa cooperativa
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

func deleteCooperativa(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
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

func getCooperativa(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
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

func createCooperativa(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		_,err:=fmt.Fprintf(w, "405 - Method Not Allowed")
		if err != nil {
			http.Error(w, err.Error(), 500)
		}
		return
	}

	var newCooperativa cooperativa
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



func indexRoute(w http.ResponseWriter, r *http.Request) {
	_,err:=fmt.Fprintf(w, "Bienvenido to my cooperativas api ")
	if err != nil {
		http.Error(w, err.Error(), 500)
	}

}

func getCooperativas(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
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

func main() {

	if err:=run(); err != nil {
		_,err:= fmt.Fprintf(os.Stderr, "%s\n", err)
		if err != nil {
			//TODO handle error
		}
	}

}

func run() error  {
	router := mux.NewRouter().StrictSlash(true)
	router.UseEncodedPath()
	router.HandleFunc("/", indexRoute).Name("Index")
	router.HandleFunc("/cooperativas", getCooperativas).Methods(http.MethodGet).Name("getCooperativas")
	router.HandleFunc("/cooperativas", createCooperativa).Methods(http.MethodPost).Name("postCooperativa")
	router.HandleFunc("/cooperativas/{id:[0-9]+}", deleteCooperativa).Methods(http.MethodDelete).Name("deleteCooperativa")
	router.HandleFunc("/cooperativas/{id:[0-9]+}", getCooperativa).Methods(http.MethodGet).Name("getCooperativa")
	router.HandleFunc("/cooperativas/{id:[0-9]+}", updateCooperativa).Methods(http.MethodPut).Name("updateCooperativa")
	router.HandleFunc("/cooperativas/{id:[0-9]+}", updatePatchCooperativa).Methods(http.MethodPatch).Name("patchCooperativa")
	router.HandleFunc("/cooperativas/", optionsCooperativa).Methods(http.MethodOptions).Name("optionCooperativa")

	router.Use(mux.CORSMethodMiddleware(router))

	log.Fatal(http.ListenAndServe(":3001", router))

	return nil

}
