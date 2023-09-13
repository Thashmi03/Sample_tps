package controller

import (
	"context"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//	type sample struct{
//		Name string `json:"name" bson:"name"`
//	}
var (
	client     *mongo.Client
	collection *mongo.Collection
	mu         sync.Mutex // Mutex for protecting shared resources
)
type Data struct {
	Name string`json:"name" bson:"name"`
	Id int32 `json:"id" bson:"id"`
}
func init() {
	// Connect to MongoDB
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/?retryWrites=true&serverSelectionTimeoutMS=5000&connectTimeoutMS=10000")
	client, _ = mongo.Connect(context.TODO(), clientOptions)

	// Check the connection
	err := client.Ping(context.Background(), nil)
	if err != nil {
		panic(err)
	}

	// Access the database and collection
	database := client.Database("db")        // Replace "mydb" with your database name
	collection = database.Collection("sample") // Replace "tokens" with your collection name
}
func Test(c *gin.Context) {
	var requestBody map[string]interface{}
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Modify the JSON data
	requestBody["name"] = requestBody["name"].(string)

	// Send the modified JSON as a response
	c.JSON(http.StatusOK, requestBody)

}
func Token(c *gin.Context) {
	token := c.GetHeader("Authorization")
	var data []Data
	if token == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Token not found in the header"})
		return
	}
	if err := c.ShouldBindJSON(&data); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
	insertData := make([]interface{}, len(data))
    for i, doc := range data {

        insertData[i] = doc
    }
	// Store the token in MongoDB
	_, err := collection.InsertOne(context.TODO(), map[string]interface{}{"token": token,"details":insertData})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Token stored successfully"})
}

func GetToken(c *gin.Context) {
	cursor, err := collection.Find(context.TODO(), nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer cursor.Close(context.TODO())

	var tokens []string

	for cursor.Next(context.TODO()) {
		var result map[string]interface{}
		if err := cursor.Decode(&result); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		tokens = append(tokens, result["token"].(string))
	}

	// Check if there are no stored tokens
	if len(tokens) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No tokens found"})
		return
	}

	// Send the retrieved tokens as a response
	c.JSON(http.StatusOK, gin.H{"tokens": tokens})
}
