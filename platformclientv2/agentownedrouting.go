package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Agentownedrouting
type Agentownedrouting struct { 
	// EnableAgentOwnedCallbacks - Indicates if Agent Owned Callbacks are enabled for the queue
	EnableAgentOwnedCallbacks *bool `json:"enableAgentOwnedCallbacks,omitempty"`


	// MaxOwnedCallbackHours - The max amount of time a callback can be owned (in hours); Allowable range 1 - 168 hour(s) (inclusive)
	MaxOwnedCallbackHours *int `json:"maxOwnedCallbackHours,omitempty"`


	// MaxOwnedCallbackDelayHours - The max amount of time a callback can be scheduled out into the future (in hours); Allowable range 1 - 720 hour(s) (inclusive)
	MaxOwnedCallbackDelayHours *int `json:"maxOwnedCallbackDelayHours,omitempty"`

}

func (u *Agentownedrouting) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Agentownedrouting

	

	return json.Marshal(&struct { 
		EnableAgentOwnedCallbacks *bool `json:"enableAgentOwnedCallbacks,omitempty"`
		
		MaxOwnedCallbackHours *int `json:"maxOwnedCallbackHours,omitempty"`
		
		MaxOwnedCallbackDelayHours *int `json:"maxOwnedCallbackDelayHours,omitempty"`
		*Alias
	}{ 
		EnableAgentOwnedCallbacks: u.EnableAgentOwnedCallbacks,
		
		MaxOwnedCallbackHours: u.MaxOwnedCallbackHours,
		
		MaxOwnedCallbackDelayHours: u.MaxOwnedCallbackDelayHours,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Agentownedrouting) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
