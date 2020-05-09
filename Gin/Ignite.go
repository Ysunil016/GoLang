package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

// Person ...
type Person struct {
	Name string
	Age  int
}

func main() {
	router = gin.Default()
	// counter := 0

	// P1 := Person{"Sunil", 0}

	router.GET("/", home)
	router.Run(":8081")
}

// Counter ...
var Counter int = 0

// P1 ...
var P1 Person = Person{"Sunil", 0}

func home(c *gin.Context) {
	Counter++
	P1.Age = Counter
	c.JSON(http.StatusOK, gin.H{"data": P1})
}
