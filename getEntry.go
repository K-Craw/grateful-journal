package main

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	// "errors"
)

func GetEntry(c *gin.Context) {
	db, err := sql.Open("sqlite3", "./entries.db")
	checkErr(err)
	db.Exec("SELECT json_group_array(json_object('id', id, 'entry', entry, 'date', date, 'user_name', user_name)) AS json_result FROM (SELECT * FROM entries ORDER BY id);")
	db.Close()
	c.IndentedJSON(http.StatusOK, entries)
}
