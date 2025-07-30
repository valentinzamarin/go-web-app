package rest

import (
	"backend/internal/app/interfaces"
	"backend/internal/domain/entities"
	"backend/internal/interfaces/api/rest/dto/request"
	"backend/internal/interfaces/api/rest/dto/response"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type UserController struct {
	service interfaces.UserService
}

func NewUserController(e *echo.Echo, service interfaces.UserService) *UserController {

	controller := &UserController{
		service: service,
	}

	e.POST("/api/create", controller.CreateUserController)
	e.GET("/api/users", controller.GetUsersController)

	return controller

}

func (uc *UserController) CreateUserController(c echo.Context) error {
	var userRequest request.CreateUserRequest

	if err := c.Bind(&userRequest); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	newUser := entities.NewUser(
		uuid.New(),
		userRequest.Username,
		userRequest.Password,
		userRequest.Role,
	)

	if err := uc.service.CreateNewUser(newUser); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	response := response.UserResponse{
		ID:       newUser.ID().String(),
		Username: newUser.Username(),
		Role:     newUser.Role(),
	}

	return c.JSON(http.StatusCreated, response)
}

func (pc *UserController) GetUsersController(c echo.Context) error {
	users, err := pc.service.GetAllUsers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to fetch users",
		})
	}

	var usersResponse []response.UserResponse
	for _, user := range users {
		usersResponse = append(usersResponse, response.UserResponse{
			ID:       user.ID().String(),
			Username: user.Username(),
			Role:     user.Role(),
		})
	}

	return c.JSON(http.StatusOK, usersResponse)
}
