package main

import (
	"net/http"
	"time"
)

func main() {
	// r := gin.Default()
	// r.GET("/PING", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{"mess": "pong"})
	// })
	// r.Run()
	route := routers.NewRouters()
	s := &http.Server{
		Addr:           "8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
