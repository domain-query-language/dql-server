package event

import "github.com/satori/go.uuid"

var TypeCheckedOut, _ = uuid.FromString("b911900f-3eb7-41e4-8ae5-b5fcced31fb7")

type CheckedOut struct {

}