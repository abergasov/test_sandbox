package sampler

import (
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
