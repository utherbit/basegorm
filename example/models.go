package main

import "basegorm"

type User struct {
	basegorm.BaseModel[int]
	Username string
}

type Order struct {
	basegorm.BaseModel[int]
	UserId int
	User   *User        `gorm:"foreignKey:UserId"`
	Items  []*OrderItem `gorm:"foreignKey:OrderId"`
}
type OrderItem struct {
	basegorm.BaseModel[int]
	OrderId int
	Order   *Order `gorm:"foreignKey:OrderId"`
	ItemId  int
	Item    *Item `gorm:"foreignKey:ItemId"`
	Amount  int
}
type Item struct {
	basegorm.BaseModel[int]
	Title       string
	Description string
	Cost        int
}
