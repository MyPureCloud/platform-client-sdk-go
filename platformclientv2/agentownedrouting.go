package platformclientv2
import (
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

// String returns a JSON representation of the model
func (o *Agentownedrouting) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
