package posts

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func (h *Handler) GetAllPost(c *gin.Context) {
	ctx := c.Request.Context()

	pageIndexStr := c.Query("page")
	pageSizeStr := c.Query("limit")

	pageIndex, err := strconv.Atoi(pageIndexStr)
	if err != nil {
		log.Error().Err(err).Msg("failed to convert page index")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errors.New("failed to convert page index").Error(),
		})
		return
	}

	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil {
		log.Error().Err(err).Msg("failed to convert page size")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errors.New("failed to convert page size").Error(),
		})
		return
	}

	responses, err := h.postSvc.GetAllPost(ctx, pageSize, pageIndex)
	if err != nil {
		log.Error().Err(err).Msg("failed to get all post")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, responses)
}
