package controller

import (
	"DB_In_Memory/methods"
	stre "DB_In_Memory/structures"
	"fmt"
	"net/http"

	"github.com/go-playground/form"

	"github.com/labstack/echo"
)

type Controller struct {
}

func (r *Controller) Get_All(context echo.Context) error {
	result := methods.GetAllDB()
	return context.JSON(http.StatusOK, result)
}
func (r *Controller) Post_Course(context echo.Context) error {
	c := new(stre.Course)

	if context.FormValue("name") != "" { //verifica que los datos viajan en tipo form value o en el cuerpo de la peticion
		decoder := form.NewDecoder()
		form, _ := context.FormParams()
		err := decoder.Decode(&c, form)

		if err != nil {
			fmt.Println("entró aca")
			return context.JSON(http.StatusBadRequest, err)
		}
	} else {
		err := context.Bind(&c)
		if err != nil {
			fmt.Println("o aca?")
			return context.JSON(http.StatusBadRequest, err)
		}
	}
	res := methods.PostCourseDB(*c)
	if res {
		return context.String(http.StatusOK, "successful order")
	}
	return context.String(http.StatusInternalServerError, "method not performed")
}
