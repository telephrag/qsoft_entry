package main

import (
	"log"
	"qsoft_entry/when"

	"github.com/gin-gonic/gin"
)

func main() {
	err := when.Run(gin.New(), "localhost:3000")
	if err != nil {
		log.Fatal(err)
	}
}
