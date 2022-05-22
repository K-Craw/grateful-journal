package main

import (
	// "database/sql"
	"log"

	// "os"

	"github.com/gin-gonic/gin"
	// "errors"
)

type entry struct {
	ID        int    `json:"id"`
	Entry     string `json:"entry"`
	Date      string `json:"date"`
	User_name string `json:"user_name"`
}

var entries = []entry{
	{ID: 1, Entry: "My first message", Date: "05/18/2022", User_name: "user1"},
	{ID: 2, Entry: "Hi name is user2", Date: "05/19/2022", User_name: "user2"},
	{ID: 3, Entry: "I made third message", Date: "05/20/2022", User_name: "user1"},
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	router := gin.Default()
	router.GET("/entries", GetEntry)
	router.POST("/entries", createEntry)
	router.Run("localhost:8080")
}
