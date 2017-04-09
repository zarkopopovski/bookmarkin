package main

import (
	"log"
	"crypto/sha1"
	"time"
	"fmt"
)

type Bookmark struct {
	Id					string  `json:"id"`
	BookmarkUrl			string	`json:"url"`
	BookmarkTitle		string	`json:"title"`
	BookmarkGroup		string	`json:"group"`
}

type Group struct {
	Id                  string   `json:"id"`
	GroupName           string   `json:"group_name"`
}

func (bookmark *Bookmark) CreateNewGroup(dbConnection *DBConnection, groupName string) bool {

	sha1Hash := sha1.New()
	sha1Hash.Write([]byte(time.Now().String() + groupName))
	sha1HashString := sha1Hash.Sum(nil)

	groupID := fmt.Sprintf("%x", sha1HashString)

	query := "INSERT INTO groups(id, group_name) VALUES('"+groupID+"','"+groupName+"')"
	
	_, err := dbConnection.db.Exec(query)

	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

func (bookmark *Bookmark) ListAllGroups(dbConnection *DBConnection) []*Group {

	query := "SELECT id, group_name FROM groups"
	
	rows, err := dbConnection.db.Query(query)

	if err != nil {
		log.Fatal(err)
		return nil
	}

	groups := make([]*Group, 0)

	for rows.Next() {

		newGroup := new(Group)

		err := rows.Scan(
			&newGroup.Id, 
			&newGroup.GroupName)

		if err != nil {
			log.Fatal(err)
			return nil
		}

		groups = append(groups, newGroup)

	}

	return groups
}


func (bookmark *Bookmark) CreateNewBookmark(dbConnection *DBConnection) bool {

	sha1Hash := sha1.New()
	sha1Hash.Write([]byte(time.Now().String() + bookmark.BookmarkTitle + " " + bookmark.BookmarkUrl + " " + bookmark.BookmarkGroup))
	sha1HashString := sha1Hash.Sum(nil)

	bookmarkID := fmt.Sprintf("%x", sha1HashString)

	query := "INSERT INTO bookmarks(id, bookmark_url, bookmark_title, bookmark_group) VALUES('"+
	bookmarkID+"','"+bookmark.BookmarkUrl+"','"+bookmark.BookmarkTitle+"','"+bookmark.BookmarkGroup+"')"
	
	_, err := dbConnection.db.Exec(query)

	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

func (bookmark *Bookmark) ListAllBookmarks(dbConnection *DBConnection) []*Bookmark {

	query := "SELECT id, bookmark_url, bookmark_title, bookmark_group FROM bookmarks"
	
	rows, err := dbConnection.db.Query(query)

	if err != nil {
		log.Fatal(err)
		return nil
	}

	bookmarks := make([]*Bookmark, 0)

	for rows.Next() {

		newBookmark := new(Bookmark)

		err := rows.Scan(
			&newBookmark.Id, 
			&newBookmark.BookmarkUrl, 
			&newBookmark.BookmarkTitle, 
			&newBookmark.BookmarkGroup)

		if err != nil {
			log.Fatal(err)
			return nil
		}

		bookmarks = append(bookmarks, newBookmark)

	}

	return bookmarks
}
