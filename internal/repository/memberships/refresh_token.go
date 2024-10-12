package memberships

import (
	"context"

	"github.com/hilmiikhsan/situs-forum/internal/model/memberships"
	"github.com/rs/zerolog/log"
)

func (r *repository) InsertRefreshToken(ctx context.Context, model memberships.RefreshTokenModel) error {
	_, err := r.db.ExecContext(ctx, queryInsertRefreshTOken,
		model.UserID,
		model.RefreshToken,
		model.ExpiredAt,
		model.CreatedAt,
		model.CreatedBy,
		model.UpdatedAt,
		model.UpdatedBy,
	)
	if err != nil {
		log.Error().Err(err).Msg("Failed to insert refresh token")
		return err
	}

	return nil
}

func (r *repository) GetRefreshToken(ctx context.Context, userID int64) (*memberships.RefreshTokenModel, error) {
	var response memberships.RefreshTokenModel

	row := r.db.QueryRowContext(ctx, queryGetRefreshToken, userID)

	err := row.Scan(
		&response.ID,
		&response.UserID,
		&response.RefreshToken,
		&response.ExpiredAt,
		&response.CreatedAt,
		&response.CreatedBy,
		&response.UpdatedAt,
		&response.UpdatedBy,
	)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get refresh token")
		return nil, err
	}

	return &response, nil
}
