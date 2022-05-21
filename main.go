package main

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	// "errors"
)

type entry struct {
	ID      int    `json:"id"`
	Entry   string `json:"entry"`
	Date    string `json:"date"`
	User_ID int    `json:"user_id"`
}

var entries = []entry{
	{ID: "1", Entry: "My first message", Date: "05/18/2022", User_ID: "user1"},
	{ID: "2", Entry: "Hi name is user2", Date: "05/19/2022", User_ID: "user2"},
	{ID: "3", Entry: "I made third message", Date: "05/20/2022", User_ID: "user1"},
}

func createEntry(c *gin.Context) {
	var newEntry entry
	if err := c.BindJSON(&newEntry); err != nil {
		return
	}
	entries = append(entries, newEntry)
	c.IndentedJSON(http.StatusCreated, newEntry)
}

func insertRows(ctx context.Context, tx pgx.Tx, accts [4]uuid.UUID) error {
	// Insert four rows into the "accounts" table.
	log.Println("Creating new rows...")
	if _, err := tx.Exec(ctx,
		"INSERT INTO entries (id, entry, date, user_id) VALUES (1, 'hello', 2001-02-08 ), (2, 'this', 2022-01-02), (3, 'is', 2018-04-05), (4, 'a test', '2019-02-21')", accts[0], 250, accts[1], 100, accts[2], 500, accts[3], 300); err != nil {
		return err
	}
	return nil
}

func main() {
	router := gin.Default()
	router.GET("/entries", GetEntry)
	router.POST("/entries", createEntry)
	router.Run("localhost:8080")
}
