package list_databases

type Queryable interface {

	/**
		Returns an alphabetically ordered list of database names.
	 */
	List() []string
}
