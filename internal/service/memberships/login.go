package memberships

import (
	"context"
	"errors"

	"github.com/hilmiikhsan/situs-forum/internal/model/memberships"
	"github.com/hilmiikhsan/situs-forum/pkg/jwt"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

func (s *service) Login(ctx context.Context, req memberships.LoginRequest) (string, error) {
	user, err := s.membershipRepo.GetUseByEmailOrUsername(ctx, req.Email, "")
	if err != nil {
		log.Error().Err(err).Msg("failed to get user by email or username")
		return "", err
	}

	if user == nil {
		return "", errors.New("email not exist")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		log.Error().Err(err).Msg("failed to compare password")
		return "", err
	}

	token, err := jwt.CreateToken(user.ID, user.Username, s.cfg.Service.SecretJWT)
	if err != nil {
		log.Error().Err(err).Msg("failed to create token")
		return "", err
	}

	return token, nil
}
