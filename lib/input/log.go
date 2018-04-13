package input

import (
	"fmt"
	"log"

	driver "github.com/arangodb/go-driver"
)

func handleLog(coll driver.Collection, msg string) {
	meta, err := coll.CreateDocument(nil, map[string]string{"msg": msg})
	if err != nil {
		fmt.Println("failed to create document: ", err, "\nmeta: ", meta)
	}
	log.Println("Created new document for ", msg)
}
