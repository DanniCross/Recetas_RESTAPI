package Receta

import "gopkg.in/mgo.v2/bson"

type Receta struct {
	ID           bson.ObjectId `bson:"_id"`
	Nombre       string        `bson:"nombre" json:"nombre"`
	Ingredientes []string      `bson:"ingredientes" json:"ingredientes"`
	Elaboracion  string        `bson:"elaboracion" json:"elaboracion"`
	Pos          string        `bson:"pos" json:"pos"`
}

type Recetas []Receta
