package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type BookmarkHandlers struct {
	dbConnection *DBConnection
}

func (bHandlers *BookmarkHandlers) ShowFavIcon(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./web/bookmarkin-flat.png")
}

func (bHandlers *BookmarkHandlers) CreateBookmarkGroup(w http.ResponseWriter, r *http.Request) {
	groupName := r.FormValue("group_name")
	userID := r.FormValue("user_id")

	if groupName != "" {

		bookmark := &Bookmark{}

		result := bookmark.CreateNewGroup(bHandlers.dbConnection, groupName, userID)

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

func (bHandlers *BookmarkHandlers) UpdateBookmarkGroup(w http.ResponseWriter, r *http.Request) {
	groupName := r.FormValue("group_name")
	groupID := r.FormValue("group_id")
	userID := r.FormValue("user_id")

	if groupName != "" {

		bookmark := &Bookmark{}

		result := bookmark.UpdateExistingGroup(bHandlers.dbConnection, groupID, groupName, userID)

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

func (bHandlers *BookmarkHandlers) DeleteBookmarkGroup(w http.ResponseWriter, r *http.Request) {
	groupID := r.FormValue("group_id")
	forceDelete := r.FormValue("force")
	userID := r.FormValue("user_id")

	bForceDelete, err := strconv.ParseBool(forceDelete)
	if err != nil {
		bForceDelete = false
	}

	if groupID != "" {

		bookmark := &Bookmark{}

		result := bookmark.ListAllBookmarksInGroup(bHandlers.dbConnection, groupID, userID)

		if result != nil && len(result) > 0 {
			if bForceDelete == true {
				result := bookmark.DeleteBookmarksInGroup(bHandlers.dbConnection, groupID, userID)

				if result {

					result := bookmark.DeleteBookmarkGroupByID(bHandlers.dbConnection, groupID, userID)

					if result {
						w.Header().Set("Content-Type", "application/json; charset=UTF-8")
						w.WriteHeader(http.StatusOK)

						if err := json.NewEncoder(w).Encode(result); err != nil {
							panic(err)
						}
						return
					} else {
						w.Header().Set("Content-Type", "application/json; charset=UTF-8")
						w.WriteHeader(http.StatusNotFound)
						if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Error deleting group."}); err != nil {
							panic(err)
						}
					}
				} else {
					w.Header().Set("Content-Type", "application/json; charset=UTF-8")
					w.WriteHeader(http.StatusNotFound)
					if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Error deleting bookmarks from group."}); err != nil {
						panic(err)
					}
				}
			} else {
				w.Header().Set("Content-Type", "application/json; charset=UTF-8")
				w.WriteHeader(http.StatusNotFound)
				if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Bookmarks exists in the group."}); err != nil {
					panic(err)
				}
			}
		} else {
			result := bookmark.DeleteBookmarkGroupByID(bHandlers.dbConnection, groupID, userID)

			if result {
				w.Header().Set("Content-Type", "application/json; charset=UTF-8")
				w.WriteHeader(http.StatusOK)

				if err := json.NewEncoder(w).Encode(result); err != nil {
					panic(err)
				}
				return
			} else {
				w.Header().Set("Content-Type", "application/json; charset=UTF-8")
				w.WriteHeader(http.StatusNotFound)
				if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Error deleting group."}); err != nil {
					panic(err)
				}
			}
		}

		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)
	if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Not Found"}); err != nil {
		panic(err)
	}
}

func (bHandlers *BookmarkHandlers) ListBookmarkGroups(w http.ResponseWriter, r *http.Request) {
	userID := r.FormValue("user_id")

	bookmark := &Bookmark{}

	result := bookmark.ListAllGroups(bHandlers.dbConnection, userID)

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
	userID := r.FormValue("user_id")
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
			BookmarkUrl:   bookmarkURL,
			BookmarkTitle: bookmarkTitle,
			BookmarkGroup: bookmarkGroup}

		result = bookmark.CreateNewBookmark(bHandlers.dbConnection, userID)

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
	userID := r.FormValue("user_id")
	result := false

	if bookmarkTitle != "" {
		fmt.Println(bookmarkTitle)

		bookmark := &Bookmark{
			BookmarkUrl:   bookmarkURL,
			BookmarkTitle: bookmarkTitle,
			BookmarkGroup: bookmarkGroup}

		result = bookmark.CreateNewBookmark(bHandlers.dbConnection, userID)

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
	userID := r.FormValue("user_id")

	bookmark := &Bookmark{}

	result := bookmark.ListAllBookmarks(bHandlers.dbConnection, userID)

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

func (bHandlers *BookmarkHandlers) ListBookmarksInGroup(w http.ResponseWriter, r *http.Request) {
	groupID := r.FormValue("group_id")
	userID := r.FormValue("user_id")

	bookmark := &Bookmark{}

	result := bookmark.ListAllBookmarksInGroup(bHandlers.dbConnection, groupID, userID)

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
	bookmarkTitle := r.FormValue("bookmark_title")
	bookmarkID := r.FormValue("bookmark_id")
	userID := r.FormValue("user_id")

	result := false

	if bookmarkTitle != "" {
		fmt.Println(bookmarkTitle)

		bookmark := &Bookmark{}

		result = bookmark.UpdateExistingBookmark(bHandlers.dbConnection, bookmarkID, bookmarkTitle, userID)

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

func (bHandlers *BookmarkHandlers) DeleteBookmark(w http.ResponseWriter, r *http.Request) {
	bookmarkID := r.FormValue("bookmark_id")
	userID := r.FormValue("user_id")
	result := false

	if bookmarkID != "" {
		bookmark := &Bookmark{}

		result = bookmark.DeleteBookmarkByID(bHandlers.dbConnection, bookmarkID, userID)

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

func (bHandlers *BookmarkHandlers) ChangeBookmarkGroup(w http.ResponseWriter, r *http.Request) {
	bookmarkID := r.FormValue("bookmark_id")
	newGroupID := r.FormValue("group_id")
	result := false

	if bookmarkID != "" {
		bookmark := &Bookmark{Id: bookmarkID}
		group := &Group{Id: newGroupID}

		result = bookmark.ChangeGroup(bHandlers.dbConnection, group)

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
