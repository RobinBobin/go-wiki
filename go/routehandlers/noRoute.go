package routehandlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NoRoute(ctx *gin.Context) {
	ctx.Redirect(http.StatusFound, "/arbitrary")
}
