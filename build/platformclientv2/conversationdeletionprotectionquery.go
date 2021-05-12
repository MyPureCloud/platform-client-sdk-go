package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationdeletionprotectionquery
type Conversationdeletionprotectionquery struct { 
	// ConversationIds - This is a list of ConversationIds. The list cannot exceed 100 conversationids.
	ConversationIds *[]string `json:"conversationIds,omitempty"`

}

// String returns a JSON representation of the model
func (o *Conversationdeletionprotectionquery) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
