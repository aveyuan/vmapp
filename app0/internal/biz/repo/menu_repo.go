package repo

import (
	"context"
	"vmapp/app0/internal/models"
)

type MenuRepo interface {
	CreateMenu(ctx context.Context, one *models.Menu) error
	UpdateMenu(ctx context.Context, one *models.Menu) error
	DeleteMenu(ctx context.Context, id int32) error
	GetMenu(ctx context.Context, id int32) (one *models.Menu, err error)
	GetMenuByPid(ctx context.Context, pid int32) (all []*models.Menu, err error)
	ListMenu(ctx context.Context) (all []*models.Menu, err error)
	GetMenuByIds(ctx context.Context, ids []string) (all []*models.Menu, err error)
}
