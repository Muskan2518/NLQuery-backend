package main

import (
	"log"
	"NLQuery-backend/handlers"
	"NLQuery-backend/db"

	"github.com/gin-gonic/gin"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

func main() {
	// Generate a new private key
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		fmt.Println("Error generating RSA key:", err)
		return
	}

	// Encode and save private key to file
	privateKeyFile, err := os.Create("private_key.pem")
	if err != nil {
		fmt.Println("Error creating private key file:", err)
		return
	}
	defer privateKeyFile.Close()

	privateKeyPEM := pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	}
	if err := pem.Encode(privateKeyFile, &privateKeyPEM); err != nil {
		fmt.Println("Error encoding private key:", err)
		return
	}

	// Extract and encode the public key
	publicKey := &privateKey.PublicKey
	pubASN1, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		fmt.Println("Error marshaling public key:", err)
		return
	}

	// Save public key to file
	publicKeyFile, err := os.Create("public_key.pem")
	if err != nil {
		fmt.Println("Error creating public key file:", err)
		return
	}
	defer publicKeyFile.Close()

	publicKeyPEM := pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: pubASN1,
	}
	if err := pem.Encode(publicKeyFile, &publicKeyPEM); err != nil {
		fmt.Println("Error encoding public key:", err)
		return
	}

	fmt.Println("RSA key pair generated and saved to private_key.pem and public_key.pem")
	db.ConnectDB()
	// Initialize the Gin router
	router := gin.Default()

	// Define GET route /ping
	router.GET("/tell", handlers.Pong)
	router.POST("/signup",handlers.Signup)
	router.POST("/signin",handlers.Signin)
	router.POST("/getoption",handlers.Getoption)
	//router.POST("/analyze",handlers.Analyze)
	router.GET("/getpublickey",handlers.GetPublicKey)

	// Start the server on port 8080
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}

