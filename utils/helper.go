package utils

import (
	"io"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
	"rental-app/models"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

func LoadEnv() {
	// Load environment variables from .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}
}

func UploadFile(file *multipart.FileHeader, uploadDir string) (string, error) {
	// Buat direktori jika belum ada
	if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
		return "", err
	}

	// Tentukan path file
	filePath := filepath.Join(uploadDir, file.Filename)

	// Buka file sumber
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	// Buka atau buat file tujuan
	dst, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	defer dst.Close()

	// Salin isi file dari sumber ke tujuan
	if _, err := io.Copy(dst, src); err != nil {
		return "", err
	}

	return filePath, nil
}

// GenerateToken generates a JWT token
func GenerateToken(user models.User) (string, error) {
	// Create a new JWT token
	claim := jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"role":    user.Role,
		"name":    user.Name,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
		"iat":     time.Now().Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	// Sign the token with a secret key

	// Sign and get the complete encoded token as a string using the secret
	LoadEnv()
	jwtKey := os.Getenv("JWT_KEY")
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}


// hash password
func HashPassword(password string) (string, error) {
	// Tentukan cost untuk bcrypt (default: bcrypt.DefaultCost)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// ComparePassword compares a hashed password with a plaintext password
func ComparePassword(hashedPassword, password string) error {
	// Cocokkan password yang di-hash dengan password plaintext
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
