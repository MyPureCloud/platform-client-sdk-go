package platformclientv2
import (
	"encoding/json"
)

// Forecastplanninggroupdata
type Forecastplanninggroupdata struct { 
	// PlanningGroupId - The id of the planning group to which this data applies
	PlanningGroupId *string `json:"planningGroupId,omitempty"`


	// OfferedPerInterval - Forecast offered counts per interval for this week of the forecast
	OfferedPerInterval *[]float64 `json:"offeredPerInterval,omitempty"`


	// AverageHandleTimeSecondsPerInterval - Forecast average handle time per interval in seconds
	AverageHandleTimeSecondsPerInterval *[]float64 `json:"averageHandleTimeSecondsPerInterval,omitempty"`

}

// String returns a JSON representation of the model
func (o *Forecastplanninggroupdata) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
