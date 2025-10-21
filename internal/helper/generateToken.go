package helper

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/yogs696/skilltest/pkg/rsa256"
)

const jwtRSAChar = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var jwtRSASecretKeyLen uint

func GenerateToken(userID uint64, username, email string) (string, error) {
	// Read RSA private key
	privk, err := rsa256.ReadPrivateKey("private-key.pem")
	if err != nil {
		fmt.Printf("[generateToken]-failed to read RSA private key: %s", err.Error())
		return "reading RSA private key", fmt.Errorf("failed to read RSA private key: %s", err.Error())
	}

	// Generate and write the secret key
	secretKey, err := writeSecretKey()
	if err != nil {
		fmt.Printf("[generateToken] - failed to generate and write secret key:%s", err.Error())
		return secretKey, fmt.Errorf("failed to generate and write secret key: %s", err.Error())
	}

	// Signing RSA private key to JWT
	token := jwt.New(jwt.SigningMethodRS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = userID
	claims["email"] = email
	claims["username"] = username
	claims["sec"] = secretKey
	k, err := token.SignedString(privk)
	if err != nil {
		fmt.Printf("[generateToken]-%s", "failed to signed JWT with the RSA private key")
		return k, errors.New("failed to signed JWT with the RSA private key")
	}

	fmt.Printf("[generateToken]-%s", "New JWT token has been generated")
	return k, nil
}

// Helper function to generate and write the secret key
func writeSecretKey() (string, error) {
	// Create file
	f, err := os.Create("secret.key")
	if err != nil {
		return "", err
	}
	defer f.Close()

	// Make random char
	seedrand := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, jwtRSASecretKeyLen)
	for i := range b {
		b[i] = jwtRSAChar[seedrand.Intn(len(jwtRSAChar))]
	}

	// Convert random char to md5
	hasher := md5.New()
	hasher.Write([]byte(string(b)))
	secretKey := hex.EncodeToString(hasher.Sum(nil))

	// Write to file
	if _, err := f.WriteString(secretKey); err != nil {
		return "", err
	}

	return secretKey, nil
}

type authMetaContext struct {
	UserID   uint64 `json:"user_id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

// // Return structured value of parsed auth meta from context.
func CtxValue(ctx echo.Context) *authMetaContext {
	usr := ctx.Get("user").(*jwt.Token)
	claims := usr.Claims.(jwt.MapClaims)

	fmt.Println("claims")
	fmt.Println(claims["email"])

	return &authMetaContext{
		UserID:   uint64(claims["user_id"].(float64)),
		Email:    fmt.Sprintf("%s", claims["email"]),
		Username: fmt.Sprintf("%s", claims["username"]),
	}
}
