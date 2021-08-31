package platformclientv2
import (
	"github.com/leekchan/timeutil"
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

func (o *Forecastplanninggroupdata) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Forecastplanninggroupdata
	
	return json.Marshal(&struct { 
		PlanningGroupId *string `json:"planningGroupId,omitempty"`
		
		OfferedPerInterval *[]float64 `json:"offeredPerInterval,omitempty"`
		
		AverageHandleTimeSecondsPerInterval *[]float64 `json:"averageHandleTimeSecondsPerInterval,omitempty"`
		*Alias
	}{ 
		PlanningGroupId: o.PlanningGroupId,
		
		OfferedPerInterval: o.OfferedPerInterval,
		
		AverageHandleTimeSecondsPerInterval: o.AverageHandleTimeSecondsPerInterval,
		Alias:    (*Alias)(o),
	})
}

func (o *Forecastplanninggroupdata) UnmarshalJSON(b []byte) error {
	var ForecastplanninggroupdataMap map[string]interface{}
	err := json.Unmarshal(b, &ForecastplanninggroupdataMap)
	if err != nil {
		return err
	}
	
	if PlanningGroupId, ok := ForecastplanninggroupdataMap["planningGroupId"].(string); ok {
		o.PlanningGroupId = &PlanningGroupId
	}
	
	if OfferedPerInterval, ok := ForecastplanninggroupdataMap["offeredPerInterval"].([]interface{}); ok {
		OfferedPerIntervalString, _ := json.Marshal(OfferedPerInterval)
		json.Unmarshal(OfferedPerIntervalString, &o.OfferedPerInterval)
	}
	
	if AverageHandleTimeSecondsPerInterval, ok := ForecastplanninggroupdataMap["averageHandleTimeSecondsPerInterval"].([]interface{}); ok {
		AverageHandleTimeSecondsPerIntervalString, _ := json.Marshal(AverageHandleTimeSecondsPerInterval)
		json.Unmarshal(AverageHandleTimeSecondsPerIntervalString, &o.AverageHandleTimeSecondsPerInterval)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Forecastplanninggroupdata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
