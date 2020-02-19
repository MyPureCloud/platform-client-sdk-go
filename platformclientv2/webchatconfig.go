package platformclientv2
import (
	"encoding/json"
)

// Webchatconfig
type Webchatconfig struct { 
	// WebChatSkin - css class to be applied to the web chat widget.
	WebChatSkin *string `json:"webChatSkin,omitempty"`

}

// String returns a JSON representation of the model
func (o *Webchatconfig) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
