package posts

import (
	"context"

	"github.com/hilmiikhsan/situs-forum/internal/model/posts"
	"github.com/rs/zerolog/log"
)

func (s *service) GetPostByID(ctx context.Context, postID int64) (*posts.GetPostResponse, error) {
	data, err := s.postRepo.GetPostByID(ctx, postID)
	if err != nil {
		log.Error().Err(err).Msg("failed to get post by id")
		return nil, err
	}

	likeCount, err := s.postRepo.CountLikeByPostID(ctx, postID)
	if err != nil {
		log.Error().Err(err).Msg("failed to count like by post id")
		return nil, err
	}

	comments, err := s.postRepo.GetCommentByPostID(ctx, postID)
	if err != nil {
		log.Error().Err(err).Msg("failed to get comment by post id")
		return nil, err
	}

	return &posts.GetPostResponse{
		Data: posts.Post{
			ID:           data.ID,
			UserID:       data.UserID,
			Username:     data.Username,
			PostTitle:    data.PostTitle,
			PostContent:  data.PostContent,
			PostHashtags: data.PostHashtags,
			IsLiked:      data.IsLiked,
		},
		LikeCount: likeCount,
		Comments:  comments,
	}, nil
}
