package auth

type UserPayload struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserCretedResponse struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	IsBanned bool   `json:"is_banned"`
	IsActive bool   `json:"is_active"`
	Token    string `json:"token"`
}
