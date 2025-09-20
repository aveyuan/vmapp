package role_repo

import (
	"context"
	"vmapp/app0/internal/biz/repo"
	"vmapp/app0/internal/conf"
	"vmapp/app0/internal/models"

	"vmapp/app0/internal/data/base"

	"github.com/go-kratos/kratos/v2/log"
)

type roleRepo struct {
	data *base.Data
	log  *log.Helper
}

func NewRoleRepo(data *base.Data, component *conf.BootComponent) repo.RoleRepo {
	return &roleRepo{
		data: data,
		log:  component.Logger,
	}
}

func (t *roleRepo) CreateRole(ctx context.Context, one *models.Role) error {
	return t.data.DB(ctx).Create(&one).Error
}

// UpdateRole rolecode不能被更新
func (t *roleRepo) UpdateRole(ctx context.Context, one *models.Role) error {
	return t.data.DB(ctx).Model(new(models.Role)).Select("*").Omit("role_code").Where("id = ?", one.Id).UpdateColumns(one).Error
}

func (t *roleRepo) DeleteRole(ctx context.Context, id int32) error {
	return t.data.DB(ctx).Delete(new(models.Role), "id = ?", id).Error
}

func (t *roleRepo) GetRole(ctx context.Context, id int32) (one *models.Role, err error) {
	err = t.data.DB(ctx).First(&one, "id = ?", id).Error
	return
}

func (t *roleRepo) ListRole(ctx context.Context) (all []*models.Role, err error) {
	err = t.data.DB(ctx).Model(new(models.Role)).Order("id desc").Find(&all).Error
	return
}

func (t *roleRepo) GetRoleByCodes(ctx context.Context, codes []string) (all []*models.Role, err error) {
	err = t.data.DB(ctx).Find(&all, "role_code in (?)", codes).Error
	return
}

func (t *roleRepo) GetAllMenuIDS(ctx context.Context) (all []string, err error) {
	err = t.data.DB(ctx).Model(new(models.Role)).Select("menu_ids").Scan(&all).Error
	return
}
