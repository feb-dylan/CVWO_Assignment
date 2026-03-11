package services

import (
	"errors"
	"forum/models"
	"forum/repositories"
	"forum/dto"
	"time"
)

type CommentService struct {
	commentRepo *repositories.CommentRepository
	postRepo    *repositories.PostRepository
	userRepo    *repositories.UserRepository
}

func NewCommentService(
	commentRepo *repositories.CommentRepository,
	postRepo *repositories.PostRepository,
	userRepo *repositories.UserRepository,
) *CommentService {
	return &CommentService{
		commentRepo: commentRepo,
		postRepo:    postRepo,
		userRepo:    userRepo,
	}
}

func (s *CommentService) CreateComment(input dto.CreateCommentDTO, userID uint) (*dto.CommentResponseDTO, error) {
	user, err := s.userRepo.GetByID(userID)
	if err != nil {
		return nil, errors.New("user not found")
	}

	post, err := s.postRepo.GetByID(input.PostID)
	if err != nil {
		return nil, errors.New("post not found")
	}

	if input.ParentID != nil {
		parentExists, err := s.commentRepo.Exists(*input.ParentID)
		if err != nil || !parentExists {
			return nil, errors.New("parent comment not found")
		}
		
		parent, _ := s.commentRepo.GetByID(*input.ParentID)
		if parent.PostID != input.PostID {
			return nil, errors.New("parent comment does not belong to this post")
		}
	}

	comment := &models.Comment{
		Content:   input.Content,
		PostID:    post.ID,
		UserID:    user.ID,
		ParentID:  input.ParentID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := s.commentRepo.Create(comment); err != nil {
		return nil, errors.New("failed to create comment")
	}

	createdComment, err := s.commentRepo.GetByID(comment.ID)
	if err != nil {
		return nil, errors.New("comment created but failed to retrieve")
	}

	return toCommentResponseDTO(createdComment), nil
}

func (s *CommentService) GetCommentsByPostID(postID uint) ([]dto.CommentResponseDTO, error) {
	_, err := s.postRepo.GetByID(postID)
	if err != nil {
		return nil, errors.New("post not found")
	}

	comments, err := s.commentRepo.GetByPostID(postID)
	if err != nil {
		return nil, errors.New("failed to retrieve comments")
	}

	var response []dto.CommentResponseDTO
	for _, comment := range comments {
		response = append(response, *toCommentResponseDTO(&comment))
	}
	return response, nil
}

func (s *CommentService) GetCommentByID(id uint) (*dto.CommentResponseDTO, error) {
	comment, err := s.commentRepo.GetByID(id)
	if err != nil {
		return nil, errors.New("comment not found")
	}
	return toCommentResponseDTO(comment), nil
}

func (s *CommentService) UpdateComment(id uint, input dto.UpdateCommentDTO, userID uint) (*dto.CommentResponseDTO, error) {
	comment, err := s.commentRepo.GetByID(id)
	if err != nil {
		return nil, errors.New("comment not found")
	}

	if comment.UserID != userID {
		return nil, errors.New("unauthorized")
	}

	if input.Content != nil {
		if len(*input.Content) == 0 {
			return nil, errors.New("content cannot be empty")
		}
		comment.Content = *input.Content
	}
	comment.UpdatedAt = time.Now()

	if err := s.commentRepo.Update(comment); err != nil {
		return nil, errors.New("failed to update comment")
	}

	updatedComment, err := s.commentRepo.GetByID(id)
	if err != nil {
		return nil, errors.New("comment updated but failed to retrieve")
	}

	return toCommentResponseDTO(updatedComment), nil
}

func (s *CommentService) DeleteComment(id uint, userID uint) error {
	comment, err := s.commentRepo.GetByID(id)
	if err != nil {
		return errors.New("comment not found")
	}

	if comment.UserID != userID {
		return errors.New("unauthorized")
	}

	if err := s.commentRepo.Delete(id); err != nil {
		return errors.New("failed to delete comment")
	}

	return nil
}

func (s *CommentService) GetRepliesByCommentID(commentID uint) ([]dto.CommentResponseDTO, error) {
	exists, err := s.commentRepo.Exists(commentID)
	if err != nil || !exists {
		return nil, errors.New("comment not found")
	}

	replies, err := s.commentRepo.GetRepliesByCommentID(commentID)
	if err != nil {
		return nil, errors.New("failed to retrieve replies")
	}

	var response []dto.CommentResponseDTO
	for _, reply := range replies {
		response = append(response, *toCommentResponseDTO(&reply))
	}
	return response, nil
}
func toCommentResponseDTO(comment *models.Comment) *dto.CommentResponseDTO {
	response := &dto.CommentResponseDTO{
		ID:        comment.ID,
		Content:   comment.Content,
		PostID:    comment.PostID,
		UserID:    comment.UserID,
		Username:  comment.User.Username,
		ParentID:  comment.ParentID,
		CreatedAt: comment.CreatedAt.Format(time.RFC3339),
		UpdatedAt: comment.UpdatedAt.Format(time.RFC3339),
	}

	if len(comment.Replies) > 0 {
		response.Replies = make([]dto.CommentResponseDTO, len(comment.Replies))
		for i, reply := range comment.Replies {
			replyDTO := toCommentResponseDTO(&reply)
			response.Replies[i] = *replyDTO
		}
	}

	return response
}