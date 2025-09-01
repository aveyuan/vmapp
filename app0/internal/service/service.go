package service

import "github.com/google/wire"

var ProviderService = wire.NewSet(NewGfUserService, NewGinUserService, NewIrisUserService)
