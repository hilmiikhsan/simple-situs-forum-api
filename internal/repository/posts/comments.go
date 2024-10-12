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

func (r *repository) GetCommentByPostID(ctx context.Context, postID int64) ([]posts.Comments, error) {
	rows, err := r.db.QueryContext(ctx, queryGetCommentByPostID, postID)
	if err != nil {
		log.Error().Err(err).Msg("failed to query comment")
		return nil, err
	}
	defer rows.Close()

	responses := make([]posts.Comments, 0)

	for rows.Next() {
		var (
			comments posts.Comments
			username string
		)

		err = rows.Scan(
			&comments.ID,
			&comments.UserID,
			&username,
			&comments.CommentContent,
		)
		if err != nil {
			log.Error().Err(err).Msg("failed to scan comment")
			return nil, err
		}

		responses = append(responses, posts.Comments{
			ID:             comments.ID,
			UserID:         comments.UserID,
			Username:       username,
			CommentContent: comments.CommentContent,
		})
	}

	return responses, nil
}
