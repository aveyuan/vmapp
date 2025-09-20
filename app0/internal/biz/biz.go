package biz

import (
	"vmapp/app0/internal/biz/usecase"

	"github.com/google/wire"
)

// ProviderSet is biz providers.
var ProviderBiz = wire.NewSet(usecase.NewUserUseCase)
