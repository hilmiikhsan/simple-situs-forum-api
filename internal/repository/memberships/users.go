package memberships

import (
	"context"
	"database/sql"

	"github.com/hilmiikhsan/situs-forum/internal/model/memberships"
	"github.com/rs/zerolog/log"
)

func (r *repository) GetUseByEmailOrUsername(ctx context.Context, email, username string) (*memberships.UserModel, error) {
	row := r.db.QueryRowContext(ctx, queryGetUserByEmailOrUsername,
		email,
		username,
	)

	var response memberships.UserModel

	err := row.Scan(
		&response.ID,
		&response.Email,
		&response.Username,
		&response.Password,
		&response.CreatedAt,
		&response.CreatedBy,
		&response.UpdatedAt,
		&response.UpdatedBy,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Error().Err(err).Msg("user not found")
			return nil, nil
		}

		log.Error().Err(err).Msg("failed to get user by email or username")
		return nil, err
	}

	return &response, nil
}

func (r *repository) CreateUser(ctx context.Context, model memberships.UserModel) error {
	_, err := r.db.ExecContext(ctx, queryInsertUser,
		model.Email,
		model.Username,
		model.Password,
		model.CreatedAt,
		model.CreatedBy,
		model.UpdatedAt,
		model.UpdatedBy,
	)
	if err != nil {
		log.Error().Err(err).Msg("failed to insert user")
		return err
	}

	return nil
}
