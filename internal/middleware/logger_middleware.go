package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func init() {

	logrus.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
		PrettyPrint:     true,
	})

}

func LogrusLogger() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			// Read request body
			var bodyBytes []byte
			if c.Request().Body != nil {
				bodyBytes, _ = io.ReadAll(c.Request().Body)
				c.Request().Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
			}
			fmt.Println()
			fmt.Println(strings.Repeat("=", 80))
			fmt.Println("🚀 URL : ", c.Request().URL.String())
			fmt.Println(strings.Repeat("-", 80))
			err := next(c)

			fields := logrus.Fields{
				"method":       c.Request().Method,
				"status":       c.Response().Status,
				"remote_ip":    c.RealIP(),
				"query_params": c.Request().URL.Query(),
				"host":         c.Request().URL.Host,
				"host_name":    c.Request().URL.Hostname(),
			}

			// Add error if present
			if err != nil {
				fields["error"] = err.Error()
			}

			// Add query parameters
			if len(c.QueryParams()) > 0 {
				fields["query_params"] = c.QueryParams()
			}

			// Add path parameters
			pathParams := make(map[string]string)
			for _, name := range c.ParamNames() {
				pathParams[name] = c.Param(name)
			}
			if len(pathParams) > 0 {
				fields["path_params"] = pathParams
			}

			// Add request body
			if len(bodyBytes) > 0 {
				var bodyData interface{}
				if json.Unmarshal(bodyBytes, &bodyData) == nil {
					fields["request_body"] = bodyData
				} else {
					fields["request_body"] = string(bodyBytes)
				}
			}

			logrus.WithFields(fields).Info("Request processed")
			fmt.Println(strings.Repeat("=", 80))
			fmt.Println()
			return err
		}
	}
}
