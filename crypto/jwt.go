package crypto
import(
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"os"
	"time"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

func LoadPrivateKey(path string) (*rsa.PrivateKey, error) {
	keyData, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(keyData)
	if block == nil || block.Type != "RSA PRIVATE KEY" {
		return nil, errors.New("failed to decode PEM block containing private key")
	}

	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return privateKey, nil
}

func GenerateJWT(privateKey *rsa.PrivateKey,username string) (string, error) {
	// Create claims
	claims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(24* time.Hour).Unix(),
		"iat":      time.Now().Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	// Sign using RSA private key
	return token.SignedString(privateKey)
}
func Validate_jwt(publicKey *rsa.PublicKey, tokenString string) {
	// Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validate the signing method
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return publicKey, nil
	})

	if err != nil {
		fmt.Println("Error validating token:", err)
		return
	}

	// Extract and print claims
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println("Token is valid.")
		fmt.Println("Username:", claims["username"])
		fmt.Println("Issued At:", time.Unix(int64(claims["iat"].(float64)), 0))
		fmt.Println("Expires At:", time.Unix(int64(claims["exp"].(float64)), 0))
	} else {
		fmt.Println("Invalid token")
	}
}
