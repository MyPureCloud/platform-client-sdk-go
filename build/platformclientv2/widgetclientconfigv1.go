package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Widgetclientconfigv1
type Widgetclientconfigv1 struct { 
	// WebChatSkin
	WebChatSkin *string `json:"webChatSkin,omitempty"`


	// AuthenticationUrl
	AuthenticationUrl *string `json:"authenticationUrl,omitempty"`

}

// String returns a JSON representation of the model
func (o *Widgetclientconfigv1) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
