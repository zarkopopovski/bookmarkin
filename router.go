package main

import (
	"flag"
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter(api *ApiConnection) *mux.Router {

	router := mux.NewRouter().StrictSlash(true)

	var routes = RoutesMap(api)

	dir := flag.String("directory", "./web/scripts", "Directory of static files")
	flag.Parse()

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

	fs := http.Dir(*dir)
	fileHandler := http.FileServer(fs)

	router.PathPrefix("/files").Handler(http.StripPrefix("/files", fileHandler))

	//router.HandleFunc("/export_bookmarks", api.bHandlers.ExportBookmarks)

	return router
}
