package app

import (
	"github.com/julienschmidt/httprouter"
	"github.com/raviMukti/training-golaang-restful-api/controller"
	"github.com/raviMukti/training-golaang-restful-api/exception"
)

func NewRouter(categoryController controller.CategoryController) *httprouter.Router {
	router := httprouter.New()

	// Register Route
	router.GET("/api/categories", categoryController.FindAll)
	router.GET("/api/categories/:categoryId", categoryController.FindById)
	router.POST("/api/categories", categoryController.Create)
	router.PUT("/api/categories/:categoryId", categoryController.Update)
	router.DELETE("/api/categories/:categoryId", categoryController.Delete)

	// Register Default Error Handler
	router.PanicHandler = exception.ErrorHandler

	return router
}
