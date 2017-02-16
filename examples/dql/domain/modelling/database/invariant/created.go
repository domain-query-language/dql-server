package invariant

import (
	"github.com/domain-query-language/dql-server/examples/dql/domain/modelling/database"
	"github.com/satori/go.uuid"
)

var Created = uuid.FromStringOrNil("560c26db-00b1-40ab-8bb8-a729447f637c")

type CreatedInvariant struct {


}

func (self CreatedInvariant) Assumptions(assumptions struct{}) {

}

func (self CreatedInvariant) Satisfier(projection database.Projection) bool {

	return projection.IsCreated == true
}
