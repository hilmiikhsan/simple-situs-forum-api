package memberships

import (
	"context"
	"errors"
	"strconv"
	"time"

	"github.com/hilmiikhsan/situs-forum/internal/model/memberships"
	"github.com/hilmiikhsan/situs-forum/pkg/jwt"
	"github.com/hilmiikhsan/situs-forum/pkg/refresh_token"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

func (s *service) Login(ctx context.Context, req memberships.LoginRequest) (string, string, error) {
	user, err := s.membershipRepo.GetUser(ctx, req.Email, "", 0)
	if err != nil {
		log.Error().Err(err).Msg("failed to get user by email or username")
		return "", "", err
	}

	if user == nil {
		return "", "", errors.New("email not exist")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		log.Error().Err(err).Msg("failed to compare password")
		return "", "", err
	}

	token, err := jwt.CreateToken(user.ID, user.Username, s.cfg.Service.SecretJWT)
	if err != nil {
		log.Error().Err(err).Msg("failed to create token")
		return "", "", err
	}

	existingRefreshToken, err := s.membershipRepo.GetRefreshToken(ctx, user.ID)
	if err != nil {
		log.Error().Err(err).Msg("failed to get refresh token")
		return "", "", err
	}

	if existingRefreshToken != nil {
		return token, existingRefreshToken.RefreshToken, nil
	}

	refreshToken := refresh_token.GenerateRefreshToken()

	if refreshToken == "" {
		return token, "", errors.New("failed to generate refresh token")
	}

	err = s.membershipRepo.InsertRefreshToken(ctx, memberships.RefreshTokenModel{
		UserID:       user.ID,
		RefreshToken: refreshToken,
		ExpiredAt:    time.Now().Add(time.Hour * 24 * 10),
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
		CreatedBy:    strconv.FormatInt(user.ID, 10),
		UpdatedBy:    strconv.FormatInt(user.ID, 10),
	})
	if err != nil {
		log.Error().Err(err).Msg("failed to insert refresh token")
		return "", "", err
	}

	return token, refreshToken, nil
}
