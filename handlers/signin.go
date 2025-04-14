package handlers

import (
	"context"
	"crypto/sha256"

	"encoding/hex"
	"net/http"
	"time"
	"NLQuery-backend/crypto"
	"NLQuery-backend/db" // Import your db package
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

type Request struct {
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
}

func Signin(c *gin.Context) {
	var newUser Request

	// Bind JSON input to struct
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Hash the password using SHA-256 (for comparison)
	hashedPassword := sha256.New()
	hashedPassword.Write([]byte(newUser.Password))
	hashedPasswordString := hex.EncodeToString(hashedPassword.Sum(nil))

	// Get MongoDB collection
	collection := db.Client.Database("nlquery").Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Check if the user exists
	var existing Request
	err := collection.FindOne(ctx, bson.M{"username": newUser.Username}).Decode(&existing)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User doesn't exist, please sign up"})
		return
	}

	// Compare the hashed password
	if hashedPasswordString != existing.Password {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username or password is invalid"})

		return
	}

	// If successful, return success response
	valid,err:=crypto.LoadPrivateKey("private_key.pem")

	token,err:=crypto.GenerateJWT(valid,newUser.Username)
	c.JSON(http.StatusOK,gin.H{"token": token})
	return 





}
