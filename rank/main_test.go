package main

import (
	"testing"

	"github.com/gin-gonic/gin"
)

func TestMainExecution(t *testing.T) {
	gin.SetMode(gin.TestMode)
	go main()
}

func TestRankExecution(t *testing.T) {
	gin.SetMode(gin.TestMode)
	go Rank()
}
