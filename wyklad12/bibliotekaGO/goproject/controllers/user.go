package controllers

import (
	"context"
	types "gin/types"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetAllUsers(c *gin.Context) {

	collection := DB.Collection("users")

	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to retrieve users")
		return
	}
	defer cursor.Close(context.TODO())

	var users []types.User
	for cursor.Next(context.TODO()) {
		var user types.User
		if err := cursor.Decode(&user); err != nil {
			c.String(http.StatusInternalServerError, "Failed to decode users")
			return
		}
		users = append(users, user)
	}

	if err := cursor.Err(); err != nil {
		c.String(http.StatusInternalServerError, "Failed to retrieve users")
		return
	}

	c.JSON(http.StatusOK, users)
}

func DeleteUserById(c *gin.Context) {

	userID := c.Query("userId")

	collection := DB.Collection("users")

	objectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		c.String(http.StatusBadRequest, "Invalid user ID")
		return
	}

	_, err = collection.DeleteOne(context.TODO(), bson.M{"_id": objectID})
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to delete user")
		return
	}

	c.Status(http.StatusOK)
}

func EditUserNameAndSurname(c *gin.Context) {

	var requestBody types.EditUserNameRequestBody
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.String(http.StatusBadRequest, "Invalid request body")
		return
	}

	collection := DB.Collection("users")

	objectID, err := primitive.ObjectIDFromHex(requestBody.UserID)
	if err != nil {
		c.String(http.StatusBadRequest, "Invalid user ID")
		return
	}

	update := bson.M{"$set": bson.M{"name": requestBody.Name, "surName": requestBody.SurName}}
	_, err = collection.UpdateOne(context.TODO(), bson.M{"_id": objectID}, update)
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to update user")
		return
	}

	c.Status(http.StatusOK)
}
