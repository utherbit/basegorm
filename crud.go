package basegorm

import (
	"context"
	"gorm.io/gorm"
	"time"
)

func NewBaseCrud[ID IIdentifier, MODEL IBaseModel[ID]](db *gorm.DB, object MODEL) *BaseCrud[ID, MODEL] {
	return &BaseCrud[ID, MODEL]{
		Model: object,
		DB:    db,
	}

}

type BaseCrud[ID IIdentifier, MODEL IBaseModel[ID]] struct {
	Model MODEL
	DB    *gorm.DB
}

func (b *BaseCrud[ID, MODEL]) Get(ctx context.Context, option TxOption) (MODEL, error) {
	b.DB.WithContext(ctx)
	var model MODEL
	tx := b.DB.Model(b.Model)

	tx = option(tx)

	tx = tx.First(&model)
	if tx.Error != nil {
		return model, tx.Error
	}

	return model, nil
}

func (b *BaseCrud[ID, MODEL]) Create(ctx context.Context, value MODEL) (MODEL, error) {
	b.DB.WithContext(ctx)
	var model MODEL

	// Применяю значения по умолчанию для новой записи
	var now = time.Now()
	value.SetDateCreate(now)
	value.SetDateUpdate(now)
	value.SetIsDelete(false)

	tx := b.DB.Create(value)

	tx.Scan(&model)
	if tx.Error != nil {
		return model, tx.Error
	}

	return model, nil
}

// Delete Удалить модель по параметрам транзакции
func (b *BaseCrud[ID, MODEL]) Delete(ctx context.Context, option TxOption) (MODEL, error) {
	b.DB.WithContext(ctx)
	var model MODEL
	tx := b.DB.Model(model)

	// Применяю опции запроса
	tx = option(tx)

	tx = tx.First(model)
	if tx.Error != nil {
		return model, tx.Error
	}

	model.SetIsDelete(true)
	tx = b.DB.Updates(model)
	if tx.Error != nil {
		return model, tx.Error
	}

	return model, nil
}

func (b *BaseCrud[ID, MODEL]) Update(ctx context.Context, value MODEL) (MODEL, error) {
	b.DB.WithContext(ctx)
	value.SetDateUpdate(time.Now())
	tx := b.DB.Updates(value)

	if tx.Error != nil {
		return value, tx.Error
	}

	return value, nil
}

func (b *BaseCrud[ID, MODEL]) GetList(ctx context.Context, options ...TxOption) ([]MODEL, error) {
	b.DB.WithContext(ctx)
	var model MODEL
	tx := b.DB.Model(model)

	for _, option := range options {
		tx = option(tx)
	}

	var models = make([]MODEL, 0)
	tx = tx.Scan(models)
	if tx.Error != nil {
		return models, tx.Error
	}

	return models, nil
}
