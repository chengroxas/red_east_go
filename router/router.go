package router

import (
	"red-east/controller/user"

	"github.com/gin-gonic/gin"
)

func RegisterRouter(r *gin.Engine) {
	userRouter := r.Group("/user")
	{
		userController := &user.UserController{}
		userRouter.POST("/bypass", userController.LoginByPass)
		userRouter.POST("/bysms", userController.LoginBySms)
		userRouter.GET("/userlist", userController.GetUserList)
	}
	// v1 := r.Group("/v1")
	// {
	// 	userController := &user.User{}
	// 	userRouter := v1.Group("/user")
	// 	{
	// 		userRouter.POST("/login", userController.Login)
	// 		userRouter.POST("/register", userController.Register)
	// 		userRouter.GET("/userinfo", userController.GetInfo)
	// 	}
	// 	touristController := &user.Tourist{}
	// 	touristRouter := v1.Group("/tourist")
	// 	{
	// 		touristRouter.POST("/login", touristController.Login)
	// 	}
	// }
}
