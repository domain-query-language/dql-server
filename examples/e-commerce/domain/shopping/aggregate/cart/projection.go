package cart

import (
	"github.com/domain-query-language/dql-server/src/server/domain/vm"
	"github.com/domain-query-language/dql-server/examples/e-commerce/domain/shopping/entity"
	"errors"
	"github.com/domain-query-language/dql-server/examples/e-commerce/domain/shopping/value"
)

type Projection struct {

	Id vm.Identifier

	IsCreated bool
	IsCheckedOut bool

	Cart entity.Cart
	Products map[vm.Identifier] entity.Product
}

func (self *Projection) Id() {
	return self.Id
}

func (self *Projection) Reset() {

	self.IsCreated = false
	self.IsCheckedOut = false
}

func (self *Projection) Create() {
	self.IsCreated = true
}

func(self *Projection) AddProduct(product entity.Product) error {

	_, exists := self.Products[product.Id]

	if exists {
		return errors.New("Product already exists.")
	}

	self.Products[product.Id] = product

	return nil
}

func (self *Projection) ChangeProductQuantity(id vm.Identifier, quantity value.Quantity) error {

	product, exists := self.Products[id]

	if exists {
		return errors.New("Product already exists.")
	}

	product.Quantity = quantity

	self.Products[id] = product

	return nil
}

func (self *Projection) RemoveProduct(id vm.Identifier) error {

	_, exists := self.Products[id]

	if !exists {
		return errors.New("Product does not exist.")
	}

	delete(self.Products, id)

	return nil
}

func (self *Projection) Checkout() {
	self.IsCheckedOut = true
}

func NewProjection(id vm.Identifier) *Projection {

	projection := &Projection {
		Id: id,
	}

	projection.Reset()

	return projection
}
