package repo

import (
	"context"
	"vmapp/app0/internal/models"
)

type RoleRepo interface {
	CreateRole(ctx context.Context, one *models.Role) error
	UpdateRole(ctx context.Context, one *models.Role) error
	DeleteRole(ctx context.Context, id int32) error
	GetRole(ctx context.Context, id int32) (one *models.Role, err error)
	ListRole(ctx context.Context) (all []*models.Role, err error)
	GetRoleByCodes(ctx context.Context, codes []string) (all []*models.Role, err error)
	GetAllMenuIDS(ctx context.Context) (all []string, err error)
}
