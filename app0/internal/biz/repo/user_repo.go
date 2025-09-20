package repo

import (
	"context"
	"vmapp/app0/internal/dto"
	"vmapp/app0/internal/models"
)

type UserKey string

const (
	Username     UserKey = "username"
	Phone        UserKey = "phone"
	Email        UserKey = "email"
	TokenUserkey UserKey = "token"
)

type UserRepo interface {
	CreateUser(ctx context.Context, one *models.User) error
	ListUser(ctx context.Context, req *dto.ListUserReq) (count int64, all []*models.User, err error)
	UpdateUser(ctx context.Context, one *models.User) error
	DeleteUser(ctx context.Context, uid int64) error
	RepassUser(ctx context.Context, uid int64, password string, salt string) error
	GetUser(ctx context.Context, uid int64) (one *models.User, err error)
	GetUserByKey(ctx context.Context, key UserKey, value interface{}) (one *models.User, err error)
	UpdateUserKey(ctx context.Context, uid int64, key UserKey, value interface{}) (err error)
	UPdateUserState(ctx context.Context, uid int64, state int8) error
	GetAllRoleCodes(ctx context.Context) (all []string, err error)
}
