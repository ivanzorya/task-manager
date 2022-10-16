package routes

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"server/models"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/bson"
	"github.com/gin-gonic/gin"
)

var validate = validator.New()

var taskCollection *mongo.Collection = OpenCollection(Client, "tasks")

func AddTask(c *gin.Context) {

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	var task models.Task

	if err := c.BindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	validationErr := validate.Struct(task)
	if validationErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
		fmt.Println(validationErr)
		return
	}
	task.ID = primitive.NewObjectID()

	result, insertErr := taskCollection.InsertOne(ctx, task)
	if insertErr != nil {
		msg := fmt.Sprintf("task item was not created")
		c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
		fmt.Println(insertErr)
		return
	}
	defer cancel()

	c.JSON(http.StatusOK, result)
}

func GetTasks(c *gin.Context){

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	
	var tasks []bson.M

	cursor, err := taskCollection.Find(ctx, bson.M{})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}
	
	if err = cursor.All(ctx, &tasks); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	defer cancel()

	fmt.Println(tasks)

	c.JSON(http.StatusOK, tasks)
}

func GetTaskById(c *gin.Context){

	taskID := c.Params.ByName("id")
	docID, _ := primitive.ObjectIDFromHex(taskID)

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	var task bson.M

	if err := taskCollection.FindOne(ctx, bson.M{"_id": docID}).Decode(&task); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	defer cancel()

	fmt.Println(task)

	c.JSON(http.StatusOK, task)
}

func UpdateTask(c *gin.Context){

	taskID := c.Params.ByName("id")
	docID, _ := primitive.ObjectIDFromHex(taskID)

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	var task models.Task

	if err := c.BindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	validationErr := validate.Struct(task)
	if validationErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
		fmt.Println(validationErr)
		return
	}

	result, err := taskCollection.ReplaceOne(
		ctx,
		bson.M{"_id": docID},
		bson.M{
			"subject":  task.Subject,
			"done": task.Done,
		},
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	defer cancel()

	c.JSON(http.StatusOK, result.ModifiedCount)
}

func DeleteTask(c * gin.Context){
	
	taskID := c.Params.ByName("id")
	docID, _ := primitive.ObjectIDFromHex(taskID)

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	result, err := taskCollection.DeleteOne(ctx, bson.M{"_id": docID})
	
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	defer cancel()

	c.JSON(http.StatusOK, result.DeletedCount)

}
