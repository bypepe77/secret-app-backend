package secret

import (
	"errors"
	"fmt"

	"github.com/bypepe77/secret-app-backend/internal/models"
	"gorm.io/gorm"
)

type secretRepository struct {
	DB *gorm.DB
}

func NewSecretRepository(db *gorm.DB) SecretRepositoryInterface {
	return &secretRepository{
		DB: db,
	}
}

func (r *secretRepository) Create(p *SecretPayload) (*models.Confession, error) {
	var confession models.Confession
	confession.Content = p.Content
	confession.UserID = p.userID
	categories := []*models.Category{}
	//TODO: improve this
	for _, name := range p.Categories {
		var category *models.Category
		r.DB.FirstOrCreate(&category, Category{Name: name.Name})
		categories = append(categories, category)
	}
	confession.Categories = categories
	if err := r.DB.Create(&confession).Error; err != nil {
		return nil, err
	}
	return &confession, nil
}

func (repository *secretRepository) GetByIDWithUser(id int) (*models.Confession, error) {
	var confession models.Confession
	if err := repository.DB.Preload("User").Preload("Categories").Preload("LikesList").First(&confession, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("secret not found")
		}
		return nil, err
	}
	return &confession, nil
}

func (repository *secretRepository) GetMySecrets(userID int, pagination *Pagination) ([]*models.Confession, error) {
	var confessions []*models.Confession
	if err := repository.DB.Preload("User").
		Preload("Categories").
		Where("user_id = ?", userID).
		Limit(pagination.Limit).
		Offset(pagination.Offset).
		Find(&confessions).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("secrets not found")
		}
		return nil, err
	}
	return confessions, nil
}

func (repository *secretRepository) HasLiked(confessionID int, userID int) (bool, error) {
	var like *models.Like
	err := repository.DB.Where("confession_id = ? AND user_id = ?", confessionID, userID).First(&like).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			fmt.Println("like 2", err)
			return false, nil
		}
		fmt.Println("like 3", err)
		return false, err
	}
	return true, nil
}

func (repository *secretRepository) AddLikeToConfession(confessionID int, userID int) error {
	// Find the confession
	var confession models.Confession
	if err := repository.DB.First(&confession, confessionID).Error; err != nil {
		return err
	}
	// Create a new like
	newLike := &models.Like{
		ConfessionID: confessionID,
		UserID:       userID,
	}
	// Append the new like to the confession
	confession.LikesList = append(confession.LikesList, newLike)
	// Increase the likes count
	confession.LikesCount = confession.LikesCount + 1
	// Save the confession
	repository.DB.Save(&confession)
	return nil
}

func (repository *secretRepository) DeleteLikeFromConfession(confessionID int, userID int) error {
	// Delete the like from the association table
	result := repository.DB.Where("confession_id = ? AND user_id = ?", confessionID, userID).Delete(&models.Like{})
	if result.Error != nil {
		return result.Error
	}
	// Update the like count on the confession
	result = repository.DB.Model(&models.Confession{}).Where("id = ?", confessionID).Update("likes_count", gorm.Expr("likes_count - ?", 1))
	if result.Error != nil {
		return result.Error
	}
	return nil
}
