package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
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
	json.NewEncoder(w).Encode(methodsList)
}

func updatePatchCooperativa(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	coopertivaId, err := strconv.Atoi(vars["id"])

	w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Method == http.MethodOptions {
		return
	}

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Id inválido")
		return
	}

	var updateCooperativa cooperativa
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Inserte datos válidos")
	}

	json.Unmarshal(reqBody, &updateCooperativa)

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
			json.NewEncoder(w).Encode(cooperativasList[i])
		}
	}
}

func updateCooperativa(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	coopertivaId, err := strconv.Atoi(vars["id"])

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Id inválido")
		return
	}

	var updateCooperativa cooperativa
	reqBody, err := ioutil.ReadAll(r.Body)

	w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Method == http.MethodOptions {
		return
	}

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Inserte datos válidos")
	}

	json.Unmarshal(reqBody, &updateCooperativa)

	for i, cooperativa := range cooperativasList {
		if cooperativa.ID == coopertivaId {

			cooperativasList = append(cooperativasList[:i], cooperativasList[i+1:]...)
			updateCooperativa.ID = coopertivaId

			cooperativasList = append(cooperativasList, updateCooperativa)

			w.Header().Set("Content-type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(cooperativasList[i])
		}
	}

}

func deleteCooperativa(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	coopertivaId, err := strconv.Atoi(vars["id"])

	w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Method == http.MethodOptions {
		return
	}

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Id inválido")
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
			json.NewEncoder(w).Encode(cooperativa)
		}
	}
}

func getCooperativa(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	coopertivaId, err := strconv.Atoi(vars["id"])

	w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Method == http.MethodOptions {
		return
	}

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Id inválido")
		return
	}

	for _, cooperativa := range cooperativasList {
		if cooperativa.ID == coopertivaId {
			w.Header().Set("Content-type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(cooperativa)
		}
	}
}

func createCooperativa(w http.ResponseWriter, r *http.Request) {
	var newCooperativa cooperativa
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Inserte datos válidos")
	}

	json.Unmarshal(reqBody, &newCooperativa)
	newCooperativa.ID = len(cooperativasList) + 1
	cooperativasList = append(
		cooperativasList,
		newCooperativa,
	)

	w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Method == http.MethodOptions {
		return
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newCooperativa)
}

func getCooperativas(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Method == http.MethodOptions {
		return
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cooperativasList)

}

func indexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Bienvenido to my cooperativas api ")

}
func main() {
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
}
