package main

import (
	"github.com/domain-query-language/dql-server/src/server/domain/vm/player"
	"github.com/satori/go.uuid"
	"github.com/davecgh/go-spew/spew"
	"bytes"
	"encoding/gob"
	"github.com/domain-query-language/dql-server/src/server/domain/vm"
)

func main() {

	snapshot := player.NewSnapshot(
		uuid.NewV4(),
		uuid.NewV4(),
	)

	var encoded bytes.Buffer

	gob.Register(vm.Identifier(uuid.UUID{}))

	encoder := gob.NewEncoder(&encoded)

	encoder.Encode(snapshot)

	decoder := gob.NewDecoder(&encoded)

	var new_snapshot player.Snapshot

	decoder.Decode(&new_snapshot)

	spew.Dump(new_snapshot)


}
