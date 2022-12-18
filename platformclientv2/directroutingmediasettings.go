package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Directroutingmediasettings
type Directroutingmediasettings struct { 
	// Enabled - Toggle that enables Direct Routing for this media type.
	Enabled *bool `json:"enabled,omitempty"`


	// InboundFlow - The Direct Routing inbound flow id for this media type.
	InboundFlow *Addressableentityref `json:"inboundFlow,omitempty"`

}

func (o *Directroutingmediasettings) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Directroutingmediasettings
	
	return json.Marshal(&struct { 
		Enabled *bool `json:"enabled,omitempty"`
		
		InboundFlow *Addressableentityref `json:"inboundFlow,omitempty"`
		*Alias
	}{ 
		Enabled: o.Enabled,
		
		InboundFlow: o.InboundFlow,
		Alias:    (*Alias)(o),
	})
}

func (o *Directroutingmediasettings) UnmarshalJSON(b []byte) error {
	var DirectroutingmediasettingsMap map[string]interface{}
	err := json.Unmarshal(b, &DirectroutingmediasettingsMap)
	if err != nil {
		return err
	}
	
	if Enabled, ok := DirectroutingmediasettingsMap["enabled"].(bool); ok {
		o.Enabled = &Enabled
	}
    
	if InboundFlow, ok := DirectroutingmediasettingsMap["inboundFlow"].(map[string]interface{}); ok {
		InboundFlowString, _ := json.Marshal(InboundFlow)
		json.Unmarshal(InboundFlowString, &o.InboundFlow)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Directroutingmediasettings) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
