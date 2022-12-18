package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Directroutingcallmediasettings
type Directroutingcallmediasettings struct { 
	// Enabled - Toggle that enables Direct Routing for this media type.
	Enabled *bool `json:"enabled,omitempty"`


	// InboundFlow - The Direct Routing inbound flow id for this media type.
	InboundFlow *Addressableentityref `json:"inboundFlow,omitempty"`


	// VoicemailFlow - ID of the in-queue flow that handles voicemails for Direct Routing and to transfer calls to ACD voicemail.
	VoicemailFlow *Addressableentityref `json:"voicemailFlow,omitempty"`

}

func (o *Directroutingcallmediasettings) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Directroutingcallmediasettings
	
	return json.Marshal(&struct { 
		Enabled *bool `json:"enabled,omitempty"`
		
		InboundFlow *Addressableentityref `json:"inboundFlow,omitempty"`
		
		VoicemailFlow *Addressableentityref `json:"voicemailFlow,omitempty"`
		*Alias
	}{ 
		Enabled: o.Enabled,
		
		InboundFlow: o.InboundFlow,
		
		VoicemailFlow: o.VoicemailFlow,
		Alias:    (*Alias)(o),
	})
}

func (o *Directroutingcallmediasettings) UnmarshalJSON(b []byte) error {
	var DirectroutingcallmediasettingsMap map[string]interface{}
	err := json.Unmarshal(b, &DirectroutingcallmediasettingsMap)
	if err != nil {
		return err
	}
	
	if Enabled, ok := DirectroutingcallmediasettingsMap["enabled"].(bool); ok {
		o.Enabled = &Enabled
	}
    
	if InboundFlow, ok := DirectroutingcallmediasettingsMap["inboundFlow"].(map[string]interface{}); ok {
		InboundFlowString, _ := json.Marshal(InboundFlow)
		json.Unmarshal(InboundFlowString, &o.InboundFlow)
	}
	
	if VoicemailFlow, ok := DirectroutingcallmediasettingsMap["voicemailFlow"].(map[string]interface{}); ok {
		VoicemailFlowString, _ := json.Marshal(VoicemailFlow)
		json.Unmarshal(VoicemailFlowString, &o.VoicemailFlow)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Directroutingcallmediasettings) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
