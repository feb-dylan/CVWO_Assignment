package repositories

import (
	"forum/models"
	"gorm.io/gorm"
)

type TopicRepository struct {
	db *gorm.DB
}

func NewTopicRepository(db *gorm.DB) *TopicRepository {
	return &TopicRepository{db: db}
}

func (r *TopicRepository) Create(topic *models.Topic) error {
	return r.db.Create(topic).Error
}

func (r *TopicRepository) GetAll() ([]models.Topic, error) {
	var topics []models.Topic
	err := r.db.Preload("User").Order("created_at DESC").Find(&topics).Error
	if err != nil {
		return nil, err
	}
	return topics, nil
}

func (r *TopicRepository) GetByID(id uint) (*models.Topic, error) {
	var topic models.Topic
	err := r.db.Preload("User").First(&topic, id).Error
	if err != nil {
		return nil, err
	}
	return &topic, nil
}

func (r *TopicRepository) Update(topic *models.Topic) error {
	return r.db.Save(topic).Error
}

func (r *TopicRepository) Delete(id uint) error {
	return r.db.Delete(&models.Topic{}, id).Error
}
func (r *TopicRepository) Exists(id uint) (bool, error) {
	var count int64
	err := r.db.Model(&models.Topic{}).Where("id = ?", id).Count(&count).Error
	return count > 0, err
}
