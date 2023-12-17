package main

import (
	"basegorm"
	"context"
	"gorm.io/gorm"
)

func newRepository(db *gorm.DB) (*repository, error) {
	var (
		err error
		r   repository
	)
	r.db = db

	r.itemRepo = newItemsRepository(db)
	r.orderRepo = newOrdersRepository(db)
	r.orderItemRepo = newOrderItemsRepository(db)
	r.userRepo = newUsersRepository(db)

	err = r.migration()
	if err != nil {
		return nil, err
	}

	return &r, nil
}

type repository struct {
	db            *gorm.DB
	userRepo      *usersRepository
	orderRepo     *ordersRepository
	itemRepo      *itemsRepository
	orderItemRepo *orderItemsRepository
}

func (r *repository) migration() error {
	return r.db.AutoMigrate(&User{}, &Item{}, &Order{}, &OrderItem{})
}

type usersRepository struct {
	crud *basegorm.BaseCrud[int, *User]
}

func newUsersRepository(db *gorm.DB) *usersRepository {
	return &usersRepository{
		crud: basegorm.NewBaseCrud[int, *User](db, new(User)),
	}
}

type itemsRepository struct {
	crud *basegorm.BaseCrud[int, *Item]
}

func newItemsRepository(db *gorm.DB) *itemsRepository {
	return &itemsRepository{crud: basegorm.NewBaseCrud[int, *Item](db, new(Item))}
}

type ordersRepository struct {
	crud *basegorm.BaseCrud[int, *Order]
}

func newOrdersRepository(db *gorm.DB) *ordersRepository {
	return &ordersRepository{crud: basegorm.NewBaseCrud[int, *Order](db, new(Order))}
}

type orderItemsRepository struct {
	crud *basegorm.BaseCrud[int, *OrderItem]
}

func newOrderItemsRepository(db *gorm.DB) *orderItemsRepository {
	return &orderItemsRepository{crud: basegorm.NewBaseCrud[int, *OrderItem](db, new(OrderItem))}
}

func (r *usersRepository) AddUser(ctx context.Context, username string) (*User, error) {
	user, err := r.crud.Create(ctx, &User{
		Username: username,
	})
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *itemsRepository) AddItem(ctx context.Context, title, description string, cost int) (*Item, error) {
	item, err := r.crud.Create(ctx, &Item{
		Title:       title,
		Description: description,
		Cost:        cost,
	})
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (r *ordersRepository) AddOrder(ctx context.Context, userId int) (*Order, error) {
	order, err := r.crud.Create(ctx, &Order{
		UserId: userId,
	})
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (r *orderItemsRepository) AddOrderItem(ctx context.Context, orderId, itemId, amount int) (*OrderItem, error) {
	orderItem, err := r.crud.Create(ctx, &OrderItem{
		OrderId: orderId,
		ItemId:  itemId,
		Amount:  amount,
	})
	if err != nil {
		return nil, err
	}
	return orderItem, nil
}

func (r *ordersRepository) GetOrder(ctx context.Context, orderId int, joinItems bool) (*Order, error) {

	var option = r.crud.Model.TxOptionById(orderId)

	if joinItems {
		option = option.Apply(r.crud.Model.TxOptionPreloadNotDelete("Items.Item"))
	}

	order, err := r.crud.Get(ctx, option)
	if err != nil {
		return nil, err
	}
	return order, nil
}
func (r *ordersRepository) GetOrderByUser(ctx context.Context, userId int, joinItems bool) (*Order, error) {
	var option = r.crud.Model.TxOptionWhere("UserId = ?", userId)
	if joinItems {
		option = option.Apply(r.crud.Model.TxOptionPreloadNotDelete("Items.Item"))
	}

	order, err := r.crud.Get(ctx, option)
	if err != nil {
		return nil, err
	}
	return order, nil
}
