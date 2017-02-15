package list_databases

import (
	"github.com/domain-query-language/dql-server/src/server/domain/vm/projection"
	"github.com/domain-query-language/dql-server/src/server/domain/vm"
)

type Projection interface {

	projection.Projection

	/**
		Adds a database name.
	 */

	Add(id vm.Identifier, name string)

	/**
		Renames a database.
 	*/

	Rename(id vm.Identifier, name string)

	/**
		Removes a database name.
	 */

	Remove(id vm.Identifier)
}
