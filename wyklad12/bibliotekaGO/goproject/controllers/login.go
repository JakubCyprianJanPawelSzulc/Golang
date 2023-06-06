package controllers

import (
	"context"
	"crypto/sha256"
	"fmt"
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"

	types "gin/types"
)

func Login(c *gin.Context) {
	var userLoginCredentials types.LoginUser
	var user types.User
	if err := c.ShouldBindJSON(&userLoginCredentials); err != nil {
		c.String(http.StatusBadRequest, "Invalid request")
		return
	}

	collection := DB.Collection("users")
	err := collection.FindOne(context.TODO(), userLoginCredentials).Decode(&user)
	// fmt.Println(user.ID.Hex())
	if err != nil {
		c.String(http.StatusUnauthorized, "")
		return
	}

	token := generateToken()

	Authorize(token, user)

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func generateToken() string {
	token := sha256.Sum256([]byte(fmt.Sprintf("%f", rand.Float64())))
	return fmt.Sprintf("%x", token)
}

func Logout(c *gin.Context) {

	var requestBody types.LogoutRequestBody
	if err := c.BindJSON(&requestBody); err != nil {
		c.String(http.StatusBadRequest, "Invalid request")
		return
	}

	userToken := requestBody.UserToken

	deAuthStatus := DeAuthorize(userToken)

	if !deAuthStatus {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Token not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Token deauthorized successfully",
	})
}

func Create(c *gin.Context) {

	var user types.UserCreate
	if err := c.ShouldBindJSON(&user); err != nil {
		c.String(http.StatusBadRequest, "Invalid request body")
		return
	}

	collection := DB.Collection("users")

	_, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to create user")
		return
	}

	c.String(http.StatusOK, "ok")
}
