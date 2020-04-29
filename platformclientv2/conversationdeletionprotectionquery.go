package platformclientv2
import (
	"encoding/json"
)

// Conversationdeletionprotectionquery
type Conversationdeletionprotectionquery struct { 
	// ConversationIds - This is a list of ConversationIds. The list cannot exceed 100 conversationids.
	ConversationIds *[]string `json:"conversationIds,omitempty"`

}

// String returns a JSON representation of the model
func (o *Conversationdeletionprotectionquery) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
