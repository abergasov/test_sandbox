package sampler

import (
	"context"
	"fmt"
	"net/http"
	"sandbox/internal/entities"
	"sandbox/internal/logger"
	"sandbox/internal/utils"
)

func (s *Service) PrepareState(ctx context.Context) error {
	s.log.Info("preparing sampling state...")

	s.log.Info("erasing all tables...")
	if err := s.repo.CleanupTables(ctx); err != nil {
		s.log.Error("failed to erase tables", err)
		return err
	}
	s.log.Info("preparing sampling state...")
	s.log.Info("loading all observe chats...")
	items, err := s.loadAllObserveChats(ctx)
	if err != nil {
		s.log.Error("failed to load observe chats", err)
		return err
	}
	if err = s.repo.SaveObserveChats(ctx, items); err != nil {
		s.log.Error("failed to save observe chats", err)
		return err
	}
	s.log.Info("loading all observe chats...done")
	s.log.Info("loading sample messages...")
	for _, item := range items {
		messages, errL := s.loadAllObserveMessages(ctx, item.ChatID)
		if errL != nil {
			s.log.Error("failed to load observe messages", errL, logger.WithUnt64("chat_id", item.ChatID))
			continue
		}
		if err = s.repo.SaveChatMessages(ctx, messages); err != nil {
			s.log.Error("failed to save observe messages", err)
		}
	}
	return nil
}

func (s *Service) loadAllObserveChats(ctx context.Context) ([]entities.ObserveChat, error) {
	u := fmt.Sprintf("%s/api/get_observe_chats", s.remoteHost)
	res, code, err := utils.GetCurl[[]entities.ObserveChat](ctx, u, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get observe chats: %w", err)
	}
	if code != http.StatusOK {
		return nil, fmt.Errorf("failed to get observe chats: %d", code)
	}
	return *res, nil
}

func (s *Service) loadAllObserveMessages(ctx context.Context, chatID uint64) ([]entities.ChatMessage, error) {
	u := fmt.Sprintf("%s/api/get_tg_messages/%d", s.remoteHost, chatID)
	res, code, err := utils.GetCurl[[]entities.ChatMessage](ctx, u, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get observe chats: %w", err)
	}
	if code != http.StatusOK {
		return nil, fmt.Errorf("failed to get observe messages: %d", code)
	}
	return *res, nil
}
