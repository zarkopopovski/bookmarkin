package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func RoutesMap(api *ApiConnection) Routes {
	var routes = Routes{
		Route{"Index", "GET", "/", api.Index},
		Route{"ReadPageTitle", "POST", "/read_page_title", api.bHandlers.ReadPageTitle},
		Route{"SaveBookmark", "POST", "/save_bookmark", api.bHandlers.SaveBookmark},
		Route{"CreateBookmark", "POST", "/create_bookmark", api.bHandlers.CreateBookmark},
		Route{"ListBookmarks", "POST", "/list_bookmarks", api.bHandlers.ListBookmarks},
		Route{"UpdateBookmark", "POST", "/update_bookmarks", api.bHandlers.UpdateBookmarks},
		Route{"DeleteBookmark", "POST", "/delete_bookmarks", api.bHandlers.DeleteBookmark},
	}

	return routes
}
