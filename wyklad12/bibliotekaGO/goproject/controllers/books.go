package controllers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	types "gin/types"
)

func GetAllBooks(c *gin.Context) {
	// Uzyskiwanie dostępu do kolekcji "books"
	collection := DB.Collection("books")

	// Zwrócenie wszystkich dokuemntów w kolekcji
	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to fetch books")
		return
	}
	defer cursor.Close(context.TODO())

	// Iteracja po kursorze i pobranie dokumentów
	var books []bson.M
	for cursor.Next(context.TODO()) {
		var book bson.M
		err := cursor.Decode(&book)
		if err != nil {
			c.String(http.StatusInternalServerError, "Failed to decode books")
			return
		}
		books = append(books, book)
	}

	// Zwraca książki jako odpowiedź JSON
	c.JSON(http.StatusOK, books)
}

func GetAllRatesCountByBookId(c *gin.Context) {

	// Uzyskiwanie dostępu do kolekcji "rates"
	collection := DB.Collection("rates")

	pipeline := bson.A{
		bson.M{
			"$group": bson.M{
				"_id":  "$bookId",
				"rate": bson.M{"$avg": bson.M{"$toInt": "$rate"}},
			},
		},
	}

	// wykonanie agregacji
	cursor, err := collection.Aggregate(context.TODO(), pipeline)
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to fetch rates count")
		return
	}
	defer cursor.Close(context.TODO())

	var rates []bson.M
	for cursor.Next(context.TODO()) {
		var rate bson.M
		err := cursor.Decode(&rate)
		if err != nil {
			c.String(http.StatusInternalServerError, "Failed to decode rates count")
			return
		}
		rates = append(rates, rate)
	}

	c.JSON(http.StatusOK, rates)
}

func GetCommentsByBookId(c *gin.Context) {
	// Uzyskaj pole bookId z zapytania
	bookId := c.Query("bookId")

	collection := DB.Collection("comments")

	// zdefiniowanie filtra
	filter := bson.M{"bookId": bookId}

	// Komentarze które pasują do filtra
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to fetch comments")
		return
	}
	defer cursor.Close(context.TODO())

	var comments []bson.M
	for cursor.Next(context.TODO()) {
		var comment bson.M
		err := cursor.Decode(&comment)
		if err != nil {
			c.String(http.StatusInternalServerError, "Failed to decode comments")
			return
		}
		comments = append(comments, comment)
	}

	c.JSON(http.StatusOK, comments)
}

func GetUserRatedBook(c *gin.Context) {

	bookId := c.Query("bookId")

	// Uzyskaj userId z tokena autoryzacyjnego
	userId := GetUser(c).ID.Hex()

	collection := DB.Collection("rates")

	filter := bson.M{"bookId": bookId, "userId": userId}

	var rate bson.M
	err := collection.FindOne(context.TODO(), filter).Decode(&rate)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			rate = nil
		} else {
			c.String(http.StatusInternalServerError, "Failed to fetch user rate")
			return
		}
	}

	c.JSON(http.StatusOK, rate)
}

func GetBookRateById(c *gin.Context) {

	bookID := c.Query("bookId")

	// Parsowanie identyfikatora książki jako MongoDB ObjectID
	objID, err := primitive.ObjectIDFromHex(bookID)
	if err != nil {
		c.String(http.StatusBadRequest, "Invalid book ID")
		return
	}

	collection := DB.Collection("books")

	// Filtr wybierający identyfikator książki
	filter := bson.M{
		"_id": objID,
	}

	// Znajdz ksiazke po id
	var book types.Book
	err = collection.FindOne(context.TODO(), filter).Decode(&book)
	if err != nil {
		c.String(http.StatusNotFound, "Book not found")
		return
	}

	// Obsluga wartości oceny NaN
	if isNaN(book.Rate) {
		book.Rate = 0
	}

	c.JSON(http.StatusOK, gin.H{"rate": book.Rate})
}

func isNaN(f float64) bool {
	return f != f
}

func GetBookById(c *gin.Context) {

	bookID := c.Query("bookId")

	objID, err := primitive.ObjectIDFromHex(bookID)
	if err != nil {
		c.String(http.StatusBadRequest, "Invalid book ID")
		return
	}

	collection := DB.Collection("books")

	filter := bson.M{
		"_id": objID,
	}

	var book bson.M
	err = collection.FindOne(context.TODO(), filter).Decode(&book)
	if err != nil {
		c.String(http.StatusNotFound, "Book not found")
		return
	}

	c.JSON(http.StatusOK, book)
}

