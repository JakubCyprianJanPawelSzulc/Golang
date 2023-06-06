package types

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID      primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	Email   string             `json:"email"`
	Name    string             `json:"name"`
	Surname string             `json:"surName"`
	Postal  string             `json:"postal"`
	City    string             `json:"city"`
	Pass    string             `json:"pass"`
}

type Token struct {
	Token          string `json:"token"`
	User           User   `json:"user"`
	ExpirationDate int64  `json:"expirationDate"`
}

type Book struct {
	ID   primitive.ObjectID `bson:"_id"`
	Rate float64            `bson:"rate"`
}

type LoginUser struct {
	Email string `json:"email"`
	Pass  string `json:"pass"`
}

type TokenRequest struct {
	UserToken string `json:"userToken"`
}

type LogoutRequestBody struct {
	UserToken string `json:"userToken"`
}

type EditReservationRequestBody struct {
	ReservationId string `json:"reservationId"`
	BorrowTime    int64  `json:"borrowTime"`
}

type ReservationStatus struct {
	Status string `bson:"status"`
}

type BookAddFull struct {
	Title       string   `json:"title" bson:"title"`
	Description string   `json:"description" bson:"description"`
	DateRelease int64    `json:"dateRelease" bson:"dateRelease"`
	Genres      []string `json:"genres" bson:"genres"`
	Authors     []string `json:"authors" bson:"authors"`
	Publisher   string   `json:"publisher" bson:"publisher"`
	Quantity    int      `json:"quantity" bson:"quantity"`
	DateAdded   int64    `json:"dateAdded" bson:"dateAdded"`
	Photo       string   `json:"photo" bson:"photo"`
}

type BookFull struct {
	ID          string   `json:"_id" bson:"_id"`
	Title       string   `json:"title" bson:"title"`
	Description string   `json:"description" bson:"description"`
	DateRelease int64    `json:"dateRelease" bson:"dateRelease"`
	Genres      []string `json:"genres" bson:"genres"`
	Authors     []string `json:"authors" bson:"authors"`
	Publisher   string   `json:"publisher" bson:"publisher"`
	Quantity    int      `json:"quantity" bson:"quantity"`
	DateAdded   int64    `json:"dateAdded" bson:"dateAdded"`
	Photo       string   `json:"photo" bson:"photo"`
}

type EditUserNameRequestBody struct {
	UserID  string `json:"userId"`
	Name    string `json:"name"`
	SurName string `json:"surName"`
}

type UserCreate struct {
	Email   string `json:"email"`
	Name    string `json:"name"`
	SurName string `json:"surName" bson:"surName"`
	Postal  string `json:"postal"`
	City    string `json:"city"`
	Pass    string `json:"pass"`
}

type UpdateCommentRequestBody struct {
	ID          string `json:"_id"`
	BookID      string `json:"bookId"`
	Comment     string `json:"comment"`
	UserID      string `json:"userId"`
	UserName    string `json:"userName"`
	UserSurName string `json:"userSurName"`
}

type AddCommentByAdminRequestBody struct {
	BookID      string `json:"bookId"`
	Comment     string `json:"comment"`
	UserID      string `json:"userId"`
	UserName    string `json:"userName"`
	UserSurName string `json:"userSurName"`
}

type EditGivenReservationPayload struct {
	ReservationID string `json:"reservationId"`
	HandOverTime  int64  `json:"handOverTime"`
}

type EditReturnedReservationPayload struct {
	ReservationID string `json:"reservationId"`
	ReturnTime    int64  `json:"returnTime"`
}
