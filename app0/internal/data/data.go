package data

import (
	"vmapp/app0/internal/data/base"
	"vmapp/app0/internal/data/user_repo"

	"github.com/google/wire"
)

var ProviderData = wire.NewSet(base.NewData, user_repo.NewUserRepo, base.NewTransaction)
