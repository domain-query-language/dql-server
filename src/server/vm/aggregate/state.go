package aggregate

type State interface {

	GobEncode() []byte

	GobDecode([]byte) error
}
