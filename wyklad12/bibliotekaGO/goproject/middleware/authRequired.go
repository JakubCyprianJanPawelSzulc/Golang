package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"

	controllers "gin/controllers"
)

func AuthMiddleware(c *gin.Context) {

	userToken := c.GetHeader("userToken")

	if controllers.IsTokenAuthorized(userToken) {
		// Przejście do następnego middleware'u
		c.Next()
	} else {
		// Przerwanie z nieautoryzowanym statusem
		c.AbortWithStatus(http.StatusUnauthorized)
	}

}
