package basegorm

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type IIdentifier interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~[16]byte /* uuid */ |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 | ~string
}

var _ IBaseModel[int] = (*BaseModel[int])(nil)
var _ IBaseModel[uuid.UUID] = (*BaseModel[uuid.UUID])(nil)

type IBaseModel[ID IIdentifier] interface {
	SetId(v ID)
	GetId() ID
	SetIsDelete(v bool)
	GetIsDelete() bool
	SetDateCreate(v time.Time)
	GetDateCreate() time.Time
	SetDateUpdate(v time.Time)
	GetDateUpdate() time.Time
}

type BaseModel[ID IIdentifier] struct {
	ID         ID        `db:"id"`
	DateCreate time.Time `db:"date_create"`
	DateUpdate time.Time `db:"date_update"`
	IsDelete   *bool     `db:"is_delete"`
}

func (b *BaseModel[ID]) SetId(id ID) {
	b.ID = id
}

func (b *BaseModel[ID]) GetId() ID {
	return b.ID
}

func (b *BaseModel[ID]) SetDateCreate(v time.Time) {
	b.DateCreate = v
}
func (b *BaseModel[ID]) GetDateCreate() time.Time {
	return b.DateCreate
}

func (b *BaseModel[ID]) SetDateUpdate(v time.Time) {
	b.DateUpdate = v
}
func (b *BaseModel[ID]) GetDateUpdate() time.Time {
	return b.DateUpdate
}

func (b *BaseModel[ID]) SetIsDelete(v bool) {
	b.IsDelete = &v
}
func (b *BaseModel[ID]) GetIsDelete() bool {
	if b.IsDelete == nil {
		return false
	}
	return *b.IsDelete
}

func (_ *BaseModel[ID]) TxOptionById(id ID) TxOption {
	return func(tx *gorm.DB) *gorm.DB {
		return tx.Where("id = ?", id)
	}
}

func (_ *BaseModel[ID]) TxOptionPreloadNotDelete(filed string) TxOption {
	return func(tx *gorm.DB) *gorm.DB {
		return tx.Preload(filed, "is_delete = false")
	}
}
func (_ *BaseModel[ID]) TxOptionNotDelete() TxOption {
	return func(tx *gorm.DB) *gorm.DB {
		return tx.Where("is_delete = false")
	}
}
func (_ *BaseModel[ID]) TxOptionWhere(query interface{}, args ...interface{}) TxOption {
	return func(tx *gorm.DB) *gorm.DB {
		return tx.Where(query, args...)
	}
}
