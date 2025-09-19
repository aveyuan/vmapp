package service

import "github.com/google/wire"

<<<<<<< HEAD
var ProviderService = wire.NewSet( NewGinUserService)
=======
var ProviderService = wire.NewSet( NewGinUserService, NewIrisUserService)
>>>>>>> 946887a5673d2eb48b1be7a4647b1c62b912f534
