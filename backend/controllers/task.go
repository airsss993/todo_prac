package controllers

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
	"trackit/backend/initializers"
	"trackit/backend/models"
)

func CreateTask(c *gin.Context) {
	type TaskRequest struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		OwnerID     string `json:"owner_id"`
	}

	var body TaskRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": "invalid request body"})
		return
	}

	ownerID, err := primitive.ObjectIDFromHex(body.OwnerID)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid owner_id format"})
		return
	}

	taskCollection := initializers.GetCollection("tasks")
	task := models.Task{
		ID:          primitive.NewObjectID(),
		Name:        body.Name,
		Description: body.Description,
		OwnerID:     ownerID,
		CreatedAt:   time.Now(),
	}

	_, err = taskCollection.InsertOne(c, task)
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to create task", "details": err.Error()})
		return
	}

	c.JSON(200, gin.H{"task": task})
}

func GetTasks(c *gin.Context) {
	taskCollection := initializers.GetCollection("tasks")
	var tasks []models.Task

	cursor, err := taskCollection.Find(c, bson.M{})
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	if err = cursor.All(c, &tasks); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return

	}

	c.JSON(200, gin.H{"task": tasks})
}

func UpdateTask(c *gin.Context) {
	var body models.Task

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": "invalid request body"})
		return
	}

	coll := initializers.GetCollection("tasks")

	filter := bson.M{"_id": body.ID}
	update := bson.M{
		"$set": bson.M{
			"name":        body.Name,
			"description": body.Description,
			"updated_at":  time.Now(),
		},
	}

	result, err := coll.UpdateOne(c, filter, update)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if result.MatchedCount == 0 {
		c.JSON(400, gin.H{"error": "task not found"})
	}
	c.JSON(200, gin.H{"task": body})
}
