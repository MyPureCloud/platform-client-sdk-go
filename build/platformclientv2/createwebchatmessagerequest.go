package platformclientv2
import (
	"encoding/json"
)

// Createwebchatmessagerequest
type Createwebchatmessagerequest struct { 
	// Body - The message body. Note that message bodies are limited to 4,000 characters.
	Body *string `json:"body,omitempty"`


	// BodyType - The purpose of the message within the conversation, such as a standard text entry versus a greeting.
	BodyType *string `json:"bodyType,omitempty"`

}

// String returns a JSON representation of the model
func (o *Createwebchatmessagerequest) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
