package secret

import "github.com/bypepe77/secret-app-backend/internal/models"

func toSecretResponse(secret *models.Confession) *Secret {
	return &Secret{
		Content:     secret.Content,
		IsDestroyed: secret.IsDestroyed,
		Categories:  toCategory(secret.Categories),
		User: &User{
			ID:       secret.UserID,
			Username: secret.User.Username,
		},
	}
}

func toCategory(categories []*models.Category) []string {
	var categoriesParsed []string
	for _, category := range categories {
		categoriesParsed = append(categoriesParsed, category.Name)
	}
	return categoriesParsed
}
