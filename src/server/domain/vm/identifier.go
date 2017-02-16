package vm

type Identifier interface {

	Bytes() []byte

	MarshalText() (text []byte, err error)
}

type IdentifierGenerator interface {

	Nil() Identifier

	FromBytes(input []byte) (id Identifier, err error)

}
