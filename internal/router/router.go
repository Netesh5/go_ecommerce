package router

import "github.com/labstack/echo"

type Router struct {
	Name       string
	Method     string
	Path       string
	HandleFunc func(*echo.Context)
}
