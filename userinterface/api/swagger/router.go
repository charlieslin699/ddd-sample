package swagger

import (
	"ddd-sample/pkg/httpserver"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginswagger "github.com/swaggo/gin-swagger"
)

func InitRouter(server httpserver.HTTPServer) {
	server.Route(
		httpserver.GET("/swagger/*any", func(ctx *gin.Context) (httpserver.RestfulResult, error) {
			ginswagger.WrapHandler(
				swaggerfiles.Handler,
				ginswagger.DefaultModelsExpandDepth(-1), // 隱藏Model
			)(ctx)
			return httpserver.Next()
		}),
	)
}
