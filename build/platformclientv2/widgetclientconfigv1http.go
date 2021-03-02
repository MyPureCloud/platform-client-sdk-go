package platformclientv2
import (
	"encoding/json"
)

// Widgetclientconfigv1http
type Widgetclientconfigv1http struct { 
	// WebChatSkin
	WebChatSkin *string `json:"webChatSkin,omitempty"`


	// AuthenticationUrl
	AuthenticationUrl *string `json:"authenticationUrl,omitempty"`

}

// String returns a JSON representation of the model
func (o *Widgetclientconfigv1http) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
