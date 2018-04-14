package cards

import (
	"fmt"

	"github.com/rubencaro/cardo/lib/output"
	"github.com/rubencaro/cardo/lib/types"
)

// HandleAddCard will add a card from given data and report back to the socket
func HandleAddCard(data *types.DispatchData) error {
	meta, err := data.Coll.CreateDocument(nil, data.Doc)
	if err != nil {
		return fmt.Errorf("failed to create card: %s\nmeta: %s", err.Error(), meta)
	}
	// fmt.Println("Created new card for ", data.payload)
	output.Send(data.Conn, "Card created")
	return nil
}
