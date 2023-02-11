package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/herizal95/golang-jwt-gin/helper"
)

func main() {

	router := gin.Default()

	router.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "welcome golang gin jwt")
	})

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.ErrorPanic(err)
}
