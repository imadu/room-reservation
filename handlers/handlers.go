package handlers

import (
	"github.com/labstack/echo"
	"net/http"
)

type Response struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

type ErrorResponseBody struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Status  string `json:"status"`
}

type ErrorResponse struct {
	Error ErrorResponseBody `json:"error"`
}

func sendSuccess(c echo.Context, data interface{}) error {
	s := Response{}
	s.Status = "success"
	s.Data = data
	return c.JSON(http.StatusOK, s)
}

func sendError(c echo.Context, code string, message string, status string) error {
	e := ErrorResponseBody{}
	e.Status = status
	e.Code = code
	e.Message = message

	s := ErrorResponse{}
	s.Error = e
	return c.JSON(http.StatusOK, s)
}

func sendData(c echo.Context, data interface{}) error {
	return c.JSON(http.StatusOK, data)
}
