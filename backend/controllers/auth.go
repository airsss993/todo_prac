package controllers

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
	"trackit/backend/initializers"
	"trackit/backend/models"
)

func SignUp(c *gin.Context) {
	var body *models.User
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to hash password"})
	}

	body.Password = string(hash)

	userCollection := initializers.GetCollection("users")

	_, err = userCollection.InsertOne(c, bson.M{
		"name":       body.Name,
		"password":   body.Password,
		"created_at": time.Now(),
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to create user"})
	}

	c.JSON(http.StatusOK, gin.H{"message": "successfully created user"})
}
