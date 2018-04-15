package cards

import (
	"fmt"

	"github.com/rubencaro/cardo/lib/db"

	"github.com/rubencaro/cardo/lib/output"
	"github.com/rubencaro/cardo/lib/types"
)

// HandleAddCard will add a card from given data and report back to the socket
func HandleAddCard(data *types.DispatchData) error {
	meta, err := data.Coll.CreateDocument(nil, data.Doc)
	if err != nil {
		return fmt.Errorf("failed to create card: %s\nmeta: %s", err.Error(), meta)
	}
	return reportBack(data, meta.Key)
}

func reportBack(data *types.DispatchData, key string) error {
	card, err := db.ReadCardJSON(data.Coll, key)
	if err != nil {
		return err
	}

	err = output.Send(data.Conn, "cards_upsertCard: "+card)
	if err != nil {
		return fmt.Errorf("failed to send card back: %s", err.Error())
	}

	return nil
}
