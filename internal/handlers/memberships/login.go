package memberships

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hilmiikhsan/situs-forum/internal/model/memberships"
	"github.com/sirupsen/logrus"
)

func (h *Handler) Login(c *gin.Context) {
	ctx := c.Request.Context()

	var req memberships.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		logrus.Error("failed to bind json: ", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	accessToken, err := h.membershipSvc.Login(ctx, req)
	if err != nil {
		logrus.Error("failed to login: ", err)
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
