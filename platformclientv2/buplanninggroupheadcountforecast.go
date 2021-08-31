package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Buplanninggroupheadcountforecast
type Buplanninggroupheadcountforecast struct { 
	// PlanningGroup - The planning group to which this portion of the headcount forecast applies
	PlanningGroup *Planninggroupreference `json:"planningGroup,omitempty"`


	// RequiredPerInterval - Required headcount per interval, referenced against the reference start date
	RequiredPerInterval *[]float64 `json:"requiredPerInterval,omitempty"`


	// RequiredWithoutShrinkagePerInterval - Required headcount per interval without accounting for shrinkage, referenced against the reference start date
	RequiredWithoutShrinkagePerInterval *[]float64 `json:"requiredWithoutShrinkagePerInterval,omitempty"`

}

func (o *Buplanninggroupheadcountforecast) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Buplanninggroupheadcountforecast
	
	return json.Marshal(&struct { 
		PlanningGroup *Planninggroupreference `json:"planningGroup,omitempty"`
		
		RequiredPerInterval *[]float64 `json:"requiredPerInterval,omitempty"`
		
		RequiredWithoutShrinkagePerInterval *[]float64 `json:"requiredWithoutShrinkagePerInterval,omitempty"`
		*Alias
	}{ 
		PlanningGroup: o.PlanningGroup,
		
		RequiredPerInterval: o.RequiredPerInterval,
		
		RequiredWithoutShrinkagePerInterval: o.RequiredWithoutShrinkagePerInterval,
		Alias:    (*Alias)(o),
	})
}

func (o *Buplanninggroupheadcountforecast) UnmarshalJSON(b []byte) error {
	var BuplanninggroupheadcountforecastMap map[string]interface{}
	err := json.Unmarshal(b, &BuplanninggroupheadcountforecastMap)
	if err != nil {
		return err
	}
	
	if PlanningGroup, ok := BuplanninggroupheadcountforecastMap["planningGroup"].(map[string]interface{}); ok {
		PlanningGroupString, _ := json.Marshal(PlanningGroup)
		json.Unmarshal(PlanningGroupString, &o.PlanningGroup)
	}
	
	if RequiredPerInterval, ok := BuplanninggroupheadcountforecastMap["requiredPerInterval"].([]interface{}); ok {
		RequiredPerIntervalString, _ := json.Marshal(RequiredPerInterval)
		json.Unmarshal(RequiredPerIntervalString, &o.RequiredPerInterval)
	}
	
	if RequiredWithoutShrinkagePerInterval, ok := BuplanninggroupheadcountforecastMap["requiredWithoutShrinkagePerInterval"].([]interface{}); ok {
		RequiredWithoutShrinkagePerIntervalString, _ := json.Marshal(RequiredWithoutShrinkagePerInterval)
		json.Unmarshal(RequiredWithoutShrinkagePerIntervalString, &o.RequiredWithoutShrinkagePerInterval)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Buplanninggroupheadcountforecast) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
