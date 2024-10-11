package posts

import (
	"context"

	"github.com/hilmiikhsan/situs-forum/internal/model/posts"
	"github.com/rs/zerolog/log"
)

func (r *repository) CreateComment(ctx context.Context, model posts.CommentModel) error {
	_, err := r.db.ExecContext(ctx, queryInsertComment,
		model.PostID,
		model.UserID,
		model.CommentContent,
		model.CreatedAt,
		model.CreatedBy,
		model.UpdatedAt,
		model.UpdatedBy,
	)
	if err != nil {
		log.Error().Err(err).Msg("failed to insert comment")
		return err
	}

	return nil
}
