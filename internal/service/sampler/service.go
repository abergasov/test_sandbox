package sampler

import (
	"context"
	"sandbox/internal/entities"
	"sandbox/internal/logger"
	"sandbox/internal/repository"
)

type Service struct {
	log  logger.AppLogger
	repo *repository.Repo
}

func InitService(log logger.AppLogger, repo *repository.Repo) *Service {
	return &Service{
		repo: repo,
		log:  log.With(logger.WithService("sampler")),
	}
}

func (s *Service) Init(ctx context.Context) error {
	s.log.Info("initializing service")
	return nil
}

func (s *Service) GetAllMessages(ctx context.Context) ([]*entities.ChatMessage, error) {
	return s.repo.GetAllChatMessages(ctx)
}
