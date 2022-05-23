package models

import (
	"time"
)

// type CustomTime struct {
// 	time.Time
// }

// const ctLayout = "15:04"

// func (ct *CustomTime) UnmarshalJSON(b []byte) (err error) {
// 	if b[0] == '"' && b[len(b)-1] == '"' {
// 		b = b[1 : len(b)-1]
// 	}
// 	ct.Time, err = time.Parse(ctLayout, string(b))
// 	return
// }

// func (ct *CustomTime) MarshalJSON() ([]byte, error) {
// 	return []byte(ct.Time.Format(ctLayout)), nil
// }

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
	Id           uint64 `json:"id,omitempty"`
	Email        string `json:"email,omitempty"`
	Password     string `json:"password,omitempty"`
	FullName     string `json:"fullName,omitempty"`
	Number       string `json:"number,omitempty"`
	IsAdmin      bool   `json:"isAdmin,omitempty"`
	IsRestaurant bool   `json:"isRestaurant,omitempty"`
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
type DishOrder struct {
	Id     uint64 `json:"id,omitempty"`
	Number int32  `json:"number"`
}

//easyjson:json
type OrderList []Order

//easyjson:json
type Order struct {
	Id                uint64    `json:"id,omitempty"`
	UserId            uint64    `json:"userId,omitempty"`
	PlaceId           uint64    `json:"placeId,omitempty"`
	RestaurantTitle   string    `json:"restaurantTitle,omitempty"`
	RestaurantAddress string    `json:"restaurantAddress,omitempty"`
	RestaurantMetro   string    `json:"restaurantMetro,omitempty"`
	PlaceNumber       int32     `json:"placeNumber,omitempty"`
	PlaceCapacity     int32     `json:"placeCapacity,omitempty"`
	Dishes            []string  `json:"dishes,omitempty"`
	DishesPrices      []int32   `json:"dishesPrices,omitempty"`
	DishesCounts      []int32   `json:"dishesCounts,omitempty"`
	StartTime         time.Time `json:"startTime,omitempty"`
	EndTime           time.Time `json:"endTime,omitempty"`
	Cost              int32     `json:"cost,omitempty"`
	CreatedTime       time.Time `json:"createdTime,omitempty"`
	Status            string    `json:"status,omitempty"`
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
	Id           uint64 `json:"id,omitempty"`
	RestaurantId uint64 `json:"restaurantId,omitempty"`
	Capacity     int32  `json:"capacity,omitempty"`
	Number       int32  `json:"number,omitempty"`
	LeftTop      int32  `json:"leftTop,omitempty"`
	RightBottom  int32  `json:"rightBottom,omitempty"`
	Width        int32  `json:"width,omitempty"`
	Height       int32  `json:"height,omitempty"`
	Floor        int32  `json:"floor,omitempty"`
	IsBooked     bool   `json:"isBooked"`
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
	Metro       string    `json:"metro,omitempty"`
	Number      string    `json:"number,omitempty"`
	OpenTime    time.Time `json:"openTime,omitempty"`
	CloseTime   time.Time `json:"closeTime,omitempty"`
	Kitchen     string    `json:"kitchen,omitempty"`
	Img         string    `json:"img,omitempty"`
}
