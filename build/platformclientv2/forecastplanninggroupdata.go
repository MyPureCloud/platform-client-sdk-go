package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Forecastplanninggroupdata
type Forecastplanninggroupdata struct { 
	// PlanningGroupId - The ID of the planning group to which this data applies. Note this is a snapshot of the planning group at the time of forecast creation and may not correspond to the current configuration
	PlanningGroupId *string `json:"planningGroupId,omitempty"`


	// OfferedPerInterval - Forecast offered counts per interval for this week of the forecast
	OfferedPerInterval *[]float64 `json:"offeredPerInterval,omitempty"`


	// AverageHandleTimeSecondsPerInterval - Forecast average handle time per interval in seconds
	AverageHandleTimeSecondsPerInterval *[]float64 `json:"averageHandleTimeSecondsPerInterval,omitempty"`

}

// String returns a JSON representation of the model
func (o *Forecastplanninggroupdata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
