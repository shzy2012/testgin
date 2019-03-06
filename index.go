package main

import (
	"log"
	"net/http"
	"time"

	"github.com/shzy2012/testgin/logger"

	"github.com/gin-gonic/gin"
)

//Logger middleware
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		// Set example variable
		c.Set("example", "12345")

		// before request

		c.Next()

		// after request
		latency := time.Since(t)
		log.Print(latency)

		// access the status we are sending
		status := c.Writer.Status()
		log.Println(status)
	}
}

func index(context *gin.Context) {
	context.JSON(200, gin.H{
		"index": "i am index",
	})
}

func main() {

	var thisLog logger.ILogger
	thisLog = logger.NewLogger()
	thisLog.Trace("helo")
	thisLog.Info("helo")
	thisLog.Warning("helo")
	thisLog.Error("helo")

	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.GET("/index", func(c *gin.Context) {
		time.Sleep(time.Second * 2)
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	v1 := router.Group("/v1")
	{
		v1.GET("/", index)
	}

	server := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	server.ListenAndServe()

	//router.Run(":8080")
}
