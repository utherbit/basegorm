# BaseGorm

BaseGorm is a lightweight Golang library designed to simplify the creation of base models and utilize Gorm for database interactions. It provides convenient tools for implementing basic CRUD operations, reducing redundant code.

## Features

- Base model with support for creation date, deletion date, deletion flag, and identifier.
- CRUD repository with five essential functions: Get, Create, Delete, Update, GetList.
- TxOption for maximum flexibility in query execution.

## Installation

```shell
go get -u github.com/utherbit/basegorm
```

## Usage
```go
package main

import (
	"context"
	"fmt"
	"github.com/utherbit/basegorm"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	usersRepo := newUsersRepository(db)

	_, err = usersRepo.AddUser(context.Background(), "utherbit")
	if err != nil {
		panic(err)
	}

	user, err := usersRepo.GetUserByUsername(context.Background(), "utherbit")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%v", user)
}

type User struct {
	basegorm.BaseModel[int] // int - identifier type
	Username                string
}

type usersRepository struct {
	crud *basegorm.BaseCrud[int, *User]
}

func newUsersRepository(db *gorm.DB) *usersRepository {
	return &usersRepository{
		crud: basegorm.NewBaseCrud[int, *User](db, new(User)),
	}
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
func (r *usersRepository) GetUserById(ctx context.Context, id int) (*User, error) {
	option := r.crud.Model.TxOptionById(id)
	user, err := r.crud.Get(ctx, option)
	if err != nil {
		return nil, err
	}
	return user, nil
}
func (r *usersRepository) GetUserByUsername(ctx context.Context, username string) (*User, error) {
	var option basegorm.TxOption = func(tx *gorm.DB) *gorm.DB {
		return tx.Where("Username = ?", username)
	}
	user, err := r.crud.Get(ctx, option)
	if err != nil {
		return nil, err
	}
	return user, nil
}
```

## Examples
Detailed usage examples of the library can be found in the "example" directory of the project.

## Contributions
If you have suggestions for improving the library, feel free to create an Issue or Pull Request. We are open to discussions and contributions.

## License
Copyright (c) 2023-present [Utherbit](https://github.com/utherbit)  and Contributors. BaseGorm is free and open-source software licensed under the MIT License.