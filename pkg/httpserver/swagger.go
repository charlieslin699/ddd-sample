package httpserver

import (
	swaggerfiles "github.com/swaggo/files"
	ginswagger "github.com/swaggo/gin-swagger"
)

func RouteSwagger(server HTTPServer) {
	server.Route(
		GET("/swagger/*any", func(ctx *Context) (RestfulResult, error) {
			ginswagger.WrapHandler(
				swaggerfiles.Handler,
				ginswagger.DefaultModelsExpandDepth(-1), // 隱藏Model
			)(ctx.Context)

			return ctx.Next()
		}),
	)
}
