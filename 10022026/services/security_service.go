package services

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword Hashing af password
func HashPassword(password []byte) (string, error) {

	// Laver zeroing på password input for at sikre at det ikke ligger i hukommelsen efter brug
	defer func() {
		for i := range password {
			password[i] = 0
		}
	}()

	hashBytes, err := bcrypt.GenerateFromPassword(password, 14)
	// Sikre at error ikke returnere vigtig data tilbage
	if err != nil {
		return "", err
	}
	return string(hashBytes), err
}

// CheckPasswordHash Check password hash
func CheckPasswordHash(password, hash []byte) bool {
	// Bruger bcrypt til at validere password hash
	err := bcrypt.CompareHashAndPassword(hash, password)

	// Laver zeroing på password input for at sikre at det ikke ligger i hukommelsen efter brug
	for i := range password {
		password[i] = 0
	}

	return err == nil
}

var cipherKey = []byte("strong-ass-key-for-encryption!!!") // 32 Byte

func Encrypt(text []byte) (string, error) {
	// // Laver zeroing på password input for at sikre at det ikke ligger i hukommelsen efter brug
	defer func() {
		for i := range text {
			text[i] = 0
		}
	}()

	// Opretter cipher block ud fra 32-byte cipher key - automatisk AES-256
	block, err := aes.NewCipher(cipherKey)
	if err != nil {
		return "", err
	}

	// Cipher block bliver pakket ind i GCM for at beskytte mod manipulation
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	// Nonce sikre mod at mønstre af kryptering af det samme data
	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	// Vi gemmer nonce forrest i pakken, så vi kan finde den ved dekryptering
	ciphertext := gcm.Seal(nonce, nonce, []byte(text), nil)

	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

func Decrypt(cryptoText string) ([]byte, error) {
	// Tekststreng -> binære bytes
	data, err := base64.StdEncoding.DecodeString(cryptoText)
	if err != nil {
		return nil, err
	}

	defer func() {
		for i := range data {
			data[i] = 0
		}
	}()

	// Vi opretter GCM-objekt med samme algoritme og nøgle
	block, err := aes.NewCipher(cipherKey)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	// Vi opretter nonce med 12 bytes
	nonceSize := gcm.NonceSize()
	if len(data) < nonceSize {
		return nil, fmt.Errorf("ciphertext for kort")
	}

	// Vi deler op i nonce og ciphertext
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	// vi dekryptere
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err // Returnerer fejl hvis data er manipuleret!
	}

	return plaintext, nil
}
