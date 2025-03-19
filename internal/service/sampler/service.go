package sampler

import (
	"context"
	"sandbox/internal/entities"
	"sandbox/internal/logger"
	"sandbox/internal/repository"
	"strings"
)

type Service struct {
	log        logger.AppLogger
	repo       *repository.Repo
	remoteHost string
}

func InitService(log logger.AppLogger, repo *repository.Repo, remoteHost string) *Service {
	return &Service{
		remoteHost: remoteHost,
		repo:       repo,
		log:        log.With(logger.WithService("sampler")),
	}
}

func (s *Service) GetAllMessages(ctx context.Context) ([]*entities.ChatMessage, error) {
	msgList, err := s.repo.GetAllChatMessages(ctx)
	if err != nil {
		return nil, err
	}
	for i := range msgList {
		msgList[i].Message = strings.Repeat("*", len(msgList[i].Message))
	}
	return msgList, nil
}

func (s *Service) GetMessageByID(ctx context.Context, messageID uint64) (*entities.ChatMessage, error) {
	return s.repo.GetChatMessageByID(ctx, messageID)
}

func (s *Service) DeleteMessageID(ctx context.Context, messageID uint64) error {
	return s.repo.DeleteChatMessageByID(ctx, messageID)
}