func AddRate(c *gin.Context) {

	userId := GetUserId(c)

	collection := DB.Collection("rates")

	var rateData bson.M
	err := c.BindJSON(&rateData)
	if err != nil {
		c.String(http.StatusBadRequest, "Invalid rate data")
		return
	}

	// Dodanie userId do danych
	rateData["userId"] = userId

	// Wstawianie oceny
	_, err = collection.InsertOne(context.TODO(), rateData)
	if err != nil {
		if err == mongo.ErrNilDocument {
			c.String(http.StatusInternalServerError, "Failed to add rate")
		} else {
			c.String(http.StatusBadRequest, "Invalid rate data")
		}
		return
	}

	c.Status(http.StatusCreated)
}

func AddComment(c *gin.Context) {

	user := GetUser(c)

	collection := DB.Collection("comments")

	var commentData bson.M
	err := c.BindJSON(&commentData)
	if err != nil {
		c.String(http.StatusBadRequest, "Invalid comment data")
		return
	}

	commentData["userId"] = user.ID.Hex()
	commentData["userName"] = user.Name
	commentData["userSurName"] = user.Surname

	_, err = collection.InsertOne(context.TODO(), commentData)
	if err != nil {
		if err == mongo.ErrNilDocument {
			c.String(http.StatusInternalServerError, "Failed to add comment")
		} else {
			c.String(http.StatusBadRequest, "Invalid comment data")
		}
		return
	}

	c.Status(http.StatusCreated)
}

func AddBook(c *gin.Context) {
	var book types.BookAddFull

	if err := c.ShouldBindJSON(&book); err != nil {
		c.String(http.StatusBadRequest, "Invalid request body")
		return
	}

	collection := DB.Collection("books")

	_, err := collection.InsertOne(context.TODO(), book)
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to add book")
		return
	}

	c.Status(http.StatusCreated)
}

func UpdateBook(c *gin.Context) {

	var requestBody types.BookFull
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.String(http.StatusBadRequest, "Invalid request body")
		return
	}

	collection := DB.Collection("books")

	objectId, err := primitive.ObjectIDFromHex(requestBody.ID)
	if err != nil {
		c.String(http.StatusBadRequest, "Invalid book ID")
		return
	}

	// Del Id
	updateData := bson.M{
		"title":       requestBody.Title,
		"description": requestBody.Description,
		"dateRelease": requestBody.DateRelease,
		"genres":      requestBody.Genres,
		"authors":     requestBody.Authors,
		"publisher":   requestBody.Publisher,
		"quantity":    requestBody.Quantity,
		"dateAdded":   requestBody.DateAdded,
		"photo":       requestBody.Photo,
	}

	update := bson.M{"$set": updateData}
	result, err := collection.UpdateOne(context.TODO(), bson.M{"_id": objectId}, update)
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to update book")
		return
	}

	if result.ModifiedCount == 0 {
		c.Status(http.StatusNotFound)
		return
	}

	c.Status(http.StatusCreated)
}

func DeleteBookById(c *gin.Context) {
	bookID := c.Query("bookId")

	collection := DB.Collection("books")

	objectID, err := primitive.ObjectIDFromHex(bookID)
	if err != nil {
		c.String(http.StatusBadRequest, "Invalid book ID")
		return
	}

	_, err = collection.DeleteOne(context.TODO(), bson.M{"_id": objectID})
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to delete book")
		return
	}

	c.Status(http.StatusOK)
}

func GetAllComments(c *gin.Context) {

	collection := DB.Collection("comments")

	pipeline := []bson.M{
		{
			"$match": bson.M{},
		},
		{
			"$addFields": bson.M{
				"idd": bson.M{
					"$toObjectId": "$bookId",
				},
			},
		},
		{
			"$lookup": bson.M{
				"from":         "books",
				"localField":   "idd",
				"foreignField": "_id",
				"as":           "book",
			},
		},
		{
			"$unwind": bson.M{
				"path":                       "$book",
				"preserveNullAndEmptyArrays": true,
			},
		},
		{
			"$addFields": bson.M{
				"bookTitle": "$book.title",
			},
		},
	}

	cursor, err := collection.Aggregate(context.TODO(), pipeline)
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to retrieve comments")
		return
	}

	var comments []bson.M
	if err := cursor.All(context.TODO(), &comments); err != nil {
		c.String(http.StatusInternalServerError, "Failed to retrieve comments")
		return
	}

	c.JSON(http.StatusOK, comments)
}

