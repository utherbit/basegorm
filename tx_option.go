package basegorm

import (
	"gorm.io/gorm"
)

type TxOption func(tx *gorm.DB) *gorm.DB

func (option TxOption) Apply(apply TxOption) TxOption {
	return func(tx *gorm.DB) *gorm.DB {
		return apply(option(tx))
	}
}

//type TxOptions struct {
//	options []TxOption
//}

//func NewTxOptions() *TxOptions {
//	return &TxOptions{
//		options: make([]TxOption, 0),
//	}
//}
//func (op *TxOptions) Add(option TxOption) {
//	op.options = append(op.options, option)
//}
//func (op *TxOptions) apply(tx *gorm.DB) *gorm.DB {
//	for _, option := range op.options {
//		tx = option(tx)
//	}
//	return tx
//}
