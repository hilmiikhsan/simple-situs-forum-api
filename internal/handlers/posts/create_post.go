package posts

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hilmiikhsan/situs-forum/internal/model/posts"
	"github.com/rs/zerolog/log"
)

func (h *Handler) CreatePost(c *gin.Context) {
	ctx := c.Request.Context()

	var req posts.CreatePostRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		log.Error().Err(err).Msg("failed to bind json")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	userID := c.GetInt64("user_id")

	err := h.postSvc.CreatePost(ctx, userID, req)
	if err != nil {
		log.Error().Err(err).Msg("failed to create post")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.Status(http.StatusCreated)
}
