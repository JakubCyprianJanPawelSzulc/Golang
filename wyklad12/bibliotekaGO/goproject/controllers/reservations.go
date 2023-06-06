package controllers

import (
	"context"
	"gin/types"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetReservationCountByBookId(c *gin.Context) {

	bookId := c.Query("bookId")

	collection := DB.Collection("reservations")

	// Filtr dopasowujący bookId oraz unikający wartości "RETURNED" oraz "CANCELLED"
	filter := bson.M{
		"bookId": bookId,
		"status": bson.M{"$nin": []string{"RETURNED", "CANCELLED"}},
	}

	pipeline := bson.A{
		bson.M{"$match": filter},
		bson.M{
			"$group": bson.M{
				"_id":   "$bookId",
				"count": bson.M{"$sum": 1},
			},
		},
	}

	cursor, err := collection.Aggregate(context.TODO(), pipeline)
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to fetch reservation count")
		return
	}
	defer cursor.Close(context.TODO())

	var results []bson.M
	err = cursor.All(context.TODO(), &results)
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to decode reservation count")
		return
	}

	if len(results) == 0 {
		c.JSON(http.StatusOK, gin.H{"count": 0})
		return
	}

	count := results[0]["count"].(int32)

	c.JSON(http.StatusOK, gin.H{"count": count})
}

func CreateReservation(c *gin.Context) {

	user := GetUser(c)

	collection := DB.Collection("reservations")

	var reservationData bson.M
	err := c.BindJSON(&reservationData)
	if err != nil {
		c.String(http.StatusBadRequest, "Invalid reservation data")
		return
	}

	reservationData["userId"] = user.ID.Hex()
	reservationData["status"] = "NEW"
	reservationData["payed"] = false
	reservationData["cancelled"] = false
	reservationData["returnTime"] = nil
	reservationData["handOverTime"] = nil

	_, err = collection.InsertOne(context.TODO(), reservationData)
	if err != nil {
		if err == mongo.ErrNilDocument {
			c.String(http.StatusInternalServerError, "Failed to create reservation")
		} else {
			c.String(http.StatusBadRequest, "Invalid reservation data")
		}
		return
	}

	c.Status(http.StatusOK)
}

func GetUserReservations(c *gin.Context) {

	user := GetUser(c)

	collection := DB.Collection("reservations")

	pipeline := bson.A{
		bson.M{"$match": bson.M{"userId": user.ID.Hex()}},
		bson.M{"$addFields": bson.M{"idd": bson.M{"$toObjectId": "$bookId"}}},
		bson.M{
			"$lookup": bson.M{
				"from":         "books",
				"localField":   "idd",
				"foreignField": "_id",
				"as":           "book",
			},
		},
		bson.M{"$unwind": bson.M{"path": "$book", "preserveNullAndEmptyArrays": true}},
		bson.M{"$addFields": bson.M{"bookTitle": "$book.title"}},
		bson.M{
			"$lookup": bson.M{
				"from": "rates",
				"let":  bson.M{"userIdFirst": "$userId", "bookIdFirst": "$bookId"},
				"pipeline": bson.A{
					bson.M{
						"$match": bson.M{
							"$expr": bson.M{
								"$and": bson.A{
									bson.M{"$eq": bson.A{"$userId", "$$userIdFirst"}},
									bson.M{"$eq": bson.A{"$bookId", "$$bookIdFirst"}},
								},
							},
						},
					},
				},
				"as": "rated",
			},
		},
		bson.M{"$unwind": bson.M{"path": "$rated", "preserveNullAndEmptyArrays": true}},
		bson.M{"$addFields": bson.M{"ratedValue": bson.M{"$toInt": "$rated.rate"}}},
	}

	cursor, err := collection.Aggregate(context.TODO(), pipeline)
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to fetch user reservations")
		return
	}
	defer cursor.Close(context.TODO())

	var results []bson.M
	err = cursor.All(context.TODO(), &results)
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to decode user reservations")
		return
	}

	for _, reservation := range results {
		delete(reservation, "book")
		delete(reservation, "idd")
		delete(reservation, "rated")
	}

	c.JSON(http.StatusOK, results)
}

func PayReservation(c *gin.Context) {

	reservationId := c.Query("reservationId")

	collection := DB.Collection("reservations")

	objectId, err := primitive.ObjectIDFromHex(reservationId)
	if err != nil {
		c.String(http.StatusBadRequest, "Invalid reservation ID")
		return
	}

	update := bson.M{"$set": bson.M{"payed": true}}
	_, err = collection.UpdateOne(context.TODO(), bson.M{"_id": objectId}, update)
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to update reservation")
		return
	}

	c.Status(http.StatusOK)
}

