package schema

import (
	"github.com/domain-query-language/dql-server/src/server/domain/vm"
)

type QueryID string;

func (q QueryID) Bytes() []byte {

	return []byte(q);
}

/**
 * Query IDs
 */
const (
	LIST_DATABASES QueryID = "aa31cf30-f33e-4267-ade6-ee914d907c8c"
)


/**
 * Queries
 */
type ListDatabases struct {

}

func (q *ListDatabases) String () string {

	return "ListDatabases";
}

func (q *ListDatabases) Id () vm.Identifier {

	return LIST_DATABASES;
}
