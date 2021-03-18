package platformclientv2
import (
	"time"
	"encoding/json"
	"strconv"
	"strings"
)

// Webchattyping
type Webchattyping struct { 
	// Id - The event identifier of this typing indicator event (useful to guard against event re-delivery
	Id *string `json:"id,omitempty"`


	// Conversation - The identifier of the conversation
	Conversation *Webchatconversation `json:"conversation,omitempty"`


	// Sender - The member who sent the message
	Sender *Webchatmemberinfo `json:"sender,omitempty"`


	// Timestamp - The timestamp of the message, in ISO-8601 format
	Timestamp *time.Time `json:"timestamp,omitempty"`

}

// String returns a JSON representation of the model
func (o *Webchattyping) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
