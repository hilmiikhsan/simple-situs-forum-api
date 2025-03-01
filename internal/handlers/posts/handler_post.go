package posts

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/hilmiikhsan/situs-forum/internal/middleware"
	"github.com/hilmiikhsan/situs-forum/internal/model/posts"
)

type postService interface {
	CreatePost(ctx context.Context, userID int64, req posts.CreatePostRequest) error
	CreateComment(ctx context.Context, postID, userID int64, req posts.CreateCommentRequest) error
	UpsertUserActivity(ctx context.Context, postID, userID int64, req posts.UserActivityRequest) error
	GetAllPost(ctx context.Context, pageSize, pageIndex int) (posts.GetAllPostResponse, error)
	GetPostByID(ctx context.Context, postID int64) (*posts.GetPostResponse, error)
}

type Handler struct {
	*gin.Engine

	postSvc postService
}

func NewHandler(api *gin.Engine, postSVc postService) *Handler {
	return &Handler{
		Engine:  api,
		postSvc: postSVc,
	}
}

func (h *Handler) RegisterRoute() {
	route := h.Group("/posts")
	route.Use(middleware.AuthMiddleware())

	route.POST("/create", h.CreatePost)
	route.POST("/comment/:post_id", h.CreateComment)
	route.POST("/activity/:post_id", h.UpsertUserActivity)
	route.GET("/", h.GetAllPost)
	route.GET("/:post_id", h.GetPostByID)
}
