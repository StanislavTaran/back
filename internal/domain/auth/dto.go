package auth

// TODO Rename dto to input / output
type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
