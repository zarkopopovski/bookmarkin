package main

import (
	"encoding/json"
)

type IExporter interface {
	exportBookmarks(bookmarks []*Bookmark) (export string)
}

type Exporter struct {
}

type HTMLExporter struct {
}

type JSONExporter struct {
}

func (exporter *Exporter) getExporterByType(expType string) IExporter {
	if expType == "html" {
		return &HTMLExporter{}
	} else if expType == "json" {
		return &JSONExporter{}
	}

	return nil
}

func (exporter *HTMLExporter) exportBookmarks(bookmarks []*Bookmark) (export string) {
	htmlString := "<html><head><title>Bookmarks</title></head><body><ul>"

	for index := 0; index < len(bookmarks); index++ {
		bookmark := bookmarks[index]
		htmlString += "<li><img src='data:image/png;base64," + bookmark.IconEncoded + "' width=16 height=16/>" + bookmark.BookmarkTitle + "<br/>" + bookmark.BookmarkUrl + "</li>"
	}

	htmlString += "</ul></body></html>"
	return htmlString
}

func (exporter *JSONExporter) exportBookmarks(bookmarks []*Bookmark) (export string) {
	jsonString, err := json.Marshal(bookmarks)

	if err != nil {
		return ""
	}
	return string(jsonString)
}
