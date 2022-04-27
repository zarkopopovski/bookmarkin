package main

import (
	"github.com/julienschmidt/httprouter"
)

func NewRouter(api *ApiConnection) *httprouter.Router {
	router := httprouter.New()

	router.GET("/", api.Index)
	
	router.POST("/create_group", api.bHandlers.CreateBookmarkGroup)
	router.POST("/update_group", api.bHandlers.UpdateBookmarkGroup)
	router.POST("/delete_group", api.bHandlers.DeleteBookmarkGroup)
	router.POST("/list_groups", api.bHandlers.ListBookmarkGroups)
	router.POST("/read_page_title", api.bHandlers.ReadPageTitle)
	router.POST("/save_bookmark", api.bHandlers.SaveBookmark)
	router.POST("/create_bookmark", api.bHandlers.CreateBookmark)
	router.POST("/list_bookmarks", api.bHandlers.ListBookmarks)
	router.POST("/list_group_bookmarks", api.bHandlers.ListBookmarksInGroup)
  router.POST("/update_bookmarks", api.bHandlers.UpdateBookmarks)	
  router.POST("/delete_bookmarks", api.bHandlers.DeleteBookmark)
  router.POST("/export_bookmarks", api.bHandlers.ExportBookmarks)
  router.POST("/change_bookmark_group", api.bHandlers.ChangeBookmarkGroup)
  router.POST("/create_user", api.uHandlers.CreateUserAccount)
  router.POST("/change_password", api.uHandlers.ChangeUserPassword)    
  router.POST("/login_user", api.uHandlers.LoginWithCredentials) 
  router.GET("/favicon.ico", api.bHandlers.ShowFavIcon) 
	
	//router.ServeFiles("/web/*filepath", http.Dir("./web"))
	
	return router
}
