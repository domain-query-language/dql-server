package application

import "github.com/domain-query-language/dql-server/src/server/domain/vm/handler/query"

var QueryHandler = query.CreateHandler(ProjectionsRepository)
