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

func (u *Emergencycallflow) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Emergencycallflow

	

	return json.Marshal(&struct { 
		EmergencyFlow *Domainentityref `json:"emergencyFlow,omitempty"`
		
		Ivrs *[]Domainentityref `json:"ivrs,omitempty"`
		*Alias
	}{ 
		EmergencyFlow: u.EmergencyFlow,
		
		Ivrs: u.Ivrs,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Emergencycallflow) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
