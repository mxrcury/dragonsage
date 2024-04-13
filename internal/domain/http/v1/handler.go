package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/mxrcury/dragonsage/internal/service"
)

type Handlers struct {
	usersHandler Users

	group *gin.RouterGroup
}

type Deps struct {
	Router *gin.Engine

	Services *service.Services
}

type (
	baseHandler interface {
		init(*gin.RouterGroup)
	}

	Users interface {
		baseHandler

		signUp(*gin.Context)
	}
)

func NewHandlers(deps *Deps) *Handlers {
	usersHandler := NewUsersHandler("/users", deps.Services.UsersService)

	group := deps.Router.Group("/v1")

	return &Handlers{usersHandler: usersHandler, group: group}
}

func (h *Handlers) Init() {
	h.usersHandler.init(h.group)
}
