package main

import (
	ginTimeout "github.com/gin-contrib/timeout"
	gin "github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type EchoResponse struct {
	FieldOne string `json:"fieldOne"`
	FieldTwo int64  `json:"fieldTwo"`
}

// hello is a handler function that simulates a long-running process
func hello(c *gin.Context) {
	// intentional delay of 3 seconds
	time.Sleep(3 * time.Second)

	// do not continue if cancelled
	if !c.IsAborted() {
		data := EchoResponse{
			FieldOne: "string-value",
			FieldTwo: 1000,
		}
		c.JSON(http.StatusOK, data)

		// to test after timeout whether we could skip request processing
		println("Not aborted.")
	}
}

func main() {
	router := gin.Default()
	router.Use(ginTimeout.New(
		ginTimeout.WithTimeout(2*time.Second),
		ginTimeout.WithHandler(func(context *gin.Context) {
			context.Next()
		}),
	))
	router.GET("/echo", hello)
	err := router.Run(":8080")
	if err != nil {
		panic(err)
	}
}
