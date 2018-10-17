package Receta

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type Controller struct {
	Repository Repository
}

func (c *Controller) Index(w http.ResponseWriter, r *http.Request) {
	recetas := c.Repository.GetRecetas()
	log.Println(recetas)
	data, _ := json.Marshal(recetas)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
	return
}

func (c *Controller) DelIndex(w http.ResponseWriter, r *http.Request) {
	log.Println("!No ha ingresado datos para borrar o la base de datos está vacia¡")
	return
}

func (c *Controller) GetReceta(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	pos := params["pos"]
	recetas := c.Repository.GetReceta(pos)
	log.Println(recetas)
	data, _ := json.Marshal(recetas)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
	return
}

func (c *Controller) AddReceta(w http.ResponseWriter, r *http.Request) {

	var receta Receta
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		log.Fatalln("Error AddReceta", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err := r.Body.Close(); err != nil {
		log.Fatalln("Error AddReceta", err)
	}
	if err := json.Unmarshal(body, &receta); err != nil {
		w.WriteHeader(422)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			log.Fatalln("Error AddReceta unmarshalling data", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	success := c.Repository.AddReceta(receta)
	if !success {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	return
}

func (c *Controller) UpdateReceta(w http.ResponseWriter, r *http.Request) {
	var receta Receta
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		log.Fatalln("Error UpdateReceta", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err := r.Body.Close(); err != nil {
		log.Fatalln("Error AddaUpdateRecetareceta", err)
	}
	if err := json.Unmarshal(body, &receta); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			log.Fatalln("Error Updatereceta unmarshalling data", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	success := c.Repository.UpdateReceta(receta)
	if !success {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	return
}

func (c *Controller) DeleteReceta(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pos := vars["pos"]
	if err := c.Repository.DeleteReceta(pos); err != "" {
		if strings.Contains(err, "404") {
			w.WriteHeader(http.StatusNotFound)
		} else if strings.Contains(err, "500") {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	return
}
