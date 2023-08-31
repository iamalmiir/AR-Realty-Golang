package utils

import (
	"github.com/gofiber/fiber/v2"
)

var responseCodes = map[int]string{
	200: "OK",
	400: "Bad Request",
	401: "Unauthorized",
	403: "Forbidden",
	404: "Not Found",
	405: "Method Not Allowed",
	406: "Not Acceptable",
	407: "Proxy Authentication Required",
	408: "Request Timeout",
	409: "Conflict",
	500: "Internal Server Error",
}

func ServerResponse(statusCode int, message string, data ...interface{}) fiber.Map {
	return fiber.Map{
		"status":  responseCodes[statusCode],
		"message": message,
		"data":    data,
	}
}
