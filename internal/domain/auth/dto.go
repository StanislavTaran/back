package auth

type CredentialsInputDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
