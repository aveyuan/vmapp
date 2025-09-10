package repo

import (
	"context"

	"gorm.io/gorm"
)

type Transaction interface {
	InTx(context.Context, func(ctx context.Context) error) error
	Tx(context.Context) (*gorm.DB, context.Context)
}
