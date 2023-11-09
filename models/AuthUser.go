package models

type AuthUser struct {
	Username string
	Email    string
	Role     string
}

func NewAuthUser(username, email, role string) *AuthUser {
	return &AuthUser{
		Username: username,
		Email:    email,
		Role:     role,
	}
}
