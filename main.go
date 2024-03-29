package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	// "errors"
)

type entry struct {
	ID      string `json:"id"`
	Entry   string `json:"entry"`
	Date    string `json:"date"`
	User_ID string `json:"user_id"`
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

func main() {
	router := gin.Default()
	router.GET("/entries", GetEntry)
	router.POST("/entries", createEntry)
	router.Run("localhost:8080")
}
