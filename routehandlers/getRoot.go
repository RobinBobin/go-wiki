package routehandlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetRoot(ctx *gin.Context) {
	data := make(map[string]string)

	for _, param := range ctx.Params {
		data[param.Key] = param.Value
	}

	ctx.HTML(http.StatusOK, "index.html", data)
}
