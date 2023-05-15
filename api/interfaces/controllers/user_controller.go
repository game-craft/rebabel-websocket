package controllers

import (
	"fmt"
	"strconv"
	"strings"
	"time"
	"golang.org/x/crypto/bcrypt"
	"github.com/labstack/echo"
	"github.com/dgrijalva/jwt-go"
	
	"docker-echo-template/api/domain"
	"docker-echo-template/api/interfaces/database"
	"docker-echo-template/api/usecase"
)

type UserController struct {
	Interactor usecase.UserInteractor
}

type JwtClaims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

func NewUserController(sqlHandler database.SqlHandler) *UserController {
	return &UserController{
		Interactor: usecase.UserInteractor{
			UserRepository: &database.UserRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *UserController) Users(c echo.Context) (err error) {
	users, err := controller.Interactor.Users()
	if err != nil {
		return c.JSON(500, NewError(err))
	}

	return c.JSON(200, users)
}

func (controller *UserController) Register(c echo.Context) (err error) {
	u := domain.User{}
	c.Bind(&u)
	user, err := controller.Interactor.UserByEmail(u.Email)
	if err == nil {
		return c.JSON(500, NewError(err))
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	u.Password = string(hashedPassword)
	u.Status = "Offline"
	user, err = controller.Interactor.Add(u)
	if err != nil {
		return c.JSON(500, NewError(err))
	}

	return c.JSON(201, user)
}

func (controller *UserController) Login(c echo.Context) (err error) {
	u := domain.User{}
	c.Bind(&u)
	email := u.Email
	user, err := controller.Interactor.UserByEmail(email)
	if err != nil {
		return c.JSON(500, NewError(err))
	}
	
    err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(u.Password))
	if err != nil {
		return c.JSON(500, NewError(err))
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = user.ID
	claims["email"] = user.Email
	claims["username"] = user.Username
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	tokenString, _ := token.SignedString([]byte("secret"))

	return c.JSON(200, map[string]string{"token": tokenString})
}

func (controller *UserController) Check(c echo.Context) (err error) {
	tokenString := c.Request().Header.Get("Authorization")

	if tokenString == "" {
		return c.JSON(500, map[string]string{"message": "Missing token"})
	}

	if len(tokenString) < 7 || strings.ToLower(tokenString[0:6]) != "bearer" {
		return c.JSON(500, map[string]string{"message": "Invalid token format"})
	}
	  
	tokenString = tokenString[7:]
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Invalid token")
		}

		return []byte("secret"), nil
	})

	if err != nil {
		return c.JSON(500, map[string]string{"message": err.Error()})
	}
  
	if !token.Valid {
	  	return c.JSON(500, map[string]string{"message": "Invalid token"})
	}
  
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
	  	return c.JSON(500, map[string]string{"message": "Invalid token"})
	}
  
	userId := claims["id"].(float64)
	email := claims["email"].(string)
	username := claims["username"].(string)
  
	return c.JSON(200, map[string]interface{} {
		"email": email,
		"user_id": userId,
		"user_name": username,
	})
}

func (controller *UserController) Save(c echo.Context) (err error) {
	u := domain.User{}
	c.Bind(&u)
	user, err := controller.Interactor.Update(u)
	if err != nil {
		return c.JSON(500, NewError(err))
	}

	return c.JSON(201, user)
}

func (controller *UserController) Delete(c echo.Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))
	user := domain.User{
		ID: id,
	}
	err = controller.Interactor.DeleteById(user)
	if err != nil {
		return c.JSON(500, NewError(err))
	}

	return c.JSON(200, user)
}
