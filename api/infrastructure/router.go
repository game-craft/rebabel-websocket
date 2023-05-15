package infrastructure

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	
	"docker-echo-template/api/interfaces/controllers"
)

func Init() {
	userController := controllers.NewUserController(NewSqlHandler())
	chatController := controllers.NewChatController(NewSqlHandler())
	positionController := controllers.NewPositionController(NewSqlHandler())

	e := echo.New()
	
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// auth
	e.GET("/users", func(c echo.Context) error { return userController.Users(c) })
	e.POST("/user_register", func(c echo.Context) error { return userController.Register(c) })
	e.POST("/user_login", func(c echo.Context) error { return userController.Login(c) })
	e.GET("/user_check", func(c echo.Context) error { return userController.Check(c) })
	e.PUT("/user/:id", func(c echo.Context) error { return userController.Save(c) })
	e.DELETE("/user/:id", func(c echo.Context) error { return userController.Delete(c) })

	// chat
	e.GET("/room_chat/:roomId", func(c echo.Context) error { return chatController.RoomChat(c) })

	// position
	e.GET("/room_position/:roomId", func(c echo.Context) error { return positionController.RoomPosition(c) })

	e.Logger.Fatal(e.Start(":8001"))
}
