package user

import (
	"clean/controllers/user/request"
	"clean/controllers/user/response"
	userEntities "clean/entities/user"
	"clean/utilities/base"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	userUseCase userEntities.UseCaseInterface
}

func NewUserController(userUseCase userEntities.UseCaseInterface) *UserController {
	return &UserController{
		userUseCase: userUseCase,
	}
}

func (userController *UserController) Register(c echo.Context) error {
	var userFromRequest request.UserRegisterRequest
	c.Bind(&userFromRequest)

	userEntities := userEntities.User{
		Name:     userFromRequest.Name,
		Email:    userFromRequest.Email,
		Password: userFromRequest.Password,
	}

	newUser, err := userController.userUseCase.Register(&userEntities)
	if err != nil {
		return c.JSON(base.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	userResponse := response.UserRegisterResponse{
		Id:    newUser.Id,
		Name:  newUser.Name,
		Email: newUser.Email,
	}
	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success Register", userResponse))
}

func (userController *UserController) Login(c echo.Context) error {
	var userFromRequest request.UserLoginRequest
	c.Bind(&userFromRequest)

	userEntities := userEntities.User{
		Email:    userFromRequest.Email,
		Password: userFromRequest.Password,
	}

	userFromDb, err := userController.userUseCase.Login(&userEntities)
	if err != nil {
		return c.JSON(base.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	userResponse := response.UserLoginResponse{
		Id:    userFromDb.Id,
		Name:  userFromDb.Name,
		Email: userFromDb.Email,
		Token: userFromDb.Token,
	}
	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success Login", userResponse))
}