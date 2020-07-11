package platformclientv2
import (
	"encoding/json"
)

// Posttextmessage
type Posttextmessage struct { 
	// VarType - Message type
	VarType *string `json:"type,omitempty"`


	// Text - Message text. If type is structured, used as fallback for clients that do not support particular structured content
	Text *string `json:"text,omitempty"`


	// Content - A list of content elements in message
	Content *[]Messagecontent `json:"content,omitempty"`

}

// String returns a JSON representation of the model
func (o *Posttextmessage) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
