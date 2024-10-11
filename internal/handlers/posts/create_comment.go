package posts

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hilmiikhsan/situs-forum/internal/model/posts"
	"github.com/rs/zerolog/log"
)

func (h *Handler) CreateComment(c *gin.Context) {
	ctx := c.Request.Context()

	var req posts.CreateCommentRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		log.Error().Err(err).Msg("failed to bind json")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	postIDStr := c.Param("post_id")
	postID, err := strconv.ParseInt(postIDStr, 10, 64)
	if err != nil {
		log.Error().Err(err).Msg("failed to parse post_id")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	userID := c.GetInt64("user_id")

	err = h.postSvc.CreateComment(ctx, postID, userID, req)
	if err != nil {
		log.Error().Err(err).Msg("failed to create comment")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.Status(http.StatusCreated)
}
