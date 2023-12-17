package basegorm

import (
	"gorm.io/gorm"
	"testing"
)

type ITxOptionTest interface {
	Apply(tx *gorm.DB) *gorm.DB
}

func TestName(t *testing.T) {

}
