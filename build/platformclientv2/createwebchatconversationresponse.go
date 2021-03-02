package platformclientv2
import (
	"encoding/json"
)

// Createwebchatconversationresponse
type Createwebchatconversationresponse struct { 
	// Id - Chat Conversation identifier
	Id *string `json:"id,omitempty"`


	// Jwt - The JWT that you can use to identify subsequent calls on this conversation
	Jwt *string `json:"jwt,omitempty"`


	// EventStreamUri - The URI which provides the conversation event stream.
	EventStreamUri *string `json:"eventStreamUri,omitempty"`


	// Member - Chat Member
	Member *Webchatmemberinfo `json:"member,omitempty"`

}

// String returns a JSON representation of the model
func (o *Createwebchatconversationresponse) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
