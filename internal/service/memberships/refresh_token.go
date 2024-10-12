package memberships

import (
	"context"
	"errors"

	"github.com/hilmiikhsan/situs-forum/internal/model/memberships"
	"github.com/hilmiikhsan/situs-forum/pkg/jwt"
	"github.com/rs/zerolog/log"
)

func (s *service) ValidateRefreshToken(ctx context.Context, userID int64, req memberships.RefreshTokenRequest) (string, error) {
	existingRefreshToken, err := s.membershipRepo.GetRefreshToken(ctx, userID)
	if err != nil {
		log.Error().Err(err).Msg("failed to get refresh token")
		return "", err
	}

	if existingRefreshToken == nil {
		return "", errors.New("refresh token has been expired")
	}

	if existingRefreshToken.RefreshToken != req.Token {
		return "", errors.New("refresh token is invalid")
	}

	user, err := s.membershipRepo.GetUser(ctx, "", "", userID)
	if err != nil {
		log.Error().Err(err).Msg("failed to get user by id")
		return "", err
	}

	token, err := jwt.CreateToken(user.ID, user.Username, s.cfg.Service.SecretJWT)
	if err != nil {
		log.Error().Err(err).Msg("failed to create token")
		return "", err
	}

	return token, nil
}
