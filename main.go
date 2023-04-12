package main

import (
	"ganasa18/belajar-golang-restful-api/app"
	"ganasa18/belajar-golang-restful-api/controller"
	"ganasa18/belajar-golang-restful-api/exception"
	"ganasa18/belajar-golang-restful-api/helper"
	"ganasa18/belajar-golang-restful-api/middleware"
	"ganasa18/belajar-golang-restful-api/repository"
	"ganasa18/belajar-golang-restful-api/routes"
	"ganasa18/belajar-golang-restful-api/service"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-playground/validator"
	"github.com/julienschmidt/httprouter"
)

func main() {

	db := app.NewDB()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := httprouter.New()
	routes.ApiRouteCategory(router, categoryController)

	router.PanicHandler = exception.ErrorHandler

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)

}
