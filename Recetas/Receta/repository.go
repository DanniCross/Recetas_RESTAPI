package Receta

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

type Repository struct{}

const SERVER = "localhost:27017"

const DBNAME = "REST_API"

const DOCNAME = "recetas"

func (r Repository) GetRecetas() Recetas {
	recetas, err := mgo.Dial(SERVER)
	if err != nil {
		fmt.Println("Falló conexión con el servidor de mongo: ", err)
	}
	defer recetas.Close()
	c := recetas.DB(DBNAME).C(DOCNAME)
	results := Recetas{}
	if err := c.Find(bson.M{}).All(&results); err != nil {
		fmt.Println("No se pudieron escribir los resultados: ", err)
	}

	return results
}

func (r Repository) GetReceta(pos string) Recetas {
	session, err := mgo.Dial(SERVER)
	if err != nil {
		fmt.Println("Falló conexión con el servidor de mongo:", err)
	}
	defer session.Close()
	c := session.DB(DBNAME).C(DOCNAME)
	results := Recetas{}
	if err := c.Find(bson.M{"pos": pos}).All(&results); err != nil {
		fmt.Println("No se pudieron escribir los resultados:", err)
	}
	return results
}

func (r Repository) AddReceta(receta Receta) bool {
	session, err := mgo.Dial(SERVER)
	defer session.Close()

	receta.ID = bson.NewObjectId()
	session.DB(DBNAME).C(DOCNAME).Insert(receta)

	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

func (r Repository) UpdateReceta(receta Receta) bool {
	session, err := mgo.Dial(SERVER)
	defer session.Close()
	log.Println(receta)
	session.DB(DBNAME).C(DOCNAME).UpdateId(receta.ID, receta)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

func (r Repository) DeleteReceta(pos string) string {
	session, err := mgo.Dial(SERVER)
	defer session.Close()
	if err = session.DB(DBNAME).C(DOCNAME).Remove(bson.M{"pos": pos}); err != nil {
		log.Fatal(err)
		return "INTERNAL ERR"
	}
	return "OK"
}
