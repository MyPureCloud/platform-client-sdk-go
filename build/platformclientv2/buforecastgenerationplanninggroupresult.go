package platformclientv2
import (
	"encoding/json"
)

// Buforecastgenerationplanninggroupresult
type Buforecastgenerationplanninggroupresult struct { 
	// PlanningGroupId - The ID of the planning group
	PlanningGroupId *string `json:"planningGroupId,omitempty"`


	// MetricResults - The generation results for the associated planning group
	MetricResults *[]Buforecasttimeseriesresult `json:"metricResults,omitempty"`

}

// String returns a JSON representation of the model
func (o *Buforecastgenerationplanninggroupresult) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
