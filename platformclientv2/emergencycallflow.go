package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Emergencycallflow - An emergency flow associates a call flow to use in an emergency with the ivr(s) to route to it.
type Emergencycallflow struct { 
	// EmergencyFlow - The call flow to execute in an emergency.
	EmergencyFlow *Domainentityref `json:"emergencyFlow,omitempty"`


	// Ivrs - The IVR(s) to route to the call flow during an emergency.
	Ivrs *[]Domainentityref `json:"ivrs,omitempty"`

}

func (o *Emergencycallflow) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Emergencycallflow
	
	return json.Marshal(&struct { 
		EmergencyFlow *Domainentityref `json:"emergencyFlow,omitempty"`
		
		Ivrs *[]Domainentityref `json:"ivrs,omitempty"`
		*Alias
	}{ 
		EmergencyFlow: o.EmergencyFlow,
		
		Ivrs: o.Ivrs,
		Alias:    (*Alias)(o),
	})
}

func (o *Emergencycallflow) UnmarshalJSON(b []byte) error {
	var EmergencycallflowMap map[string]interface{}
	err := json.Unmarshal(b, &EmergencycallflowMap)
	if err != nil {
		return err
	}
	
	if EmergencyFlow, ok := EmergencycallflowMap["emergencyFlow"].(map[string]interface{}); ok {
		EmergencyFlowString, _ := json.Marshal(EmergencyFlow)
		json.Unmarshal(EmergencyFlowString, &o.EmergencyFlow)
	}
	
	if Ivrs, ok := EmergencycallflowMap["ivrs"].([]interface{}); ok {
		IvrsString, _ := json.Marshal(Ivrs)
		json.Unmarshal(IvrsString, &o.Ivrs)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Emergencycallflow) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
