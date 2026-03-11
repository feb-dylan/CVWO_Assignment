package services

import (
	"errors"
	"forum/models"
	"forum/repositories"
	"forum/dto"
	"time"
)
type TopicService struct {
	topicRepo *repositories.TopicRepository
	userRepo  *repositories.UserRepository
}

func NewTopicService(topicRepo *repositories.TopicRepository, userRepo *repositories.UserRepository) *TopicService {
	return &TopicService{
		topicRepo: topicRepo,
		userRepo:  userRepo,
	}
}

func (s *TopicService) CreateTopic(input dto.CreateTopicDTO , userID uint) (*dto.TopicResponseDTO, error) {
	user , err := s.userRepo.GetByID(userID)
	if err != nil {
		return nil, errors.New("user not found")
	}
	if (len(input.Title) == 0) {
		return nil, errors.New("title cannot be empty")
	}
	topic := models.Topic{
		Title: input.Title,
		Description: input.Description,
		UserID: user.ID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	if err := s.topicRepo.Create(&topic); err != nil {
		return nil, errors.New("failed to create topic")
	}
	return &dto.TopicResponseDTO{
		ID: topic.ID,
		Title: topic.Title,
		Description: topic.Description,
		UserID: topic.UserID,
		Username: topic.User.Username,
		CreatedAt: topic.CreatedAt.Format(time.RFC3339),
		UpdatedAt: topic.UpdatedAt.Format(time.RFC3339),
	}, nil
}

func (s *TopicService) GetAllTopics() ([]dto.TopicResponseDTO, error) {
	topics , err := s.topicRepo.GetAll()
	if err != nil {
		return nil, errors.New("failed to retrieve topics")
	}
	var response []dto.TopicResponseDTO
	for _, topic := range topics {
		response = append(response, dto.TopicResponseDTO{
			ID: topic.ID,
			Title: topic.Title,
			Description: topic.Description,
			UserID: topic.UserID,
			Username: topic.User.Username,
			CreatedAt: topic.CreatedAt.Format(time.RFC3339),
			UpdatedAt: topic.UpdatedAt.Format(time.RFC3339),
		})
	}
	return response, nil
}

func (s *TopicService) GetTopicByID(id uint) (*dto.TopicResponseDTO, error) {
	topic , err := s.topicRepo.GetByID(id)
	if err != nil {
		return nil, errors.New("topic not found")
	}
	return &dto.TopicResponseDTO{
		ID: topic.ID,
		Title: topic.Title,
		Description: topic.Description,
		UserID: topic.UserID,
		Username: topic.User.Username,
		CreatedAt: topic.CreatedAt.Format(time.RFC3339),
		UpdatedAt: topic.UpdatedAt.Format(time.RFC3339),
	}, nil
}

func (s *TopicService) UpdateTopic(id uint, input dto.UpdateTopicDTO , userID uint) (*dto.TopicResponseDTO, error) {
	topic , err := s.topicRepo.GetByID(id)
	if err != nil {
		return nil, errors.New("topic not found")
	}
	if topic.UserID != userID {
		return nil, errors.New("unauthorized")
	}
	if input.Title != nil {
		if len(*input.Title) == 0 {
			return nil, errors.New("title cannot be empty")
		}
		topic.Title = *input.Title
	}
	if input.Description != nil {
		topic.Description = *input.Description
	}
	topic.UpdatedAt = time.Now()

	if err := s.topicRepo.Update(topic); err != nil {
		return nil, errors.New("failed to update topic")
	}

	updatedTopic , err := s.topicRepo.GetByID(id)
	if err != nil {
		return nil, errors.New("topic not found after update")
	}

	return &dto.TopicResponseDTO{
		ID: updatedTopic.ID,
		Title: updatedTopic.Title,
		Description: updatedTopic.Description,
		UserID: updatedTopic.UserID,
		Username: updatedTopic.User.Username,
		CreatedAt: updatedTopic.CreatedAt.Format(time.RFC3339),
		UpdatedAt: updatedTopic.UpdatedAt.Format(time.RFC3339),
	}, nil
}

func (s *TopicService) DeleteTopic(id uint, userID uint) error {
	topic , err := s.topicRepo.GetByID(id)
	if err != nil {
		return errors.New("topic not found")
	}
	if topic.UserID != userID {
		return errors.New("unauthorized")
	}
	if err := s.topicRepo.Delete(id); err != nil {
		return errors.New("failed to delete topic")
	}
	return nil
}