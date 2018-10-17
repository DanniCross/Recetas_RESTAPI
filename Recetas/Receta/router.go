package Receta

import (
	. "RESTAPI/Recetas/Logger"
	"github.com/gorilla/mux"
	"net/http"
)

var controller = &Controller{Repository: Repository{}}

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/recetas/",
		controller.Index,
	},
	Route{
		"Index",
		"GET",
		"/recetas",
		controller.Index,
	},
	Route{
		"GetReceta",
		"GET",
		"/recetas/{pos}",
		controller.GetReceta,
	},
	Route{
		"AddReceta",
		"POST",
		"/recetas/{pos}",
		controller.AddReceta,
	},
	Route{
		"UpdateReceta",
		"PUT",
		"/recetas",
		controller.UpdateReceta,
	},
	Route{
		"DelIndex",
		"DELETE",
		"/recetas/",
		controller.DelIndex,
	},
	Route{
		"DeleteReceta",
		"DELETE",
		"/recetas/{pos}",
		controller.DeleteReceta,
	},
}

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
	return router
}
