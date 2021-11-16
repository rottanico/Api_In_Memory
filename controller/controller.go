package controller

import (
	"DB_In_Memory/methods"
	stre "DB_In_Memory/structures"
	"net/http"
	"strconv"

	"github.com/go-playground/form"

	"github.com/labstack/echo"
)

type Controller struct {
}

func (r *Controller) Get_Courses(context echo.Context) error {
	result := methods.GetAllDB()
	return context.JSON(http.StatusOK, result)
}
func (r *Controller) Post_Course(context echo.Context) error {
	c := new(stre.Course)

	if context.FormValue("name") != "" { //verifica que los datos viajan en tipo form value o en el cuerpo de la peticion
		decoder := form.NewDecoder()
		form, _ := context.FormParams()

		if err := decoder.Decode(&c, form); err != nil {
			return context.JSON(http.StatusBadRequest, err)
		}
	} else {

		if err := context.Bind(&c); err != nil {
			return context.JSON(http.StatusBadRequest, err)
		}
	}

	if res := methods.PostCourseDB(*c); res {
		return context.String(http.StatusOK, "successful order")
	}
	return context.String(http.StatusInternalServerError, "method not performed")
}
func (c *Controller) Get_Course(context echo.Context) error {
	id, err := strconv.Atoi(context.QueryParam("id"))
	if err != nil {
		return context.String(http.StatusBadRequest, "id must be a number")
	}

	if value, state := methods.GetCourseDB(id); state {
		return context.JSON(http.StatusOK, value)
	} else {

		return context.String(http.StatusInternalServerError, "method not performed")
	}

}
func (c *Controller) Delete_Course(context echo.Context) error {
	id, err := strconv.Atoi(context.QueryParam("id"))
	if err != nil {
		return context.String(http.StatusBadRequest, "id must be a number")
	}
	if status := methods.DeleteCourseDB(id); status {
		return context.String(http.StatusOK, "element has been deleted")
	} else {
		return context.String(http.StatusInternalServerError, "method not performed")
	}

}
func (r *Controller) Put_Course(context echo.Context) error {
	id, err := strconv.Atoi(context.QueryParam("id"))

	if err != nil {
		return context.String(http.StatusBadRequest, "id must be a number")
	}
	course := new(stre.Course)

	if context.FormValue("name") != "" { //verifica que los datos viajan en tipo form value o en el cuerpo de la peticion
		decoder := form.NewDecoder()
		form, _ := context.FormParams()
		if err := decoder.Decode(&course, form); err != nil {
			return context.JSON(http.StatusBadRequest, err)
		}
	} else {
		if err := context.Bind(&course); err != nil {
			return context.JSON(http.StatusBadRequest, err)
		}
	}
	if res := methods.PutCourseDB(*course, id); res {
		return context.String(http.StatusOK, "element has been modified")
	} else {
		return context.String(http.StatusInternalServerError, "method not performed")
	}

}
