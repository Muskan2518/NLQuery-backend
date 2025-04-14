package handlers

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os/exec"
	"context"


	"github.com/gin-gonic/gin"
	"NLQuery-backend/db"


	"NLQuery-backend/crypto"

)
type Analyzereq struct{
	Token string `json:"token"`
	Url   string `json:"url"`
	Name        string   `json:"name"`
	Collections string `json:"collections"`
}
func Analyze(c *gin.Context){
	var urlUser Analyzereq
	if err := c.ShouldBindJSON(&urlUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	publicKey, err := crypto.LoadPublicKey("public_key.pem")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load public key"})
		return
	}
	_, err = crypto.Validate_jwt(publicKey, urlUser.Token)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JWT token validation failed"})
		return
	}
	cmd := exec.Command("python3", "analyze.py", urlUser.Url,urlUser.Name,urlUser.Collections)
output, err := cmd.Output()
if err != nil {
	log.Fatalf("Error running Python script: %v", err)
}

encodedPDF := string(output) // Base64-encoded PDF
// Get MongoDB collection
collection := db.Client.Database("nlquery").Collection("pdf")
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()
var existing Request
	err := collection.FindOne(ctx, bson.M{"username": newUser.Username}).Decode(&existing)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User doesn't exist, please sign up"})
		return
	}


}