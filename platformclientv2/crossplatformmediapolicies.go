package platformclientv2
import (
	"github.com/leekchan/timeutil"
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

func (u *Crossplatformmediapolicies) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Crossplatformmediapolicies

	

	return json.Marshal(&struct { 
		CallPolicy *Crossplatformcallmediapolicy `json:"callPolicy,omitempty"`
		
		ChatPolicy *Crossplatformchatmediapolicy `json:"chatPolicy,omitempty"`
		
		EmailPolicy *Crossplatformemailmediapolicy `json:"emailPolicy,omitempty"`
		
		MessagePolicy *Crossplatformmessagemediapolicy `json:"messagePolicy,omitempty"`
		*Alias
	}{ 
		CallPolicy: u.CallPolicy,
		
		ChatPolicy: u.ChatPolicy,
		
		EmailPolicy: u.EmailPolicy,
		
		MessagePolicy: u.MessagePolicy,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Crossplatformmediapolicies) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
