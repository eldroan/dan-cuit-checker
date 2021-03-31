package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

type Error struct {
	Error string `json:"error"`
}

type Cuit struct {
	Cuit      string     `json:"cuit"`
	Situacion *Situacion `json:"situacion"`
}

type Situacion struct {
	Nivel       int    `json:"nivel"`
	Descripcion string `json:"descripcion"`
	Riesgo      string `json:"riesgo"`
}

var situaciones = map[int]Situacion{
	1: {1, "En situación normal", "Situación normal"},
	2: {2, "Con seguimiento especial", "Riesgo bajo"},
	3: {3, "Con problemas", "Riesgo medio"},
	4: {4, "Con alto riesgo de insolvencia", "Riesgo alto"},
	5: {5, "Irrecuperable", "Irrecuperable"},
	6: {6, "Irrecuperable por disposición técnica", "Irrecuperable por disposición técnica"},
}

func getSituacionCuit(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	if _, found := params["cuit"]; !found {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Error{"No se encontró un cuit"})
	} else if cuit, err := strconv.ParseInt(params["cuit"], 10, 64); err == nil {
		rand.Seed(cuit)
		index := rand.Intn(len(situaciones))
		situacion := situaciones[index]
		json.NewEncoder(w).Encode(Cuit{params["cuit"], &situacion})
	} else if r.URL.Path == "/api/cuit/{cuit}" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Error{"Hay que reemplazar el {cuit} de la request por un numero de cuit, ej  /api/cuit/12345678"})
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Error{"Hubo un error que el programador no supo manejar mejor :c"})
	}
}

func noEncontradaHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(Error{fmt.Sprintf("La ruta '%s' no fue encontrada, la unica ruta es '/api/cuit/{cuit}'", r.URL.Path)})
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/cuit", getSituacionCuit).Methods("GET")
	router.HandleFunc("/api/cuit/", getSituacionCuit).Methods("GET")
	router.HandleFunc("/api/cuit/{cuit}", getSituacionCuit).Methods("GET")
	router.NotFoundHandler = http.HandlerFunc(noEncontradaHandler)

	port, found := os.LookupEnv("PORT")
	if !found {
		port = "8080"
	}

	fmt.Println("Ejecutando aplicación en el puerto " + port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
