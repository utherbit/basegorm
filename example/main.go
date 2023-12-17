package main

import (
	"context"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var pgDsn = "host=192.168.85.249 user=dba password=qwerty123 dbname=base_gorm_test port=5432 sslmode=disable"

func main() {
	db, err := connect(pgDsn)
	panicIfErr(err)

	repo, err := newRepository(db)
	panicIfErr(err)

	user, err := repo.userRepo.AddUser(context.Background(), "test")
	panicIfErr(err)

	item1, err := repo.itemRepo.AddItem(context.Background(), "test1", "test 1", 100)
	panicIfErr(err)
	item2, err := repo.itemRepo.AddItem(context.Background(), "test2", "test 2", 150)
	panicIfErr(err)
	item3, err := repo.itemRepo.AddItem(context.Background(), "test3", "test 3", 200)
	panicIfErr(err)

	order, err := repo.orderRepo.AddOrder(context.Background(), user.ID)
	panicIfErr(err)

	_, err = repo.orderItemRepo.AddOrderItem(context.Background(), order.ID, item1.ID, 10)
	panicIfErr(err)
	_, err = repo.orderItemRepo.AddOrderItem(context.Background(), order.ID, item2.ID, 15)
	panicIfErr(err)
	_, err = repo.orderItemRepo.AddOrderItem(context.Background(), order.ID, item3.ID, 20)
	panicIfErr(err)

	getOrder, err := repo.orderRepo.GetOrder(context.Background(), order.ID, true)
	panicIfErr(err)

	// Выведет заказ и все товары которые он содержит
	/*
		ORDER: id: 1 userId: 1
		ORDER ITEM title: test1, amount: 10
		ORDER ITEM title: test2, amount: 15
		ORDER ITEM title: test3, amount: 20
	*/
	fmt.Printf("\nORDER: id: %v userId: %v", getOrder.ID, getOrder.UserId)
	for _, item := range getOrder.Items {
		fmt.Printf("\nORDER ITEM title: %v, amount: %v", item.Item.Title, item.Amount)
	}

}
func panicIfErr(err error) {
	if err != nil {
		panic(err)
	}
}

func connect(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
