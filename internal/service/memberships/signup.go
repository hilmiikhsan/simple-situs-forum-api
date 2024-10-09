package memberships

import (
	"context"
	"errors"
	"time"

	"github.com/hilmiikhsan/situs-forum/internal/model/memberships"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

func (s *service) SignUp(ctx context.Context, req memberships.SignUpRequest) error {
	user, err := s.membershipRepo.GetUseByEmailOrUsername(ctx, req.Email, req.Username)
	if err != nil {
		logrus.Error("failed to get user by email or username: ", err)
		return err
	}

	if user != nil {
		logrus.Error("username or email already exists")
		return errors.New("username or email already exists")
	}

	password, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		logrus.Error("failed to hash password: ", err)
		return err
	}

	now := time.Now()

	model := memberships.UserModel{
		Email:     req.Email,
		Username:  req.Username,
		Password:  string(password),
		CreatedAt: now,
		CreatedBy: req.Email,
		UpdatedBy: req.Email,
	}

	return s.membershipRepo.CreateUser(ctx, model)
}
