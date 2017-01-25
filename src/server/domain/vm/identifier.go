package vm

type Identifier interface {

	Bytes() []byte
}

type IdentifierGenerator interface {

	Nil() Identifier

	FromBytes(input []byte) (id Identifier, err error)

}
