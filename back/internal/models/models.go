package models

import "time"

type Base struct {
	IsError bool   `json:"isError,omitempty"`
	Message string `json:"message,omitempty"`
}

type Credentials struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

type Dish struct {
	Id           int64  `json:"id,omitempty"`
	RestaurantId int64  `json:"restaurantId,omitempty"`
	Title        string `json:"title,omitempty"`
	Description  string `json:"description,omitempty"`
	Price        int32  `json:"price,omitempty"`
}

type Order struct {
	Id          int64     `json:"id,omitempty"`
	UserId      int64     `json:"userId,omitempty"`
	StartTime   time.Time `json:"startTime,omitempty"`
	EndTime     time.Time `json:"endTime,omitempty"`
	Cost        int32     `json:"cost,omitempty"`
	CreatedTime time.Time `json:"createdTime,omitempty"`
}

type Payment struct {
	Id        int64  `json:"id,omitempty"`
	OrderId   int64  `json:"orderId,omitempty"`
	TotalCost int32  `json:"totalCost,omitempty"`
	Format    string `json:"format,omitempty"`
	Status    string `json:"status,omitempty"`
}

type Place struct {
	Id           int64   `json:"id,omitempty"`
	OrderId      int64   `json:"orderId,omitempty"`
	RestaurantId int64   `json:"restaurantId,omitempty"`
	Capacity     int32   `json:"capacity,omitempty"`
	Number       int32   `json:"number,omitempty"`
	LeftTop      float64 `json:"leftTop,omitempty"`
	RightBottom  float64 `json:"rightBottom,omitempty"`
}

type Restaurant struct {
	Id          int64     `json:"id,omitempty"`
	Title       string    `json:"title,omitempty"`
	Description string    `json:"description,omitempty"`
	Address     string    `json:"address,omitempty"`
	Number      string    `json:"number,omitempty"`
	WorkTime    time.Time `json:"workTime,omitempty"`
	Kitchen     string    `json:"kitchen,omitempty"`
	Img         string    `json:"img,omitempty"`
}
