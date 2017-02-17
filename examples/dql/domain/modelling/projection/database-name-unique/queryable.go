package database_name_unique

type Queryable interface {

	/**
		Returns true if the name exists, otherwise false.
	 */
	Exists(name string) bool
}
