package basegorm

import (
	"fmt"
	"gorm.io/gorm"
	"testing"
)

func TestOption(t *testing.T) {
	var testOptionRowAffectAddOne TxOption = func(tx *gorm.DB) *gorm.DB {
		tx.RowsAffected += 1
		return tx
	}

	{
		tx := new(gorm.DB)
		tx.RowsAffected = 0

		tx = testOptionRowAffectAddOne(tx)

		if tx.RowsAffected != 1 {
			t.Errorf("Expected RowsAffected to be 1, got %d", tx.RowsAffected)
		}
	}

	{
		tx := new(gorm.DB)
		tx.RowsAffected = 0

		option := testOptionRowAffectAddOne
		option = option.Apply(testOptionRowAffectAddOne)
		option = option.Apply(testOptionRowAffectAddOne)

		tx = option(tx)

		if tx.RowsAffected != 3 {
			t.Errorf("Expected RowsAffected to be 3, got %d", tx.RowsAffected)
		}
	}
}

func TestOptionOrder(t *testing.T) {
	// Проверяет, что предшествующее значение в RowsAffected является before, после выполнения задаёт значение after
	var newCheckOrderOption = func(before int64, after int64) TxOption {
		return func(tx *gorm.DB) *gorm.DB {
			if tx.RowsAffected != before {
				tx.Error = fmt.Errorf("expected RowsAffected to be %d, got %d", before, tx.RowsAffected)
			}
			tx.RowsAffected = after
			return tx
		}
	}

	// Проверяет что функция newCheckOrderOption работает как ожидалось
	{
		tx := new(gorm.DB)
		tx.RowsAffected = 0

		option := newCheckOrderOption(0, 1)

		tx = option(tx)
		if tx.RowsAffected != 1 {
			t.Errorf("Expected RowsAffected to be 1, got %d", tx.RowsAffected)
		}
		if tx.Error != nil {
			t.Error(tx.Error)
		}
	}

	{
		tx := new(gorm.DB)
		tx.RowsAffected = 0

		option := newCheckOrderOption(0, 1)
		option = option.Apply(newCheckOrderOption(1, 2))
		option = option.Apply(newCheckOrderOption(2, 3))

		tx = option(tx)

		if tx.Error != nil {
			t.Error(tx.Error)
		}
	}

}
