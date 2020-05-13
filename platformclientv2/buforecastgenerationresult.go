package platformclientv2
import (
	"encoding/json"
)

// Buforecastgenerationresult
type Buforecastgenerationresult struct { 
	// PlanningGroupResults - Generation results, broken down by planning group
	PlanningGroupResults *[]Buforecastgenerationplanninggroupresult `json:"planningGroupResults,omitempty"`

}

// String returns a JSON representation of the model
func (o *Buforecastgenerationresult) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
