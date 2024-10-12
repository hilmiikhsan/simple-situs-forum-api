package posts

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func (h *Handler) GetPostByID(c *gin.Context) {
	ctx := c.Request.Context()

	postIDStr := c.Param("post_id")

	postID, err := strconv.ParseInt(postIDStr, 10, 64)
	if err != nil {
		log.Error().Err(err).Msg("failed to parse post_id")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	response, err := h.postSvc.GetPostByID(ctx, postID)
	if err != nil {
		log.Error().Err(err).Msg("failed to get post by id")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response)
}
