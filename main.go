package main

import (
	"learn-go-restful-api/app"
	"learn-go-restful-api/controller"
	"learn-go-restful-api/helper"
	"learn-go-restful-api/middleware"
	"learn-go-restful-api/repository"
	"learn-go-restful-api/service"
	"net/http"

	"github.com/go-playground/validator/v10"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db := app.NewDB()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)
	router := app.NewRouter(categoryController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleWare(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
