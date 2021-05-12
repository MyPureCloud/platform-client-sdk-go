package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
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
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
