package main

type Exporter struct {
	dbConnection *DBConnection
}

type IExporter interface {
	exportBookmarks(s string) (err error)
}

func (exporter *Exporter) exportBookmarks() (err error) {
	return nil
}
