package service

import (
	logger "bitbucket.org/gomatrix/customlogger"
	"github.com/google/uuid"
	"github.com/maksymchuk-mm/crm/internal/models"
	"github.com/maksymchuk-mm/crm/internal/repository"
	"github.com/maksymchuk-mm/crm/pkg/auth"
	"github.com/maksymchuk-mm/crm/pkg/errors"
	"time"
)

type UserService struct {
	tokenManager auth.TokenManager
	password     auth.PasswordGenerator

	repo            repository.User
	accessTokenTTL  time.Duration
	refreshTokenTTL time.Duration
}

func NewUserService(tokenManager auth.TokenManager, password auth.PasswordGenerator, repo repository.User,
	accessTokenTTL time.Duration, refreshTokenTTL time.Duration) *UserService {
	return &UserService{
		tokenManager:    tokenManager,
		password:        password,
		repo:            repo,
		accessTokenTTL:  accessTokenTTL,
		refreshTokenTTL: refreshTokenTTL,
	}
}

func (s *UserService) SignIn(input SignInInput) (Tokens, errors.CrmError) {
	user, err := s.repo.Get(input.login)
	if err != nil {
		logger.Errorf(err.Error())
		return Tokens{}, errors.Errors.DataBaseError
	}
	if err := s.password.CompareHashAndPassword(user.Password.Password, input.Password); err != nil {
		logger.Warnf(err.Error())
		return Tokens{}, errors.Errors.WrongSigInData
	}
	return s.createSession(user.PublicUserID)
}

func (s *UserService) createSession(userPublicID uuid.UUID) (res Tokens, err errors.CrmError) {
	res.AccessToken, err = s.tokenManager.NewJWT(userPublicID, s.accessTokenTTL)
	if err != nil {
		logger.Errorf(err.Error())
		return res, err
	}
	res.RefreshToken, err = s.tokenManager.NewRefreshToken()
	if err != nil {
		logger.Errorf(err.Error())
		return res, err
	}
	session := models.Session{
		RefreshToken: res.RefreshToken,
		ExpiredAt:    time.Now().Add(s.refreshTokenTTL),
	}

	err_ := s.repo.SetSession(session)
	if err_ != nil {
		logger.Errorf(err_.Error())
		err = errors.Errors.DataBaseError
		return res, err
	}
	return res, err
}
