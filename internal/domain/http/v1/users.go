package v1

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mxrcury/dragonsage/internal/service"
)

type UsersHandler struct {
	service *service.UsersService
	path    string
}

func NewUsersHandler(path string, service *service.UsersService) Users {
	return &UsersHandler{path: path, service: service}
}

func (h *UsersHandler) init(group *gin.RouterGroup) {
	usersGroup := group.Group(h.path)
	{
		usersGroup.POST("/sign-up", h.signUp)
	}
}

func (h *UsersHandler) signUp(ctx *gin.Context) {
	var userInput *service.UsersSignUpInput
	fmt.Println(userInput)
	if err := ctx.BindJSON(userInput); err != nil {
		ctx.JSON(400, map[string]string{
			"message": "You entered wrong input",
		})
	}

	if err := h.service.SignUp(userInput); err != nil {
		ctx.JSON(400, map[string]string{
			"message": fmt.Sprintf("%s\n", err),
		})
	}

	ctx.Status(http.StatusCreated)
}
