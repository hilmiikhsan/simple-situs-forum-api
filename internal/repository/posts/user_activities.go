package posts

import (
	"context"
	"database/sql"

	"github.com/hilmiikhsan/situs-forum/internal/model/posts"
	"github.com/rs/zerolog/log"
)

func (r *repository) GetUserActivity(ctx context.Context, model posts.UserActivityModel) (*posts.UserActivityModel, error) {
	var responses posts.UserActivityModel

	row := r.db.QueryRowContext(ctx, queryGetUserActivity,
		model.PostID,
		model.UserID,
	)

	err := row.Scan(
		&responses.ID,
		&responses.PostID,
		&responses.UserID,
		&responses.IsLiked,
		&responses.CreatedAt,
		&responses.CreatedBy,
		&responses.UpdatedAt,
		&responses.UpdatedBy,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Error().Err(err).Msg("user activity not found")
			return nil, nil
		}

		log.Error().Err(err).Msg("failed to get user activity")
		return nil, err
	}

	return &responses, nil
}

func (r *repository) CreateUserActivity(ctx context.Context, model posts.UserActivityModel) error {
	_, err := r.db.ExecContext(ctx, queryInsertUserActivity,
		model.PostID,
		model.UserID,
		model.IsLiked,
		model.CreatedAt,
		model.CreatedBy,
		model.UpdatedAt,
		model.UpdatedBy,
	)
	if err != nil {
		log.Error().Err(err).Msg("failed to insert user activity")
		return err
	}

	return nil
}

func (r *repository) UpdateUserActivity(ctx context.Context, model posts.UserActivityModel) error {
	_, err := r.db.ExecContext(ctx, queryUpdateUserActivity,
		model.IsLiked,
		model.UpdatedAt,
		model.UpdatedBy,
		model.PostID,
		model.UserID,
	)
	if err != nil {
		log.Error().Err(err).Msg("failed to update user activity")
		return err
	}

	return nil
}

func (r *repository) CountLikeByPostID(ctx context.Context, postID int64) (int, error) {
	var count int

	row := r.db.QueryRowContext(ctx, queryCountLikeByPostID, postID)

	err := row.Scan(&count)
	if err != nil {
		log.Error().Err(err).Msg("failed to count like by post id")
		return count, err
	}

	return count, nil
}
