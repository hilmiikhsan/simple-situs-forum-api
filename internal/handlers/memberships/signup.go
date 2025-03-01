package memberships

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hilmiikhsan/situs-forum/internal/model/memberships"
	"github.com/rs/zerolog/log"
)

func (h *Handler) SignUp(c *gin.Context) {
	ctx := c.Request.Context()

	var req memberships.SignUpRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		log.Error().Err(err).Msg("failed to bind json")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err := h.membershipSvc.SignUp(ctx, req)
	if err != nil {
		log.Error().Err(err).Msg("failed to sign up")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.Status(http.StatusCreated)
}
