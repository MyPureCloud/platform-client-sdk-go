package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Longtermforecastplanninggroupdata
type Longtermforecastplanninggroupdata struct { 
	// PlanningGroupId - The ID of the planning group to which this data applies. Note this is a snapshot of the planning group at the time of forecast creation and may not correspond to the current configuration
	PlanningGroupId *string `json:"planningGroupId,omitempty"`


	// OfferedPerDay - Forecast offered counts per day for this planning group
	OfferedPerDay *[]float64 `json:"offeredPerDay,omitempty"`


	// AverageHandleTimeSecondsPerDay - Forecast average handle time per day in seconds
	AverageHandleTimeSecondsPerDay *[]float64 `json:"averageHandleTimeSecondsPerDay,omitempty"`

}

func (o *Longtermforecastplanninggroupdata) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Longtermforecastplanninggroupdata
	
	return json.Marshal(&struct { 
		PlanningGroupId *string `json:"planningGroupId,omitempty"`
		
		OfferedPerDay *[]float64 `json:"offeredPerDay,omitempty"`
		
		AverageHandleTimeSecondsPerDay *[]float64 `json:"averageHandleTimeSecondsPerDay,omitempty"`
		*Alias
	}{ 
		PlanningGroupId: o.PlanningGroupId,
		
		OfferedPerDay: o.OfferedPerDay,
		
		AverageHandleTimeSecondsPerDay: o.AverageHandleTimeSecondsPerDay,
		Alias:    (*Alias)(o),
	})
}

func (o *Longtermforecastplanninggroupdata) UnmarshalJSON(b []byte) error {
	var LongtermforecastplanninggroupdataMap map[string]interface{}
	err := json.Unmarshal(b, &LongtermforecastplanninggroupdataMap)
	if err != nil {
		return err
	}
	
	if PlanningGroupId, ok := LongtermforecastplanninggroupdataMap["planningGroupId"].(string); ok {
		o.PlanningGroupId = &PlanningGroupId
	}
    
	if OfferedPerDay, ok := LongtermforecastplanninggroupdataMap["offeredPerDay"].([]interface{}); ok {
		OfferedPerDayString, _ := json.Marshal(OfferedPerDay)
		json.Unmarshal(OfferedPerDayString, &o.OfferedPerDay)
	}
	
	if AverageHandleTimeSecondsPerDay, ok := LongtermforecastplanninggroupdataMap["averageHandleTimeSecondsPerDay"].([]interface{}); ok {
		AverageHandleTimeSecondsPerDayString, _ := json.Marshal(AverageHandleTimeSecondsPerDay)
		json.Unmarshal(AverageHandleTimeSecondsPerDayString, &o.AverageHandleTimeSecondsPerDay)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Longtermforecastplanninggroupdata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
