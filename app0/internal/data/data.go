package data

import (
	"vmapp/app0/internal/data/base"
	"vmapp/app0/internal/data/role_repo"
	"vmapp/app0/internal/data/send_repo"
	"vmapp/app0/internal/data/user_repo"

	"github.com/google/wire"
)

var ProviderData = wire.NewSet(base.NewData, user_repo.NewUserRepo, base.NewTransaction, role_repo.NewRoleRepo, send_repo.NewSendRepo)
