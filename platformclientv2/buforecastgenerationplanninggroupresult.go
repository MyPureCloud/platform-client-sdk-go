package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Buforecastgenerationplanninggroupresult
type Buforecastgenerationplanninggroupresult struct { 
	// PlanningGroupId - The ID of the planning group
	PlanningGroupId *string `json:"planningGroupId,omitempty"`


	// MetricResults - The generation results for the associated planning group
	MetricResults *[]Buforecasttimeseriesresult `json:"metricResults,omitempty"`

}

func (o *Buforecastgenerationplanninggroupresult) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Buforecastgenerationplanninggroupresult
	
	return json.Marshal(&struct { 
		PlanningGroupId *string `json:"planningGroupId,omitempty"`
		
		MetricResults *[]Buforecasttimeseriesresult `json:"metricResults,omitempty"`
		*Alias
	}{ 
		PlanningGroupId: o.PlanningGroupId,
		
		MetricResults: o.MetricResults,
		Alias:    (*Alias)(o),
	})
}

func (o *Buforecastgenerationplanninggroupresult) UnmarshalJSON(b []byte) error {
	var BuforecastgenerationplanninggroupresultMap map[string]interface{}
	err := json.Unmarshal(b, &BuforecastgenerationplanninggroupresultMap)
	if err != nil {
		return err
	}
	
	if PlanningGroupId, ok := BuforecastgenerationplanninggroupresultMap["planningGroupId"].(string); ok {
		o.PlanningGroupId = &PlanningGroupId
	}
    
	if MetricResults, ok := BuforecastgenerationplanninggroupresultMap["metricResults"].([]interface{}); ok {
		MetricResultsString, _ := json.Marshal(MetricResults)
		json.Unmarshal(MetricResultsString, &o.MetricResults)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Buforecastgenerationplanninggroupresult) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
