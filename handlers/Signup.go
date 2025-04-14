package handlers

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"net/http"
	"time"

	"NLQuery-backend/db" // Import your db package
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

type User struct {
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
	Email    string `json:"email" bson:"email"`
}

func Signup(c *gin.Context) {
	var newUser User

	// Bind JSON input to struct
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Hash the password using SHA-256
	hashedPassword := sha256.New()
	hashedPassword.Write([]byte(newUser.Password))
	newUser.Password = hex.EncodeToString(hashedPassword.Sum(nil))

	// Get MongoDB collection
	collection := db.Client.Database("nlquery").Collection("users")

	// Check if user already exists
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var existing User
	err := collection.FindOne(ctx, bson.M{"username": newUser.Username}).Decode(&existing)
	if err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "User already exists"})
		return
	}
	if newUser.Username==""{
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	if newUser.Password==""{
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	if newUser.Email==""{
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Insert the new user
	_, err = collection.InsertOne(ctx, newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User signed up successfully"})
}
