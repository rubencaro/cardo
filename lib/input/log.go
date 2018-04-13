package input

import (
	"fmt"
)

func init() {
	routes["log"] = handleLog
}

func handleLog(data *dispatchData) error {
	meta, err := data.coll.CreateDocument(nil, map[string]string{"msg": data.msg})
	if err != nil {
		return fmt.Errorf("failed to create document: %s\nmeta: %s", err.Error(), meta)
	}
	fmt.Println("Created new document for ", data.msg)
	return nil
}
