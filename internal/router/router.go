package router

import (
	"github.com/labstack/echo"
)

type Router struct {
	Name       string
	Method     string
	Path       string
	HandleFunc func(echo.Context) error
}

type Routes []Router

func RegisterRoutes(e *echo.Echo, routes Routes, apiVersion string) {

	versionedGroups := make(map[string]*echo.Group)

	for _, route := range routes {

		prefix := "/api/" + apiVersion

		group, exits := versionedGroups[prefix]

		if !exits {
			group = e.Group(prefix)
			versionedGroups[prefix] = group
		}
		switch route.Method {
		case "GET":
			group.GET(route.Path, route.HandleFunc)
		case "POST":
			group.POST(route.Path, route.HandleFunc)
		case "PUT":
			group.PUT(route.Path, route.HandleFunc)
		case "PATCH":
			group.PATCH(route.Path, route.HandleFunc)
		case "DELETE":
			group.DELETE(route.Path, route.HandleFunc)
		default:
			e.Logger.Warnf("Unsupported method %s for path %s", route.Method, route.Path)
		}
	}
}
