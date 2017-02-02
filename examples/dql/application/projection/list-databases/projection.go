package list_databases

import (
	"github.com/domain-query-language/dql-server/src/server/domain/vm/projection"
)

type Projection interface {

	projection.Projection

	/**
		Adds a database name.
	 */

	Add(name string)

	/**
		Removes a database name.
	 */

	Remove(name string)
}
