package main

import (
	"embed"
	"html/template"

	"github.com/gin-gonic/gin"

	"example/wiki/routehandlers"
)

//go:embed html/*
var embeddedFS embed.FS

func main() {
	router := gin.Default()

	templ := template.Must(template.New("").ParseFS(embeddedFS, "html/templates/*.html"))

	router.SetHTMLTemplate(templ)
	router.SetTrustedProxies(nil)
	router.Static("/static", "html/static")

	router.NoRoute(routehandlers.NoRoute)

	router.GET("/:reason/*params", routehandlers.GetRoot)

	router.Run(":8080")
}
