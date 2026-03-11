package services

import (
	"errors"
	"forum/models"
	"forum/repositories"
	"forum/dto"
	"time"
)

type PostService struct {
	postRepo  *repositories.PostRepository
	topicRepo *repositories.TopicRepository
	userRepo  *repositories.UserRepository
}

func NewPostService(
	postRepo *repositories.PostRepository,
	topicRepo *repositories.TopicRepository,
	userRepo *repositories.UserRepository,
) *PostService {
	return &PostService{
		postRepo:  postRepo,
		topicRepo: topicRepo,
		userRepo:  userRepo,
	}
}

func (s *PostService) CreatePost(input dto.CreatePostDTO, userID uint) (*dto.PostResponseDTO, error) {
	user, err := s.userRepo.GetByID(userID)
	if err != nil {
		return nil, errors.New("user not found")
	}

	topic, err := s.topicRepo.GetByID(input.TopicID)
	if err != nil {
		return nil, errors.New("topic not found")
	}

	if len(input.Title) == 0 {
		return nil, errors.New("title cannot be empty")
	}

	post := models.Post{
		Title:     input.Title,
		Content:   input.Content,
		TopicID:   topic.ID,
		UserID:    user.ID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := s.postRepo.Create(&post); err != nil {
		return nil, errors.New("failed to create post")
	}

	return &dto.PostResponseDTO{
		ID:        post.ID,
		Title:     post.Title,
		Content:   post.Content,
		TopicID:   post.TopicID,
		UserID:    post.UserID,
		Username:  user.Username,
		CreatedAt: post.CreatedAt.Format(time.RFC3339),
		UpdatedAt: post.UpdatedAt.Format(time.RFC3339),
	}, nil
}

func (s *PostService) GetAllPosts() ([]dto.PostResponseDTO, error) {
	posts, err := s.postRepo.GetAll()
	if err != nil {
		return nil, errors.New("failed to retrieve posts")
	}

	var response []dto.PostResponseDTO
	for _, post := range posts {
		response = append(response, dto.PostResponseDTO{
			ID:        post.ID,
			Title:     post.Title,
			Content:   post.Content,
			TopicID:   post.TopicID,
			UserID:    post.UserID,
			Username:  post.User.Username,
			CreatedAt: post.CreatedAt.Format(time.RFC3339),
			UpdatedAt: post.UpdatedAt.Format(time.RFC3339),
		})
	}
	return response, nil
}

func (s *PostService) GetPostByID(id uint) (*dto.PostResponseDTO, error) {
	post, err := s.postRepo.GetByID(id)
	if err != nil {
		return nil, errors.New("post not found")
	}

	return &dto.PostResponseDTO{
		ID:        post.ID,
		Title:     post.Title,
		Content:   post.Content,
		TopicID:   post.TopicID,
		UserID:    post.UserID,
		Username:  post.User.Username,
		CreatedAt: post.CreatedAt.Format(time.RFC3339),
		UpdatedAt: post.UpdatedAt.Format(time.RFC3339),
	}, nil
}

func (s *PostService) GetPostsByTopic(topicID uint) ([]dto.PostResponseDTO, error) {
	_, err := s.topicRepo.GetByID(topicID)
	if err != nil {
		return nil, errors.New("topic not found")
	}

	posts, err := s.postRepo.GetByTopicID(topicID)
	if err != nil {
		return nil, errors.New("failed to retrieve posts")
	}

	var response []dto.PostResponseDTO
	for _, post := range posts {
		response = append(response, dto.PostResponseDTO{
			ID:        post.ID,
			Title:     post.Title,
			Content:   post.Content,
			TopicID:   post.TopicID,
			UserID:    post.UserID,
			Username:  post.User.Username,
			CreatedAt: post.CreatedAt.Format(time.RFC3339),
			UpdatedAt: post.UpdatedAt.Format(time.RFC3339),
		})
	}
	return response, nil
}

func (s *PostService) UpdatePost(id uint, input dto.UpdatePostDTO, userID uint) (*dto.PostResponseDTO, error) {
	post, err := s.postRepo.GetByID(id)
	if err != nil {
		return nil, errors.New("post not found")
	}

	if post.UserID != userID {
		return nil, errors.New("unauthorized")
	}

	if input.Title != nil {
		if len(*input.Title) == 0 {
			return nil, errors.New("title cannot be empty")
		}
		post.Title = *input.Title
	}
	if input.Content != nil {
		if len(*input.Content) == 0 {
			return nil, errors.New("content cannot be empty")
		}
		post.Content = *input.Content
	}
	post.UpdatedAt = time.Now()

	if err := s.postRepo.Update(post); err != nil {
		return nil, errors.New("failed to update post")
	}

	updatedPost, err := s.postRepo.GetByID(id)
	if err != nil {
		return nil, errors.New("post not found after update")
	}

	return &dto.PostResponseDTO{
		ID:        updatedPost.ID,
		Title:     updatedPost.Title,
		Content:   updatedPost.Content,
		TopicID:   updatedPost.TopicID,
		UserID:    updatedPost.UserID,
		Username:  updatedPost.User.Username,
		CreatedAt: updatedPost.CreatedAt.Format(time.RFC3339),
		UpdatedAt: updatedPost.UpdatedAt.Format(time.RFC3339),
	}, nil
}

func (s *PostService) DeletePost(id uint, userID uint) error {
	post, err := s.postRepo.GetByID(id)
	if err != nil {
		return errors.New("post not found")
	}

	if post.UserID != userID {
		return errors.New("unauthorized")
	}

	if err := s.postRepo.Delete(id); err != nil {
		return errors.New("failed to delete post")
	}

	return nil
}