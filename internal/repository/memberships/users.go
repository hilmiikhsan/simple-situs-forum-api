package memberships

import (
	"context"
	"database/sql"

	"github.com/hilmiikhsan/situs-forum/internal/model/memberships"
	"github.com/sirupsen/logrus"
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
			logrus.Error("user not found")
			return nil, nil
		}

		logrus.Error("failed to get user by email or username: ", err)
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
		logrus.Error("failed to create user: ", err)
		return err
	}

	return nil
}
