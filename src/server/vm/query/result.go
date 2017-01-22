package query

type Result interface {

	MarshalJSON() ([]byte, error)
}
