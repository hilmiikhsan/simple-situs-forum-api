package memberships

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hilmiikhsan/situs-forum/internal/model/memberships"
	"github.com/rs/zerolog/log"
)

func (h *Handler) Login(c *gin.Context) {
	ctx := c.Request.Context()

	var req memberships.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		log.Error().Err(err).Msg("failed to bind json")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	accessToken, err := h.membershipSvc.Login(ctx, req)
	if err != nil {
		log.Error().Err(err).Msg("failed to login")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	response := memberships.LoginResponse{
		AccessToken: accessToken,
	}

	c.JSON(http.StatusOK, response)
}
