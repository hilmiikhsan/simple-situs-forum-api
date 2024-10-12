package memberships

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hilmiikhsan/situs-forum/internal/model/memberships"
	"github.com/rs/zerolog/log"
)

func (h *Handler) RefreshToken(c *gin.Context) {
	ctx := c.Request.Context()

	var req memberships.RefreshTokenRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		log.Error().Err(err).Msg("failed to bind json")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	userID := c.GetInt64("user_id")

	accessToken, err := h.membershipSvc.ValidateRefreshToken(ctx, userID, req)
	if err != nil {
		log.Error().Err(err).Msg("failed to validate refresh token")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, memberships.RefreshTokenResponse{
		AccessToken: accessToken,
	})
}
