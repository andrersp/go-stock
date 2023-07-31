package security

import "golang.org/x/crypto/bcrypt"

func HashGenerator(passwordPlain string) (passwordHashe string, err error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(passwordPlain), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(bytes), err
}

func CheckPasswordHash(passwordHash, passwordPlain string) error {

	err := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(passwordPlain))
	return err
}