func DeleteCommentById(c *gin.Context) {
	commentID := c.Query("commentId")

	collection := DB.Collection("comments")

	objectID, err := primitive.ObjectIDFromHex(commentID)
	if err != nil {
		c.String(http.StatusBadRequest, "Invalid comment ID")
		return
	}

	_, err = collection.DeleteOne(context.TODO(), bson.M{"_id": objectID})
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to delete comment")
		return
	}

	c.Status(http.StatusOK)
}

func UpdateComment(c *gin.Context) {
	var requestBody types.UpdateCommentRequestBody
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.String(http.StatusBadRequest, "Invalid request body")
		return
	}

	collection := DB.Collection("comments")

	objectID, err := primitive.ObjectIDFromHex(requestBody.ID)
	if err != nil {
		c.String(http.StatusBadRequest, "Invalid comment ID")
		return
	}

	update := bson.M{
		"$set": bson.M{
			"bookId":      requestBody.BookID,
			"comment":     requestBody.Comment,
			"userId":      requestBody.UserID,
			"userName":    requestBody.UserName,
			"userSurName": requestBody.UserSurName,
		},
	}

	_, err = collection.UpdateOne(context.TODO(), bson.M{"_id": objectID}, update)
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to update comment")
		return
	}

	c.Status(http.StatusOK)
}

func AddCommentByAdmin(c *gin.Context) {
	var requestBody types.AddCommentByAdminRequestBody

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.String(http.StatusBadRequest, "Invalid request body")
		return
	}

	collection := DB.Collection("comments")

	_, err := collection.InsertOne(context.TODO(), bson.M{
		"bookId":      requestBody.BookID,
		"comment":     requestBody.Comment,
		"userId":      requestBody.UserID,
		"userName":    requestBody.UserName,
		"userSurName": requestBody.UserSurName,
	})
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to add comment")
		return
	}

	c.Status(http.StatusCreated)
}

func GetTenMostBooks(c *gin.Context) {
	collection := DB.Collection("books")

	pipeline := []bson.M{
		{
			"$sort": bson.M{
				"quantity": -1,
			},
		},
		{
			"$limit": 10,
		},
	}

	cursor, err := collection.Aggregate(context.TODO(), pipeline)
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to retrieve books")
		return
	}

	var books []bson.M
	if err := cursor.All(context.TODO(), &books); err != nil {
		c.String(http.StatusInternalServerError, "Failed to retrieve books")
		return
	}

	c.JSON(http.StatusOK, books)
}

func GetFiveMostPopularBooks(c *gin.Context) {
	collection := DB.Collection("books")

	pipeline := []bson.M{
		{
			"$match": bson.M{},
		},
		{
			"$addFields": bson.M{
				"idd": bson.M{
					"$toString": "$_id",
				},
			},
		},
		{
			"$lookup": bson.M{
				"from":         "reservations",
				"localField":   "idd",
				"foreignField": "bookId",
				"as":           "res",
			},
		},
		{
			"$addFields": bson.M{
				"reservationCount": bson.M{
					"$size": "$res",
				},
			},
		},
		{
			"$sort": bson.M{
				"reservationCount": -1,
			},
		},
		{
			"$limit": 5,
		},
	}

	cursor, err := collection.Aggregate(context.TODO(), pipeline)
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to retrieve books")
		return
	}

	var books []bson.M
	if err := cursor.All(context.TODO(), &books); err != nil {
		c.String(http.StatusInternalServerError, "Failed to retrieve books")
		return
	}

	for i := range books {
		delete(books[i], "idd")
		delete(books[i], "res")
	}

	c.JSON(http.StatusOK, books)
}

func GetTenOldestBooks(c *gin.Context) {
	collection := DB.Collection("books")

	pipeline := []bson.M{
		{
			"$sort": bson.M{
				"dateRelease": 1,
			},
		},
		{
			"$limit": 10,
		},
	}

	cursor, err := collection.Aggregate(context.TODO(), pipeline)
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to retrieve books")
		return
	}

	var books []bson.M
	if err := cursor.All(context.TODO(), &books); err != nil {
		c.String(http.StatusInternalServerError, "Failed to retrieve books")
		return
	}

	c.JSON(http.StatusOK, books)
}
