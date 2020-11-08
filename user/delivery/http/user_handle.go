package http

import (
	"go-gorm-echo/domain"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

//ResponseError return error message
type ResponseError struct {
	Message string `json:"message"`
}

//UserHandler represent the routing handler for User
type UserHandler struct {
	UserUcase domain.UserUseCase
}

//NewUserHandler will initialize the User/Resources endpoint
func NewUserHandler(e *echo.Echo, userUcase domain.UserUseCase) {
	handler := &UserHandler{
		UserUcase: userUcase,
	}

	e.GET("/testing", handler.Testing)
	e.GET("/user", handler.FetchAll)
	e.POST("/user", handler.Store)
	e.GET("/user/:id", handler.GetByID)
	e.PUT("/user/:id", handler.Update)
	e.DELETE("/user/:id", handler.Delete)
	e.POST("/login", handler.Login)
}

//Testing will return string
func (u *UserHandler) Testing(c echo.Context) error {
	return c.JSON(http.StatusOK, "Testing works")
}

//FetchAll will fetched All data
func (u *UserHandler) FetchAll(c echo.Context) error {
	ctx := c.Request().Context()
	data, err := u.UserUcase.FetchAll(ctx)
	if err != nil {
		log.Error(err)
	}

	return c.JSON(http.StatusOK, data)
}

//GetByID get a single data by ID
func (u *UserHandler) GetByID(c echo.Context) error {
	idP, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, "Not found")
	}
	id := int64(idP)

	ctx := c.Request().Context()
	data, err := u.UserUcase.GetByID(ctx, id)
	if err != nil {
		return c.JSON(http.StatusNotFound, ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, data)

}

//Store will store the request param.
func (u *UserHandler) Store(c echo.Context) (err error) {
	var user domain.User
	err = c.Bind(&user)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	ctx := c.Request().Context()

	err = u.UserUcase.Store(ctx, &user)

	if err != nil {
		return c.JSON(http.StatusBadRequest, "Bad Request")
	}

	return c.JSON(http.StatusCreated, user)
}

//Update will process the entity
func (u *UserHandler) Update(c echo.Context) (err error) {
	var user domain.User
	err = c.Bind(&user)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}
	ctx := c.Request().Context()
	idP, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Bad Request")
	}

	id := int64(idP)

	data, err := u.UserUcase.Update(ctx, &user, id)

	return c.JSON(http.StatusOK, data)

}

//Delete the spesific data
func (u *UserHandler) Delete(c echo.Context) (err error) {
	idP, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Bad Request")
	}

	ctx := c.Request().Context()

	id := int64(idP)
	err = u.UserUcase.Delete(ctx, id)

	return c.JSON(http.StatusOK, "Deleted")
}

//Login to generate jwt
func (u *UserHandler) Login(c echo.Context) (err error) {
	var user domain.User
	err = c.Bind(&user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Bad Request")
	}
	ctx := c.Request().Context()

	token, err := u.UserUcase.Login(ctx, &user)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, "not authorized")
	}

	return c.JSON(http.StatusOK, token)
}
