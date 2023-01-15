package secret

import (
	"errors"

	"github.com/bypepe77/secret-app-backend/internal/models"
	"gorm.io/gorm"
)

type SecretRepositoryInterface interface {
	Create(p *SecretPayload) (*models.Confession, error)
	GetByIDWithUser(id int) (*models.Confession, error)
}

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
	if err := repository.DB.Preload("User").Preload("Categories").First(&confession, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("secret not found")
		}
		return nil, err
	}
	return &confession, nil
}
