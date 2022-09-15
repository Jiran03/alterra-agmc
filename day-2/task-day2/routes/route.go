package routes

import (
	"github.com/Jiran03/agmc/task/day2/controllers"
	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()
	v1 := e.Group("/v1")

	bookGroup := v1.Group("/books")
	bookGroup.POST("", controllers.CreateBookController)
	bookGroup.GET("", controllers.GetAllBookController)
	bookGroup.GET("/:id", controllers.GetBookByIdController)
	bookGroup.PUT("/:id", controllers.UpdateBookController)
	bookGroup.DELETE("/:id", controllers.DeleteBookController)

	userGroup := v1.Group("/users")
	userGroup.POST("", controllers.CreateUserController)
	userGroup.GET("", controllers.GetAllUserController)
	userGroup.GET("/:id", controllers.GetUserByIdController)
	userGroup.PUT("/:id", controllers.UpdateUserController)
	userGroup.DELETE("/:id", controllers.DeleteUserController)

	return e
}
