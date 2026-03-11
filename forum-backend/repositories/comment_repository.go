package repositories

import (
	"forum/models"
	"gorm.io/gorm"
)

type CommentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) *CommentRepository {
	return &CommentRepository{db: db}
}

func (r *CommentRepository) Create(comment *models.Comment) error {
	return r.db.Create(comment).Error
}

func (r *CommentRepository) GetAll() ([]models.Comment, error) {
	var comments []models.Comment
	err := r.db.Preload("User").Order("created_at DESC").Find(&comments).Error
	return comments, err
}

func (r *CommentRepository) GetByID(id uint) (*models.Comment, error) {
	var comment models.Comment
	err := r.db.Preload("User").
		Preload("Replies", func(db *gorm.DB) *gorm.DB {
			return db.Preload("User").Order("created_at ASC")
		}).
		Preload("Replies.Replies").
		First(&comment, id).Error
	if err != nil {
		return nil, err
	}
	return &comment, nil
}

func (r *CommentRepository) GetByPostID(postID uint) ([]models.Comment, error) {
	var comments []models.Comment
	err := r.db.Where("post_id = ? AND parent_id IS NULL", postID).
		Preload("User").
		Preload("Replies", func(db *gorm.DB) *gorm.DB {
			return db.Preload("User").Order("created_at ASC")
		}).
		Preload("Replies.Replies").
		Preload("Replies.Replies.User").
		Order("created_at DESC").
		Find(&comments).Error
	return comments, err
}

func (r *CommentRepository) GetRepliesByCommentID(commentID uint) ([]models.Comment, error) {
	var replies []models.Comment
	err := r.db.Where("parent_id = ?", commentID).
		Preload("User").
		Order("created_at ASC").
		Find(&replies).Error
	return replies, err
}

func (r *CommentRepository) Update(comment *models.Comment) error {
	return r.db.Save(comment).Error
}

func (r *CommentRepository) Delete(id uint) error {
	return r.db.Delete(&models.Comment{}, id).Error
}

func (r *CommentRepository) Exists(id uint) (bool, error) {
	var count int64
	err := r.db.Model(&models.Comment{}).Where("id = ?", id).Count(&count).Error
	return count > 0, err
}

func (r *CommentRepository) GetCommentCountByPostID(postID uint) (int64, error) {
	var count int64
	err := r.db.Model(&models.Comment{}).Where("post_id = ?", postID).Count(&count).Error
	return count, err
}