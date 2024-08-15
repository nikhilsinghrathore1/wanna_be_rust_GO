package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type book struct {
	ID       string `json`
	Title    string
	Author   string
	Quantity int
}
