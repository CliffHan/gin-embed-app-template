package main

import (
	"embed"
	"io/fs"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ServeEmbeddedAsset(engine *gin.Engine, staticFiles embed.FS, assetPath string) {
	fs, err := fs.Sub(staticFiles, assetPath)
	if err != nil {
		panic(err)
	}
	engine.StaticFS("/", http.FS(fs))
}

func ServeAPIs(engine *gin.Engine) {
	engine.POST("/api/hello", postHello)
}

func postHello(c *gin.Context) {
	var json struct {
		User string `json:"user" binding:"required"`
	}
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User field is required"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Hello " + json.User})
}
