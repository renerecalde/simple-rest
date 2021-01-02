package main

import (
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/renerecalde/simple-rest/Controller"
	"log"
	"net/http"
	"os"
)

// DB Driver visible to whole program
var DB *sql.DB



func indexRouteHandler(w http.ResponseWriter, r *http.Request) {
	_,err:=fmt.Fprintf(w, "Bienvenido to my cooperativas api ")
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
}


func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

func setHeaderMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		next.ServeHTTP(w, r)
	})
}


func run() error  {

	router := mux.NewRouter().StrictSlash(true)
	router.UseEncodedPath()
	router.Use(loggingMiddleware)
	router.Use(setHeaderMiddleware)
	router.Use(mux.CORSMethodMiddleware(router))
	router.HandleFunc("/", indexRouteHandler).Name("Index")
	router.HandleFunc("/cooperativas",Controller.GetCooperativasHandler).Methods(http.MethodGet).Name("getCooperativas")
	router.HandleFunc("/cooperativas", Controller.CreateCooperativaHandler).Methods(http.MethodPost).Name("postCooperativa")
	router.HandleFunc("/cooperativas/{id:[0-9]+}", Controller.DeleteCooperativaHandler).Methods(http.MethodDelete).Name("deleteCooperativa")
	router.HandleFunc("/cooperativas/{id:[0-9]+}", Controller.GetCooperativaHandler).Methods(http.MethodGet).Name("getCooperativa")
	router.HandleFunc("/cooperativas/{id:[0-9]+}", Controller.UpdateCooperativaHandler).Methods(http.MethodPut).Name("updateCooperativa")
	router.HandleFunc("/cooperativas/{id:[0-9]+}", Controller.UpdatePatchCooperativaHandler).Methods(http.MethodPatch).Name("patchCooperativa")
	router.HandleFunc("/cooperativas/", Controller.OptionsCooperativaHandler).Methods(http.MethodOptions).Name("optionCooperativa")

	log.Fatal(http.ListenAndServe(":3001", router))

	return nil
}

func main() {
	if err:=run(); err != nil {
		_,err:= fmt.Fprintf(os.Stderr, "%s\n", err)
		if err != nil {
			panic(err)
		}
	}
}
