package handlers

import (
	"net/http"
	"NLQuery-backend/utils"

	"NLQuery-backend/crypto"

	"github.com/gin-gonic/gin"
)

type Urloption struct {
	Token string `json:"token"`
	Url   string `json:"url"`
}

func Getoption(c *gin.Context) {
	var urlUser Urloption
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

	
	//calling utils.go
	result,err:=utils.GetDatabasesAndCollections(urlUser.Url)
	if err!=nil{
		c.JSON(http.StatusExpectationFailed,gin.H{"err":"Invalid Url"})

		return
	}
	
	c.JSON(http.StatusAccepted,gin.H{"result":result})

	
}
