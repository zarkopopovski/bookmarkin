package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/mat/besticon/besticon"
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

//TODO: Check if the icon exist in database by selecting the base URL, and if exist return icon ID, otherwise download and save
func (bHandlers *BookmarkHandlers) CreateBookmark(w http.ResponseWriter, r *http.Request) {
	bookmarkURL := r.FormValue("bookmark_url")
	bookmarkGroup := r.FormValue("bookmark_group")
	userID := r.FormValue("user_id")
	bookmarkTitle := ""
	result := false

	besticon.SetLogOutput(ioutil.Discard)
	finder := besticon.IconFinder{}

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

	icons, errIcon := finder.FetchIcons(bookmarkURL)

	iconURL := ""
	iconEncoded := ""

	if errIcon == nil && len(icons) > 0 {
		best := icons[(len(icons) - 1)]
		iconURL = best.URL

		resp, err := http.Get(iconURL)
		defer resp.Body.Close()

		if err == nil {
			contents, err := ioutil.ReadAll(resp.Body)

			if err == nil {
				iconEncoded = base64.StdEncoding.EncodeToString(contents)
			}
		}
	}

	if bookmarkTitle != "" {
		fmt.Println(bookmarkTitle)

		bookmark := &Bookmark{
			BookmarkUrl:   bookmarkURL,
			BookmarkTitle: bookmarkTitle,
			BookmarkGroup: bookmarkGroup,
			IconURL:       iconURL,
			IconEncoded:   iconEncoded}

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

//TODO: Join with bookmark_icons table comparing with base_url because the bookmark icon may exist only once in database
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

//TODO: Join with bookmark_icons table comparing with base_url because the bookmark icon may exist only once in database
func (bHandlers *BookmarkHandlers) ExportBookmarks(w http.ResponseWriter, r *http.Request) {
	userID := r.FormValue("user_id")
	exportType := r.FormValue("export_type")

	bookmark := &Bookmark{}

	result := bookmark.ListAllBookmarks(bHandlers.dbConnection, userID)

	exporter := &Exporter{}
	expEngine := exporter.getExporterByType(exportType)

	exportData := expEngine.exportBookmarks(result)

	fileDataInBytes := []byte(exportData)

	bytesReader := bytes.NewReader(fileDataInBytes)

	io.Copy(w, bytesReader)

	w.Header().Set("Content-Disposition", "attachment; filename='export."+exportType+"'")
	w.Header().Set("Content-Type", "application/"+exportType)
	w.Header().Set("Content-Length", strconv.Itoa(len(fileDataInBytes)))

	return
}

//TODO: Join with bookmark_icons table comparing with base_url because the bookmark icon may exist only once in database
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

//TODO: Check if there are other bookmarks that are using the same icon by counting base_url, if not, delete also the icon from the bookmark_icon table
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
