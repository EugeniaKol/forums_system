package forums

import "github.com/google/wire"

// Providers for channels components.
var Providers = wire.NewSet(NewStore, HTTPHandler)
