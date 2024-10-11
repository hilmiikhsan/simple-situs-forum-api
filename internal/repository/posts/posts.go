package posts

import (
	"context"

	"github.com/rs/zerolog/log"

	"github.com/hilmiikhsan/situs-forum/internal/model/posts"
)

func (r *repository) CreatePost(ctx context.Context, model posts.PostModel) error {
	_, err := r.db.ExecContext(ctx, queryInsertPosts,
		model.UserID,
		model.PostTitle,
		model.PostContent,
		model.PostHastags,
		model.CreatedAt,
		model.CreatedBy,
		model.UpdatedAt,
		model.UpdatedBy,
	)
	if err != nil {
		log.Error().Err(err).Msg("failed to insert post")
		return err
	}

	return nil
}
