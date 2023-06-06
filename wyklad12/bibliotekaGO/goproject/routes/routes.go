package routes

import (
	"github.com/gin-gonic/gin"

	controllers "gin/controllers"
)

func PublicRoutes(g *gin.RouterGroup) {

	loginGroup := g.Group("/login")
	{
		loginGroup.POST("", controllers.Login)
		loginGroup.POST("/logout", controllers.Logout)

	}

	register := g.Group("/register")
	{
		register.POST("", controllers.Create)
	}

}

func PrivateRoutes(g *gin.RouterGroup) {

	booksGroup := g.Group("/books")
	{
		booksGroup.GET("/getAllBooks", controllers.GetAllBooks)
		booksGroup.GET("/getAllRatesCountByBookId", controllers.GetAllRatesCountByBookId)
		booksGroup.GET("/getCommentsByBookId", controllers.GetCommentsByBookId)
		booksGroup.GET("/getBookRateById", controllers.GetBookRateById)
		booksGroup.GET("/getUserRatedBook", controllers.GetUserRatedBook)
		booksGroup.GET("/getBookById", controllers.GetBookById)
		booksGroup.GET("/getAllComments", controllers.GetAllComments)
		booksGroup.GET("/getFiveMostPopularBooks", controllers.GetFiveMostPopularBooks)
		booksGroup.GET("/getTenOldestBooks", controllers.GetTenOldestBooks)
		booksGroup.GET("/getTenMostBooks", controllers.GetTenMostBooks)

		booksGroup.POST("/addRate", controllers.AddRate)
		booksGroup.POST("/addComment", controllers.AddComment)
		booksGroup.POST("/addBook", controllers.AddBook)
		booksGroup.POST("/addCommentByAdmin", controllers.AddCommentByAdmin)

		booksGroup.PUT("/updateBook", controllers.UpdateBook)
		booksGroup.PUT("/updateComment", controllers.UpdateComment)

		booksGroup.DELETE("/deleteBookbyId", controllers.DeleteBookById)
		booksGroup.DELETE("/deleteCommentbyId", controllers.DeleteCommentById)
	}

	reservations := g.Group("/reservations")
	{
		reservations.GET("/getReservationCountByBookId", controllers.GetReservationCountByBookId)
		reservations.GET("/getAllReservations", controllers.GetAllReservations)
		reservations.GET("/getUserReservations", controllers.GetUserReservations)
		reservations.GET("/updateReservationStatus", controllers.UpdateReservationStatus)
		reservations.GET("/cancelReservation", controllers.CancelReservation)
		reservations.GET("/payReservation", controllers.PayReservation)

		reservations.POST("/createReservation", controllers.CreateReservation)
		reservations.POST("/editReservation", controllers.EditReservation)
		reservations.POST("/editGivenReservation", controllers.EditGivenReservation)
		reservations.POST("/editReturnedReservation", controllers.EditReturnedReservation)
	}

	user := g.Group("/user")
	{
		user.GET("/getAllUsers", controllers.GetAllUsers)
		user.GET("/deleteUserbyId", controllers.DeleteUserById)

		user.POST("/editUserNameAndSurname", controllers.EditUserNameAndSurname)
	}
}
