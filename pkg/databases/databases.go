package databases

import "github.com/google/wire"

var DatabaseSet = wire.NewSet(NewMySQLDB, NewFirebaseClient, NewFirestoreDB)
