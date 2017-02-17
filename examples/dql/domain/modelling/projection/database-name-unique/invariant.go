package database_name_unique

type Invariant struct {

}

var Satisfier = func(queryable Queryable, name string) bool {

	return queryable.Exists(name) == true
}
