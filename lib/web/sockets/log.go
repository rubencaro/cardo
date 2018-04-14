package sockets

import (
	"fmt"

	"github.com/rubencaro/cardo/lib/types"
)

func init() {
	routes["log"] = handleLog
}

func handleLog(data *types.DispatchData) error {
	meta, err := data.Coll.CreateDocument(nil, map[string]string{"msg": data.Payload})
	if err != nil {
		return fmt.Errorf("failed to create log: %s\nmeta: %s", err.Error(), meta)
	}
	// fmt.Println("Created new log for ", data.Payload)
	return nil
}
