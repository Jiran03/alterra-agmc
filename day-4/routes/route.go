package routes

import (
	"os"

	"github.com/Jiran03/agmc/task/day4/controllers"
	"github.com/Jiran03/agmc/task/day4/middlewares"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	e := echo.New()
	middlewares.LogMiddleware(e)
	v1 := e.Group("/v1")
	authGroup := v1.Group("/auth")

	bookGroup := v1.Group("/books")
	//not authenticated
	bookGroup.GET("", controllers.GetAllBookController)
	bookGroup.GET("/:id", controllers.GetBookByIdController)
	//authenticated
	authBookGroup := authGroup.Group("/books")
	authBookGroup.Use(middleware.JWT([]byte(os.Getenv("JWT_SECRET"))))
	authBookGroup.POST("", controllers.CreateBookController)
	authBookGroup.PUT("/:id", controllers.UpdateBookController)
	authBookGroup.DELETE("/:id", controllers.DeleteBookController)

	userGroup := v1.Group("/users")
	//not authenticated
	userGroup.POST("", controllers.CreateUserController)
	userGroup.POST("/login", controllers.LoginUserController)
	//authenticated
	authUserGroup := authGroup.Group("/users")
	authUserGroup.Use(middleware.JWT([]byte(os.Getenv("JWT_SECRET"))))
	authUserGroup.GET("", controllers.GetAllUserController)
	authUserGroup.GET("/:id", controllers.GetUserByIdController)
	authUserGroup.PUT("/:id", controllers.UpdateUserController)
	authUserGroup.DELETE("/:id", controllers.DeleteUserController)

	return e
}
