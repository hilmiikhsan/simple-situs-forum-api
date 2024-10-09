package memberships

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hilmiikhsan/situs-forum/internal/model/memberships"
	"github.com/sirupsen/logrus"
)

func (h *Handler) SignUp(c *gin.Context) {
	ctx := c.Request.Context()

	var req memberships.SignUpRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		logrus.Error("failed to bind json: ", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err := h.membershipSvc.SignUp(ctx, req)
	if err != nil {
		logrus.Error("failed to sign up: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.Status(http.StatusCreated)
}
