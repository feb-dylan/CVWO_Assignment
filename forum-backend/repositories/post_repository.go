package repositories

import (
	"forum/models"
	"gorm.io/gorm"
)

type PostRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) *PostRepository {
	return &PostRepository{db: db}
}

func (r *PostRepository) Create(post *models.Post) error {
	return r.db.Create(post).Error
}

func (r *PostRepository) GetAll() ([]models.Post, error) {
	var posts []models.Post
	err := r.db.Preload("User").Preload("Topic").Order("created_at DESC").Find(&posts).Error
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (r *PostRepository) GetByID(id uint) (*models.Post, error) {
	var post models.Post
	err := r.db.Preload("User").Preload("Topic").First(&post, id).Error
	if err != nil {
		return nil, err
	}
	return &post, nil
}

func (r *PostRepository) GetByTopicID(topicID uint) ([]models.Post, error) {
	var posts []models.Post
	err := r.db.Where("topic_id = ?", topicID).Preload("User").Order("created_at DESC").Find(&posts).Error
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (r *PostRepository) Update(post *models.Post) error {
	return r.db.Save(post).Error
}

func (r *PostRepository) Delete(id uint) error {
	return r.db.Delete(&models.Post{}, id).Error
}

func (r *PostRepository) Exists(id uint) (bool, error) {
	var count int64
	err := r.db.Model(&models.Post{}).Where("id = ?", id).Count(&count).Error
	return count > 0, err
}