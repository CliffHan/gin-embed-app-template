package main

import (
	"embed"

	"github.com/gin-gonic/gin"
)

const ASSET_PATH = "app/dist"

//go:embed app/dist/*
var staticFiles embed.FS

func main() {
	r := gin.Default()

	ServeEmbeddedAsset(r, staticFiles, ASSET_PATH)
	ServeAPIs(r)

	r.Run(":8080")
}
