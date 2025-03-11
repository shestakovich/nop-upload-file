package main

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"os"
)

func NopUploadFile() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		written, err := io.Copy(io.Discard, c.Request.Body)

		if err != nil {
			c.Error(err) //nolint
			c.Status(http.StatusInternalServerError)
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"written": written,
			"took":    time.Since(start).String(),
		})
	}
}

func SetupRoutes(
	router gin.IRouter,
) {
	gapiRoutes := router.Group("/")
	gapiRoutes.POST("/nop_upload_file", NopUploadFile())

}

func run() error {
	routerEngine := gin.New()
	routerEngine.Use(gin.Recovery())

	SetupRoutes(
		routerEngine,
	)

	err := routerEngine.Run(":4000")

	return err
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
