package util

import "github.com/labstack/echo/v4"

func ErrorResponse(code int, message string) echo.Map {
	return echo.Map{
		"success": false,
		"code":    code,
		"error":   message,
	}
}

func SuccessResponse(data interface{}) echo.Map {
	return echo.Map{
		"success": true,
		"code":    200,
		"data":    data,
	}
}
