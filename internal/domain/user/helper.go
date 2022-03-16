package user

import "golang.org/x/crypto/bcrypt"

func generatePassHash(pass string) (string, error) {
	saltedBytes := []byte(pass)
	hashedBytes, err := bcrypt.GenerateFromPassword(saltedBytes, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	hash := string(hashedBytes[:])
	return hash, nil
}
