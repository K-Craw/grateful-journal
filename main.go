package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"errors"
)

type entry struct {
	ID      string json: "id"
	Entry   string json: "entry"
	Date    string json: "date"
	User_ID string json: "user_id"
}