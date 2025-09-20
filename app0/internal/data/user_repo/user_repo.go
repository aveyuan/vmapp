package user_repo

import (
	"context"
	"fmt"
	"vmapp/app0/internal/biz/repo"
	"vmapp/app0/internal/conf"
	"vmapp/app0/internal/dto"
	"vmapp/app0/internal/models"

	"vmapp/app0/internal/data/base"

	"github.com/go-kratos/kratos/v2/log"
)

type userRepo struct {
	data *base.Data
	log  *log.Helper
}

func NewUserRepo(data *base.Data, component *conf.BootComponent) repo.UserRepo {
	return &userRepo{
		data: data,
		log:  component.Logger,
	}
}

func (t *userRepo) CreateUser(ctx context.Context, one *models.User) error {
	err := t.data.DB(ctx).Create(&one).Error
	return err
}

func (t *userRepo) DeleteUser(ctx context.Context, uid int64) error {
	return t.data.DB(ctx).Delete(new(models.User), "uid = ?", uid).Error
}

func (t *userRepo) ListUser(ctx context.Context, req *dto.ListUserReq) (count int64, all []*models.User, err error) {
	if req.Page == 0 {
		req.Page = 1
	}

	if req.Limit == 0 {
		req.Limit = 10
	}

	q := t.data.DB(ctx).Model(new(models.User))

	if req.Mail != "" {
		q.Where("mail like ?", "%"+req.Mail+"%")
	}
	if req.Nickname != "" {
		q.Where("nick_name like ?", "%"+req.Nickname+"%")
	}
	if req.Username != "" {
		q.Where("username like ?", "%"+req.Username+"%")
	}

	err = q.Count(&count).Limit(req.Limit).Offset((req.Page - 1) * req.Limit).Order("uid desc").Find(&all).Error
	return
}

func (t *userRepo) GetUser(ctx context.Context, uid int64) (one *models.User, err error) {
	err = t.data.DB(ctx).First(&one, "uid = ?", uid).Error
	return
}

func (t *userRepo) GetUserByUserName(ctx context.Context, username string) (one *models.User, err error) {
	err = t.data.DB(ctx).First(&one, "username = ?", username).Error
	return
}

func (t *userRepo) UpdateUser(ctx context.Context, one *models.User) error {
	return t.data.DB(ctx).Model(new(models.User)).Select("*").Where("uid = ?", one.Uid).UpdateColumns(&one).Error
}

func (t *userRepo) UPdateUserState(ctx context.Context, uid int64, state int8) error {
	return t.data.DB(ctx).Model(new(models.User)).Where("uid = ?", uid).Update("state", state).Error
}

func (t *userRepo) RepassUser(ctx context.Context, uid int64, password string, salt string) error {
	return t.data.DB(ctx).Model(new(models.User)).Where("uid = ?", uid).UpdateColumns(map[string]interface{}{
		"password": password,
		"salt":     salt,
	}).Error
}

func (t *userRepo) UpdateUserKey(ctx context.Context, uid int64, key repo.UserKey, value interface{}) (err error) {
	return t.data.DB(ctx).Model(new(models.User)).Where("uid = ?", uid).Update(string(key), value).Error
}

func (t *userRepo) GetUserByKey(ctx context.Context, key repo.UserKey, value interface{}) (one *models.User, err error) {
	err = t.data.DB(ctx).First(&one, fmt.Sprintf("%v = ?", key), value).Error
	return
}

func (t *userRepo) GetAllRoleCodes(ctx context.Context) (all []string, err error) {
	err = t.data.DB(ctx).Model(new(models.User)).Select("role_codes").Scan(&all).Error
	return
}
