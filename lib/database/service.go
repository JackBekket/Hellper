package database

import (
	"errors"
	//"hellper/internal/database"
	"sync"

	"github.com/tmc/langchaingo/llms/openai"
)

var ErrHandlerNotFound = errors.New("handler for that user id is not found")
var ErrHandlerCast = errors.New("failed to cast LLM handler")

type Service struct {
	LLMHandlers sync.Map
	DBHandler   *Handler
}

func NewAIService(dbHandler *Handler) (*Service, error) {
	service := Service{
		DBHandler: dbHandler,
	}
	err := service.CreateTables()
	return &service, err
}

func (s *Service) GetHandler(userId int64) (*openai.LLM, error) {
	handlerAny, ok := s.LLMHandlers.Load(userId)
	if !ok {
		return nil, ErrHandlerNotFound
	}
	handler, ok := handlerAny.(*openai.LLM)
	if !ok {
		return nil, ErrHandlerCast
	}
	return handler, nil
}

func (s *Service) DropHandler(userId int64) {
	s.LLMHandlers.Delete(userId)
}

func (s *Service) UpdateHandler(userId int64, localAIToken, model, endpoint string) (*openai.LLM, error) {
	llm, err := openai.New(
		openai.WithToken(localAIToken),
		openai.WithModel(model),
		openai.WithBaseURL(endpoint),
	)
	if err != nil {
		return nil, err
	}
	s.LLMHandlers.Store(userId, llm)
	return llm, nil
}
