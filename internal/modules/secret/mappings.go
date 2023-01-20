package secret

import "github.com/bypepe77/secret-app-backend/internal/models"

func toSecretResponse(secret *models.Confession) *Secret {
	return &Secret{
		Content:     secret.Content,
		IsDestroyed: secret.IsDestroyed,
		Categories:  toCategory(secret.Categories),
		LikesCount:  secret.LikesCount,
	}
}

func toSecretResponseArray(secrets []*models.Confession) []*Secret {
	var secretsResponse []*Secret

	for _, secret := range secrets {
		secretStruct := &Secret{
			Content:     secret.Content,
			IsDestroyed: secret.IsDestroyed,
			Categories:  toCategory(secret.Categories),
			LikesCount:  secret.LikesCount,
		}
		secretsResponse = append(secretsResponse, secretStruct)
	}
	return secretsResponse
}

func toCategory(categories []*models.Category) []string {
	var categoriesParsed []string
	for _, category := range categories {
		categoriesParsed = append(categoriesParsed, category.Name)
	}
	return categoriesParsed
}
