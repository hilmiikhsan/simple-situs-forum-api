package memberships

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/hilmiikhsan/situs-forum/internal/middleware"
	"github.com/hilmiikhsan/situs-forum/internal/model/memberships"
)

type membershipService interface {
	SignUp(ctx context.Context, req memberships.SignUpRequest) error
	Login(ctx context.Context, req memberships.LoginRequest) (string, string, error)
	ValidateRefreshToken(ctx context.Context, userID int64, req memberships.RefreshTokenRequest) (string, error)
}

type Handler struct {
	*gin.Engine

	membershipSvc membershipService
}

func NewHandler(api *gin.Engine, membershipSVc membershipService) *Handler {
	return &Handler{
		Engine:        api,
		membershipSvc: membershipSVc,
	}
}

func (h *Handler) RegisterRoute() {
	route := h.Group("/memberships")

	route.GET("/ping", h.Ping)
	route.POST("/signup", h.SignUp)
	route.POST("/login", h.Login)

	routeRefreshToken := route.Group("/")
	routeRefreshToken.Use(middleware.AuthRefreshTokenMiddleware())

	routeRefreshToken.POST("/refresh-token", h.RefreshToken)
}
