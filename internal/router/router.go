package router

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/netesh5/go_ecommerce/internal/middleware"
)

type Router struct {
	Name       string
	Method     string
	Path       string
	HandleFunc func(echo.Context) error
}

type Routers []Router

func RegisterRoutes(e *echo.Echo, routes Routers, apiVersion string) {

	versionedGroups := make(map[string]*echo.Group)

	for _, route := range routes {
		prefix := "/api/" + apiVersion

		group, exists := versionedGroups[prefix]
		if !exists {
			group = e.Group(prefix)
			versionedGroups[prefix] = group
			group.Use(middleware.Authentication())
		}

		switch route.Method {
		case http.MethodGet:
			group.GET(route.Path, route.HandleFunc)
		case http.MethodPost:
			group.POST(route.Path, route.HandleFunc)
		case http.MethodPut:
			group.PUT(route.Path, route.HandleFunc)
		case http.MethodPatch:
			group.PATCH(route.Path, route.HandleFunc)
		case http.MethodDelete:
			group.DELETE(route.Path, route.HandleFunc)
		default:
			e.Logger.Warnf("Unsupported method %s for path %s", route.Method, route.Path)
		}
	}

}
