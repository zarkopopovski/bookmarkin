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
		Route{"CreateBookmarkGroup", "POST", "/create_group", api.bHandlers.CreateBookmarkGroup},
		Route{"UpdateBookmarkGroup", "POST", "/update_group", api.bHandlers.UpdateBookmarkGroup},
		Route{"DeleteBookmarkGroup", "POST", "/delete_group", api.bHandlers.DeleteBookmarkGroup},
		Route{"ListBookmarkGroups", "POST", "/list_groups", api.bHandlers.ListBookmarkGroups},
		Route{"ReadPageTitle", "POST", "/read_page_title", api.bHandlers.ReadPageTitle},
		Route{"SaveBookmark", "POST", "/save_bookmark", api.bHandlers.SaveBookmark},
		Route{"CreateBookmark", "POST", "/create_bookmark", api.bHandlers.CreateBookmark},
		Route{"ListBookmarks", "POST", "/list_bookmarks", api.bHandlers.ListBookmarks},
		Route{"ListBookmarksInGroup", "POST", "/list_group_bookmarks", api.bHandlers.ListBookmarksInGroup},
		Route{"UpdateBookmark", "POST", "/update_bookmarks", api.bHandlers.UpdateBookmarks},
		Route{"DeleteBookmark", "POST", "/delete_bookmarks", api.bHandlers.DeleteBookmark},
		Route{"ChangeBookmarkGroup", "POST", "/change_bookmark_group", api.bHandlers.ChangeBookmarkGroup},
		Route{"CreateNewUser", "POST", "/create_user", api.uHandlers.CreateUserAccount},
		Route{"ChangePassword", "POST", "/change_password", api.uHandlers.ChangeUserPassword},
		Route{"LoginUser", "POST", "/login_user", api.uHandlers.LoginWithCredentials},
	}

	return routes
}
