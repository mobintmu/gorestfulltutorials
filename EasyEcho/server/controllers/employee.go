package controllers

import (
	"net/http"

	"../models"

	"github.com/labstack/echo"
)

//GetEmployees get all list of employee
func GetEmployees(c echo.Context) error {
	result := models.GetEmployee()
	println("foo")
	return c.JSON(http.StatusOK, result)
}
