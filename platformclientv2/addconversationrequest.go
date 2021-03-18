package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Addconversationrequest - Update coaching appointment request
type Addconversationrequest struct { 
	// ConversationId - The id of the conversation to add
	ConversationId *string `json:"conversationId,omitempty"`

}

// String returns a JSON representation of the model
func (o *Addconversationrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
