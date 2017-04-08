package main

import (
	"encoding/json"
	"net/http"
	"fmt"
)

type BookmarkHandlers struct {
	dbConnection *DBConnection
}

func (bHandlers *BookmarkHandlers) CreateBookmarkGroup(w http.ResponseWriter, r *http.Request) {


}

func (bHandlers *BookmarkHandlers) UpdateBookmarkGroup(w http.ResponseWriter, r *http.Request) {


}

func (bHandlers *BookmarkHandlers) DeleteBookmarkGroup(w http.ResponseWriter, r *http.Request) {


}

func (bHandlers *BookmarkHandlers) ListBookmarkGroups(w http.ResponseWriter, r *http.Request) {


}

func (bHandlers *BookmarkHandlers) ReadPageTitle(w http.ResponseWriter, r *http.Request) {
	bookmarkURL := r.FormValue("bookmark_url")
	bookmarkTitle := ""
	
	resp, err := http.Get(bookmarkURL)
	
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if title, ok := GetHtmlTitle(resp.Body); ok {
		bookmarkTitle = title
	} else {
		println("Fail to get HTML title")
	}

	if bookmarkTitle != "" {
		fmt.Println(bookmarkTitle)
		
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)

		if err := json.NewEncoder(w).Encode(bookmarkTitle); err != nil {
			panic(err)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)
	if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Not Found"}); err != nil {
		panic(err)
	}
}

func (bHandlers *BookmarkHandlers) CreateBookmark(w http.ResponseWriter, r *http.Request) {
	bookmarkURL := r.FormValue("bookmark_url")
	bookmarkGroup := r.FormValue("bookmark_group")
	bookmarkTitle := ""
	result := false
	
	resp, err := http.Get(bookmarkURL)
	
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if title, ok := GetHtmlTitle(resp.Body); ok {
		bookmarkTitle = title
	} else {
		println("Fail to get HTML title")
	}

	if bookmarkTitle != "" {
		fmt.Println(bookmarkTitle)

		bookmark := &Bookmark{
			BookmarkUrl:bookmarkURL, 
			BookmarkTitle:bookmarkTitle, 
			BookmarkGroup:bookmarkGroup}

		result = bookmark.CreateNewBookmark(bHandlers.dbConnection)

		if result {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusOK)

			if err := json.NewEncoder(w).Encode(result); err != nil {
				panic(err)
			}
			return
		}

	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)
	if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Not Found"}); err != nil {
		panic(err)
	}
}

func (bHandlers *BookmarkHandlers) SaveBookmark(w http.ResponseWriter, r *http.Request) {
	bookmarkURL := r.FormValue("bookmark_url")
	bookmarkTitle := r.FormValue("bookmark_title")
	bookmarkGroup := r.FormValue("bookmark_group")
	result := false
	
	resp, err := http.Get(bookmarkURL)
	
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if title, ok := GetHtmlTitle(resp.Body); ok {
		bookmarkTitle = title
	} else {
		println("Fail to get HTML title")
	}

	if bookmarkTitle != "" {
		fmt.Println(bookmarkTitle)

		bookmark := &Bookmark{
			BookmarkUrl:bookmarkURL, 
			BookmarkTitle:bookmarkTitle, 
			BookmarkGroup:bookmarkGroup}

		result = bookmark.CreateNewBookmark(bHandlers.dbConnection)

		if result {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusOK)

			if err := json.NewEncoder(w).Encode(result); err != nil {
				panic(err)
			}
			return
		}

	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)
	if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Not Found"}); err != nil {
		panic(err)
	}
}

func (bHandlers *BookmarkHandlers) ListBookmarks(w http.ResponseWriter, r *http.Request) {
	bookmark := &Bookmark{}

	result := bookmark.ListAllBookmarks(bHandlers.dbConnection)

		if result != nil {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusOK)

			if err := json.NewEncoder(w).Encode(result); err != nil {
				panic(err)
			}
			return
		}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)
	if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Not Found"}); err != nil {
		panic(err)
	}
}

func (bHandlers *BookmarkHandlers) UpdateBookmarks(w http.ResponseWriter, r *http.Request) {

}

func (bHandlers *BookmarkHandlers) DeleteBookmark(w http.ResponseWriter, r *http.Request) {

}