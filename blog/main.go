package main

import (
	"blog/internal/routers"
	"net/http"
	"time"
)

func main() {
	// r := gin.Default()
	// r.GET("/PING", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{"mess": "pong"})
	// })
	// r.Run()
	router := routers.NewRouter()
	s := &http.Server{
		Addr:           "8081",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	_ = s.ListenAndServe()
}
