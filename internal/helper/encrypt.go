package helper

import "golang.org/x/crypto/bcrypt"

// function to encrypt a string
func Encrypt(s string) (string, error) {
	// generare hashed string using minimum cost (10)
	hashedString, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	return string(hashedString), nil
}

func CompareHashAndString(s1 string, s2 string) error {
	err := bcrypt.CompareHashAndPassword([]byte(s1), []byte(s2))

	return err
}
