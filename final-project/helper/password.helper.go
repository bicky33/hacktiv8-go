package helper

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPassword), err
}

func PasswordMatch(hashPassword string, currentPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(currentPassword))
	return err == nil
}
