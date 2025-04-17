package handlers

import (
	"net/http"
	"NLQuery-backend/crypto"
	"NLQuery-backend/utils"

	"github.com/gin-gonic/gin"
)

func GetPublicKey(c *gin.Context) {
	publicKey, err := crypto.LoadPublicKey("public_key.pem")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid public key"})
		return
	}

	serializedKey, err := utils.PublicKeyToString(publicKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not serialize public key"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"publickey": serializedKey})
}
