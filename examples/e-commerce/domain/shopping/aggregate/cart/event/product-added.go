package event

import (
	"github.com/domain-query-language/dql-server/examples/e-commerce/domain/shopping/entity"
	"github.com/satori/go.uuid"
)

var TypeProductAdded, _ = uuid.FromString("30de03ed-12c9-425e-983e-7e9ad8a12aa7")

type ProductAdded struct {

	Product entity.Product
}
