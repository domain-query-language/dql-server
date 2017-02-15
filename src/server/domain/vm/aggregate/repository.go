package aggregate

import "github.com/domain-query-language/dql-server/src/server/domain/vm"

type Repository interface {

	/**

		Returns an Aggregate by its Type Identifier.

	 */

	Get(id *vm.AggregateIdentifier) (Aggregate, error)

	/**

		Saves an Aggregate.

	 */

	Save(aggregate Aggregate) error

}
