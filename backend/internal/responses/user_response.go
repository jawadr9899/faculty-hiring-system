package responses

import (
	"encoding/json"
	"github.com/labstack/echo/v5"
)

type UserResponse struct{
	Status int   `json:"status"`
	Success bool  `json:"success"`
	Message any `json:"message"`
}

// Generic Response Streamer
func JsonResponse[T comparable](c *echo.Context, statusCode int,data... T ) error {
	c.Response().Header().Set(echo.HeaderContentType,echo.MIMEApplicationJSON)
	c.Response().WriteHeader(statusCode)
	return json.NewEncoder(c.Response()).Encode(data)
}

