package main

import (
	"net/http"

	"github.com/faridlan/belajar-restful-api/app"
	"github.com/faridlan/belajar-restful-api/controller"
	"github.com/faridlan/belajar-restful-api/helper"
	"github.com/faridlan/belajar-restful-api/middleware"
	"github.com/faridlan/belajar-restful-api/repository"
	"github.com/faridlan/belajar-restful-api/service"
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
		Addr:    ":3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
