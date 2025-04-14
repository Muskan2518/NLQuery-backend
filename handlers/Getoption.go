package handlers
import(
	"context"
	"crypto/sha256"
	"encoding/hex"
	"net/http"
	"time"

	"NLQuery-backend/db" // Import your db package
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

type Urloption struct{
	token string `json:"token"`
	url string `json:"url"`
}
func Getoption(c *gin.Context){
	var urlUser Urloption
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

}

