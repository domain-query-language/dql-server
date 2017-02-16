package aggregate

type Invariant interface {

	Assuming(struct {}) Invariant

	Not() Invariant

	IsSatisfied() bool

	Asserts()
}
