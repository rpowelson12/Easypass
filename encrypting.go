package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func MakeKey() string {
	bytes := make([]byte, 32)
	if _, err := rand.Read(bytes); err != nil {
		log.Fatalf("could not make key: %v", err)
	}

	key := hex.EncodeToString(bytes)

	return key
}
func saveKeyToEnvFile(path, hexKey string) error {
	content := fmt.Sprintf("ENCRYPTION_KEY=%s\n", hexKey)
	return os.WriteFile(path, []byte(content), 0600)

}

func ensureKeyFromEnv(path string) string {
	// Try to load .env
	_ = godotenv.Load(path)

	key := os.Getenv("ENCRYPTION_KEY")

	// If missing, generate and save it
	if key == "" {
		key = MakeKey()
		err := saveKeyToEnvFile(path, key)
		if err != nil {
			log.Fatalf("failed to save new .env file: %v", err)
		}

		// Reload the new .env
		if err := godotenv.Load(path); err != nil {
			log.Fatalf("failed to reload .env after saving: %v", err)
		}
	}

	return key
}

func Encode(stringToEncrypt string, keyString string) (encryptedString string) {

	key, _ := hex.DecodeString(keyString)
	plainText := []byte(stringToEncrypt)

	block, err := aes.NewCipher(key)
	if err != nil {
		log.Fatalf("could not make cipher: %v", err)
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		log.Fatalf("could not make GCM: %v", err)
	}

	nonce := make([]byte, aesGCM.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		log.Fatalf("could not make nonce: %v", err)
	}

	ciphertext := aesGCM.Seal(nonce, nonce, plainText, nil)

	return base64.StdEncoding.EncodeToString(ciphertext)
}

func Decrypt(encryptedString string, keyString string) (decryptedString string) {
	key, err := hex.DecodeString(keyString)
	if err != nil {
		log.Fatalf("could not decode key: %v", err)
	}

	enc, err := base64.StdEncoding.DecodeString(encryptedString)
	if err != nil {
		log.Fatalf("could not decode base64: %v", err)
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		log.Fatalf("could not create cipher: %v", err)
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		log.Fatalf("could not create GCM: %v", err)
	}

	nonceSize := aesGCM.NonceSize()
	if len(enc) < nonceSize {
		log.Fatalf("ciphertext too short")
	}

	nonce, ciphertext := enc[:nonceSize], enc[nonceSize:]

	plainText, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		log.Fatalf("decryption failed: %v", err)
	}

	return string(plainText)
}
