package invariant

import (
	"github.com/domain-query-language/dql-server/examples/dql/domain/modelling/database"
	"github.com/satori/go.uuid"
)

var CreatedType = uuid.FromStringOrNil("560c26db-00b1-40ab-8bb8-a729447f637c")

type Created struct {


}

func (self Created) Assumptions(assumptions struct{}) {

}

func (self Created) Satisfier(projection database.Projection) bool {

	return projection.IsCreated == true
}
