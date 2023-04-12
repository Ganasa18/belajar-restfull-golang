package routes

import (
	"fmt"
	"ganasa18/belajar-golang-restful-api/controller"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func ApiRouteCategory(router *httprouter.Router, categoryController controller.CategoryController) {

	router.GET("/api/categories", categoryController.FindAll)
	router.GET("/api/categories/:categoryId", categoryController.FindById)
	router.POST("/api/categories", categoryController.Create)
	router.PUT("/api/categories/:categoryId", categoryController.Update)
	router.DELETE("/api/categories/:categoryId", categoryController.Delete)
}

func ApiRouteTest(router *httprouter.Router) {
	router.GET("/api/test", func(writer http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		fmt.Fprint(writer, "HELLO TEST")
	})
}
