package server

import "github.com/google/wire"

var ProviderServer = wire.NewSet( NewGf, NewGin, NewIris)
