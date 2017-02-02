package event

import "github.com/satori/go.uuid"

var TypeEmpty, _ = uuid.FromString("4d510326-9a6d-4475-81db-8651ad7f2008")

type Empty struct {}
