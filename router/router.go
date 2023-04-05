package router

import (
	"go-clean-architecture/controller"

	"github.com/labstack/echo/v4"
)

func NewRouter(uc controller.IUserController) *echo.Echo {
	e := echo.New()
	e.POST("/signup", uc.SignUp)
	e.POST("/signin", uc.SignIn)
	e.POST("/signout", uc.SignOut)
	return e
}
