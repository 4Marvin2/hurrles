package usecase

import (
	"context"
	"errors"
	"hurrles/config"
	"hurrles/internal/models"
	"hurrles/internal/user/repository"
	"time"
)

type ISessionUsecase interface {
	AddSession(context.Context, models.Session) error
	DeleteSession(context.Context) error
	GetSessionByCookie(context.Context, string) (models.Session, error)
}

type sessionUsecase struct {
	Session        repository.ISessionRepository
	contextTimeout time.Duration
}

func NewSessionUsecase(sess repository.ISessionRepository, timeout time.Duration) ISessionUsecase {
	return &sessionUsecase{
		Session:        sess,
		contextTimeout: timeout,
	}
}

func (s *sessionUsecase) AddSession(ctx context.Context, session models.Session) error {
	err := s.Session.NewSessionCookie(ctx, session.Cookie, session.UserID)
	if err != nil {
		return err
	}
	return nil
}

func (s *sessionUsecase) DeleteSession(ctx context.Context) error {
	ctxSession := ctx.Value(config.ContextUserID)
	if ctxSession == nil {
		return errors.New("context nil error")
	}
	currentSession, ok := ctxSession.(models.Session)
	if !ok {
		return errors.New("convert to model session error")
	}
	err := s.Session.DeleteSessionCookie(ctx, currentSession.Cookie)
	if err != nil {
		return err
	}
	return nil
}

func (s *sessionUsecase) GetSessionByCookie(ctx context.Context, sessionCookie string) (session models.Session, err error) {
	findSession, err := s.Session.GetSessionByCookie(ctx, sessionCookie)
	if err != nil {
		return models.Session{}, err
	}
	return findSession, nil
}
