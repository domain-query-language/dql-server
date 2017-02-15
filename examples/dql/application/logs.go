package application

import "github.com/domain-query-language/dql-server/src/server/infrastructure/domain/store"

var EventLog = store.NewMemoryLog()

var CommandLog = store.NewMemoryLog()
