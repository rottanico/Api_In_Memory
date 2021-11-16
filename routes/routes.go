package routes

import (
	"DB_In_Memory/controller"

	"github.com/labstack/echo"
)

func Router(e *echo.Echo) {
	router := controller.Controller{}
	e.POST("Course", router.Post_Course)
	e.GET("/Course", router.Get_Course)
	e.GET("/Courses", router.Get_Courses)
	e.DELETE("/Course", router.Delete_Course)
}
