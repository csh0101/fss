package http

import (
	v1 "fss/api/product/app/v1"
	"net/http"

	"github.com/labstack/echo"
)

var (
	apiV1 = "/api/v1"
)

// Register
func Register(echo *echo.Echo, srv v1.TextHTTPServer) {
	// v1 /api/version/resource/action
	echo.Add(http.MethodGet, apiV1+"/text/get", TextQueryHandler(srv))
	// default global
	echo.Any("/*", NotFound())
}

func TextQueryHandler(srv v1.TextHTTPServer) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		resp, err := srv.QueryTextByFilter(ctx.Request().Context(), ctx.Request())
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, resp)
		}
		return ctx.JSON(http.StatusOK, resp)
	}
}

func NotFound() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		return ctx.String(http.StatusNotFound, "resource is not exist")
	}
}
