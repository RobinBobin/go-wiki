package routehandlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetArbitrary(ctx *gin.Context) {
	pathParams := make(map[string]string)

	for _, param := range ctx.Params {
		pathParams[param.Key] = param.Value
	}

	ctx.HTML(http.StatusOK, "arbitrary.html", gin.H{
		"Path":  pathParams,
		"Query": ctx.Request.URL.Query(),
	})
}
