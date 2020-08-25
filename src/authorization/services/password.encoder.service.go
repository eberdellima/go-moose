package services

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"fmt"
	"strings"

	"golang.org/x/crypto/argon2"
)

// PasswordConfig structure of parameters used
// for encoding and comparing passwords
type PasswordConfig struct {
	time    uint32
	memory  uint32
	threads uint8
	keyLen  uint32
}

var passwordConfig = PasswordConfig{
	time:    1,
	memory:  64 * 1024,
	threads: 4,
	keyLen:  32,
}

// GeneratePassword generates hashed password
func GeneratePassword(password string) (string, error) {

	// Generate a salt
	salt := make([]byte, 16)
	if _, err := rand.Read(salt); err != nil {
		return "", err
	}

	// Generate hash
	hash := argon2.IDKey([]byte(password), salt, passwordConfig.time, passwordConfig.memory, passwordConfig.threads, passwordConfig.keyLen)

	// Base64 encode the salt and hashed password.
	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)

	format := "$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s"
	full := fmt.Sprintf(format, argon2.Version, passwordConfig.memory, passwordConfig.time, passwordConfig.threads, b64Salt, b64Hash)
	return full, nil
}

// ComparePasswords checks if the provided password
// is the same as the hashed password
func ComparePasswords(password, hash string) (bool, error) {

	parts := strings.Split(hash, "$")

	// Get password configuration
	c := &PasswordConfig{}
	_, err := fmt.Sscanf(parts[3], "m=%d,t=%d,p=%d", &c.memory, &c.time, &c.threads)
	if err != nil {
		return false, err
	}

	// Decode salt
	salt, err := base64.RawStdEncoding.DecodeString(parts[4])
	if err != nil {
		return false, err
	}

	// Decode hash
	decodedHash, err := base64.RawStdEncoding.DecodeString(parts[5])
	if err != nil {
		return false, err
	}
	c.keyLen = uint32(len(decodedHash))

	// Generate hashed password
	comparisonHash := argon2.IDKey([]byte(password), salt, c.time, c.memory, c.threads, c.keyLen)

	return (subtle.ConstantTimeCompare(decodedHash, comparisonHash) == 1), nil
}
