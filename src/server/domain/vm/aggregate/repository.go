package aggregate

type Repository interface {

	/**

		Returns an Aggregate by its Type Identifier.

	 */

	Get(id Identifier) (Aggregate, error)

	/**

		Saves an Aggregate.

	 */

	Save(aggregate Aggregate) error

}
