package swagger

import (
	"ddd-sample/pkg/httpserver"
)

func InitRouter(server httpserver.HTTPServer) {
	httpserver.RouteSwagger(server)
}
