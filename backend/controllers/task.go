package controllers

//func CreateTask(c *gin.Context) {
//	var body models.Task
//
//	if c.ShouldBindJSON(&body) != nil {
//		c.JSON(400, gin.H{"error": "invalid request body"})
//		return
//	}
//
//	user, exists := c.Get("user")
//	if !exists {
//		c.JSON(400, gin.H{"error": "unknown user"})
//		return
//	}
//
//	userOwnID := user.(models.User)
//
//	taskCollection := initializers.GetCollection("tasks")
//	task := models.Task{
//		Name:        body.Name,
//		Description: body.Description,
//		OwnerID:     userOwnID.ID,
//		CreatedAt:   time.Now(),
//		UpdatedAt:   time.Time{},
//	}
//
//	_, err := taskCollection.InsertOne(c, task)
//	if err != nil {
//		c.JSON(500, gin.H{"error": "failed to create task"})
//		return
//	}
//
//	c.JSON(200, gin.H{"task": task})
//}
