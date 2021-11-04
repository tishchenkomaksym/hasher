//Package hasher implements hashing string password and password vc comparison
package hasher

import "golang.org/x/crypto/bcrypt"

// HashPassword Hashes string password using bcrypt, returns hash, err
func HashPassword(password string) (string, error)  {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(hash), err
}

// CheckPasswordHash verify hashed string  using bcrypt, returns err (true, false check)
func CheckPasswordHash(password, hash string) bool  {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}