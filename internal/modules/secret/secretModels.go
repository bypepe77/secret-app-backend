package secret

import "github.com/bypepe77/secret-app-backend/internal/models"

type SecretPayload struct {
	Content    string             `json:"content"`
	Categories []*models.Category `json:"category"`
	userID     int
}

type Secret struct {
	Content     string   `json:"content"`
	IsDestroyed bool     `json:"is_destroyed"`
	Categories  []string `json:"categories"`
	User        *User    `json:"user"`
}

type User struct {
	ID       int    `json:"user_id"`
	Username string `json:"username"`
}

type Category struct {
	Name string `json:"name"`
}

type Pagination struct {
	Limit  int
	Offset int
}
