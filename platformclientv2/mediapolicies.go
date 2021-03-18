package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Mediapolicies
type Mediapolicies struct { 
	// CallPolicy - Conditions and actions for calls
	CallPolicy *Callmediapolicy `json:"callPolicy,omitempty"`


	// ChatPolicy - Conditions and actions for chats
	ChatPolicy *Chatmediapolicy `json:"chatPolicy,omitempty"`


	// EmailPolicy - Conditions and actions for emails
	EmailPolicy *Emailmediapolicy `json:"emailPolicy,omitempty"`


	// MessagePolicy - Conditions and actions for messages
	MessagePolicy *Messagemediapolicy `json:"messagePolicy,omitempty"`

}

// String returns a JSON representation of the model
func (o *Mediapolicies) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
