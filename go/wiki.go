package main

import (
	"embed"
	"example/wiki/routehandlers"
	"fmt"
	"html/template"
	"io/fs"
	"log"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

//go:embed dist/*
var embeddedDist embed.FS

//go:embed html/*
var embeddedHtml embed.FS

func main() {
	router := gin.Default()

	templ := template.Must(template.New("").ParseFS(embeddedHtml, "html/templates/*.html"))

	router.SetHTMLTemplate(templ)
	router.SetTrustedProxies(nil)
	router.Static("/static", "html/static")

	dist, err := static.EmbedFolder(embeddedDist, "dist")

	if err != nil {
		log.Fatalln(err)
	}

	router.Use(static.Serve("/", dist))

	router.Use(func(ctx *gin.Context) {
		shouldInvokeNext := func() bool {
			path := ctx.Request.URL.Path

			// It's a root path or already has an extension.
			if path == "/" || strings.Contains(filepath.Base(path), ".") {
				return true
			}

			filepath := fmt.Sprintf("dist%v.html", path)

			_, err := fs.Stat(embeddedDist, filepath)

			if err != nil {
				return true
			}

			ctx.FileFromFS(filepath, http.FS(embeddedDist))
			ctx.Abort()

			return false
		}()

		if shouldInvokeNext {
			ctx.Next()
		}
	})

	router.NoRoute(routehandlers.NoRoute)

	router.GET("/arbitrary/*params", routehandlers.GetArbitrary)

	router.Run(":8080")
}
