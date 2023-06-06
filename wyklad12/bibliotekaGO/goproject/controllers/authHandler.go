package controllers

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"

	types "gin/types"
)

var authorizedTokens []types.Token

func Authorize(token string, user types.User) {

	authorizedTokens = append(authorizedTokens, types.Token{
		Token:          token,
		User:           user,
		ExpirationDate: getTomorrowTimestamp(),
	})

	PrintAuthorized()

}

func PrintAuthorized() {
	// Na potrzeby testowania, wyświetl autoryzowane tokeny
	for _, token := range authorizedTokens {
		fmt.Printf("Token: %s\nUser: %+v\nExpiration Date: %d\n\n", token.Token, token.User, token.ExpirationDate)
	}
}

func IsTokenAuthorized(token string) bool {
	for _, auth := range authorizedTokens {
		if auth.Token == token {
			return true
		}
	}

	return false
}

func GetUser(c *gin.Context) *types.User {
	token := c.GetHeader("userToken")

	for _, auth := range authorizedTokens {
		if auth.Token == token {
			return &auth.User
		}
	}
	return nil
}

func GetUserId(c *gin.Context) string {
	user := GetUser(c)
	if user != nil {
		return user.ID.Hex()
	}
	return ""
}

func getTomorrowTimestamp() int64 {
	tomorrow := time.Now().Add(24 * time.Hour)
	return tomorrow.Unix()
}

func DeAuthorize(userToken string) bool {

	// Usuwanie tokenów pasujących do userToken
	updatedTokens := []types.Token{}
	found := false
	for _, token := range authorizedTokens {
		if token.Token != userToken {
			updatedTokens = append(updatedTokens, token)
		} else {
			found = true
		}
	}

	authorizedTokens = updatedTokens

	PrintAuthorized()

	return found
}
