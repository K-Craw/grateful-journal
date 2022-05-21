package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	// "errors"
)

func GetEntry(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, entries)
}
