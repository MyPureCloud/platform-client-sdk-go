package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Crossplatformmediapolicies
type Crossplatformmediapolicies struct { 
	// CallPolicy - Conditions and actions for calls
	CallPolicy *Crossplatformcallmediapolicy `json:"callPolicy,omitempty"`


	// ChatPolicy - Conditions and actions for chats
	ChatPolicy *Crossplatformchatmediapolicy `json:"chatPolicy,omitempty"`


	// EmailPolicy - Conditions and actions for emails
	EmailPolicy *Crossplatformemailmediapolicy `json:"emailPolicy,omitempty"`


	// MessagePolicy - Conditions and actions for messages
	MessagePolicy *Crossplatformmessagemediapolicy `json:"messagePolicy,omitempty"`

}

// String returns a JSON representation of the model
func (o *Crossplatformmediapolicies) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