func CancelReservation(c *gin.Context) {

	reservationId := c.Query("reservationId")

	collection := DB.Collection("reservations")

	objectId, err := primitive.ObjectIDFromHex(reservationId)
	if err != nil {
		c.String(http.StatusBadRequest, "Invalid reservation ID")
		return
	}

	update := bson.M{"$set": bson.M{"cancelled": true, "status": "CANCELLED"}}
	_, err = collection.UpdateOne(context.TODO(), bson.M{"_id": objectId}, update)
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to update reservation")
		return
	}

	c.Status(http.StatusOK)
}

func EditReservation(c *gin.Context) {

	var requestBody types.EditReservationRequestBody
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.String(http.StatusBadRequest, "Invalid request body")
		return
	}

	collection := DB.Collection("reservations")

	objectId, err := primitive.ObjectIDFromHex(requestBody.ReservationId)
	if err != nil {
		c.String(http.StatusBadRequest, "Invalid reservation ID")
		return
	}

	update := bson.M{"$set": bson.M{"borrowTime": requestBody.BorrowTime}}
	_, err = collection.UpdateOne(context.TODO(), bson.M{"_id": objectId}, update)
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to update reservation")
		return
	}

	c.Status(http.StatusOK)
}

func UpdateReservationStatus(c *gin.Context) {

	reservationID := c.Query("reservationId")

	statuses := []string{"NEW", "CONFIRMED", "READY", "BORROWED", "RETURNED"}

	collection := DB.Collection("reservations")

	objectID, err := primitive.ObjectIDFromHex(reservationID)
	if err != nil {
		c.String(http.StatusBadRequest, "Invalid reservation ID")
		return
	}

	var reservation types.ReservationStatus

	err = collection.FindOne(context.TODO(), bson.M{"_id": objectID}).Decode(&reservation)
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to find reservation")
		return
	}

	nextIndex := -1
	for i, status := range statuses {
		if status == reservation.Status {
			nextIndex = (i + 1) % len(statuses)
			break
		}
	}

	if nextIndex == -1 {
		c.String(http.StatusInternalServerError, "Invalid reservation status")
		return
	}

	update := bson.M{"$set": bson.M{"status": statuses[nextIndex]}}
	_, err = collection.UpdateOne(context.TODO(), bson.M{"_id": objectID}, update)
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to update reservation status")
		return
	}

	c.Status(http.StatusOK)
}

func GetAllReservations(c *gin.Context) {
	collection := DB.Collection("reservations")

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
		{
			"$addFields": bson.M{
				"user_idd": bson.M{
					"$toObjectId": "$userId",
				},
			},
		},
		{
			"$lookup": bson.M{
				"from":         "users",
				"localField":   "user_idd",
				"foreignField": "_id",
				"as":           "user",
			},
		},
		{
			"$unwind": bson.M{
				"path":                       "$user",
				"preserveNullAndEmptyArrays": true,
			},
		},
		{
			"$addFields": bson.M{
				"userName":    "$user.name",
				"userSurName": "$user.surName",
			},
		},
	}

	cursor, err := collection.Aggregate(context.TODO(), pipeline)
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to retrieve reservations")
		return
	}

	var reservations []bson.M
	if err := cursor.All(context.TODO(), &reservations); err != nil {
		c.String(http.StatusInternalServerError, "Failed to retrieve reservations")
		return
	}

	c.JSON(http.StatusOK, reservations)
}

func EditGivenReservation(c *gin.Context) {
	var payload types.EditGivenReservationPayload

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.String(http.StatusBadRequest, "Invalid payload")
		return
	}

	objectID, err := primitive.ObjectIDFromHex(payload.ReservationID)
	if err != nil {
		c.String(http.StatusBadRequest, "Invalid reservation ID")
		return
	}

	collection := DB.Collection("reservations")

	exists, err := collection.CountDocuments(context.TODO(), bson.M{"_id": objectID})
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to check reservation")
		return
	}
	if exists == 0 {
		c.String(http.StatusNotFound, "Reservation not found")
		return
	}

	filter := bson.M{"_id": objectID}
	update := bson.M{"$set": bson.M{"handOverTime": payload.HandOverTime}}

	_, err = collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to update reservation")
		return
	}

	c.String(http.StatusOK, "Reservation updated successfully")
}

func EditReturnedReservation(c *gin.Context) {
	var payload types.EditReturnedReservationPayload

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.String(http.StatusBadRequest, "Invalid payload")
		return
	}

	objectID, err := primitive.ObjectIDFromHex(payload.ReservationID)
	if err != nil {
		c.String(http.StatusBadRequest, "Invalid reservation ID")
		return
	}

	collection := DB.Collection("reservations")

	exists, err := collection.CountDocuments(context.TODO(), bson.M{"_id": objectID})
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to check reservation")
		return
	}
	if exists == 0 {
		c.String(http.StatusNotFound, "Reservation not found")
		return
	}

	filter := bson.M{"_id": objectID}
	update := bson.M{"$set": bson.M{"returnTime": payload.ReturnTime}}

	_, err = collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to update reservation")
		return
	}

	c.String(http.StatusOK, "Reservation updated successfully")
}
