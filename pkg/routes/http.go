package routes

import (
	"src/pkg/myerror"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	r.GET("/health-check", myerror.HandleError(Healthcheck))
}

func Healthcheck(c *gin.Context) error {
	c.String(200, "ok")
	return nil
}
