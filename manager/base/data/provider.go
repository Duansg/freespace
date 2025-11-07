package data

import "github.com/google/wire"

// ProviderSetData Data layer provider collection
var ProviderSetData = wire.NewSet(
	NewData,
	NewDB,
)
