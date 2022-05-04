package models

import "time"

//easyjson:json
type Base struct {
	IsError bool   `json:"isError,omitempty"`
	Message string `json:"message,omitempty"`
}

//easyjson:json
type UserCredentials struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

//easyjson:json
type User struct {
	Id       uint64 `json:"id,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
	FullName string `json:"fullName,omitempty"`
	Number   string `json:"number,omitempty"`
}

//easyjson:json
type Session struct {
	Cookie string
	UserID uint64
}

//easyjson:json
type DishList []Dish

//easyjson:json
type Dish struct {
	Id           uint64 `json:"id,omitempty"`
	RestaurantId uint64 `json:"restaurantId,omitempty"`
	Title        string `json:"title,omitempty"`
	Description  string `json:"description,omitempty"`
	Price        int32  `json:"price,omitempty"`
}

//easyjson:json
type Order struct {
	Id          uint64    `json:"id,omitempty"`
	UserId      uint64    `json:"userId,omitempty"`
	StartTime   time.Time `json:"startTime,omitempty"`
	EndTime     time.Time `json:"endTime,omitempty"`
	Cost        int32     `json:"cost,omitempty"`
	CreatedTime time.Time `json:"createdTime,omitempty"`
}

//easyjson:json
type Payment struct {
	Id        uint64 `json:"id,omitempty"`
	OrderId   uint64 `json:"orderId,omitempty"`
	TotalCost int32  `json:"totalCost,omitempty"`
	Format    string `json:"format,omitempty"`
	Status    string `json:"status,omitempty"`
}

//easyjson:json
type PlaceList []Place

//easyjson:json
type Place struct {
	Id           uint64  `json:"id,omitempty"`
	RestaurantId uint64  `json:"restaurantId,omitempty"`
	Capacity     int32   `json:"capacity,omitempty"`
	Number       int32   `json:"number,omitempty"`
	LeftTop      float64 `json:"leftTop,omitempty"`
	RightBottom  float64 `json:"rightBottom,omitempty"`
	Floor        int32   `json:"floor,omitempty"`
	IsBooked     bool    `json:"isBooked"`
}

//easyjson:json
type PlaceParameters struct {
	RestaurantId uint64    `json:"restaurantId,omitempty"`
	Time         time.Time `json:"time,omitempty"`
	Floor        int32     `json:"floor,omitempty"`
}

//easyjson:json
type RestaurantList []Restaurant

//easyjson:json
type Restaurant struct {
	Id          uint64    `json:"id,omitempty"`
	Title       string    `json:"title,omitempty"`
	Description string    `json:"description,omitempty"`
	Address     string    `json:"address,omitempty"`
	Number      string    `json:"number,omitempty"`
	OpenTime    time.Time `json:"openTime,omitempty"`
	CloseTime   time.Time `json:"closeTime,omitempty"`
	Kitchen     string    `json:"kitchen,omitempty"`
	Img         string    `json:"img,omitempty"`
}
