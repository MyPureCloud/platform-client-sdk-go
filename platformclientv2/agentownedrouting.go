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

func (o *Agentownedrouting) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Agentownedrouting
	
	return json.Marshal(&struct { 
		EnableAgentOwnedCallbacks *bool `json:"enableAgentOwnedCallbacks,omitempty"`
		
		MaxOwnedCallbackHours *int `json:"maxOwnedCallbackHours,omitempty"`
		
		MaxOwnedCallbackDelayHours *int `json:"maxOwnedCallbackDelayHours,omitempty"`
		*Alias
	}{ 
		EnableAgentOwnedCallbacks: o.EnableAgentOwnedCallbacks,
		
		MaxOwnedCallbackHours: o.MaxOwnedCallbackHours,
		
		MaxOwnedCallbackDelayHours: o.MaxOwnedCallbackDelayHours,
		Alias:    (*Alias)(o),
	})
}

func (o *Agentownedrouting) UnmarshalJSON(b []byte) error {
	var AgentownedroutingMap map[string]interface{}
	err := json.Unmarshal(b, &AgentownedroutingMap)
	if err != nil {
		return err
	}
	
	if EnableAgentOwnedCallbacks, ok := AgentownedroutingMap["enableAgentOwnedCallbacks"].(bool); ok {
		o.EnableAgentOwnedCallbacks = &EnableAgentOwnedCallbacks
	}
    
	if MaxOwnedCallbackHours, ok := AgentownedroutingMap["maxOwnedCallbackHours"].(float64); ok {
		MaxOwnedCallbackHoursInt := int(MaxOwnedCallbackHours)
		o.MaxOwnedCallbackHours = &MaxOwnedCallbackHoursInt
	}
	
	if MaxOwnedCallbackDelayHours, ok := AgentownedroutingMap["maxOwnedCallbackDelayHours"].(float64); ok {
		MaxOwnedCallbackDelayHoursInt := int(MaxOwnedCallbackDelayHours)
		o.MaxOwnedCallbackDelayHours = &MaxOwnedCallbackDelayHoursInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Agentownedrouting) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
