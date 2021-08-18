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

func (u *Buplanninggroupheadcountforecast) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Buplanninggroupheadcountforecast

	

	return json.Marshal(&struct { 
		PlanningGroup *Planninggroupreference `json:"planningGroup,omitempty"`
		
		RequiredPerInterval *[]float64 `json:"requiredPerInterval,omitempty"`
		
		RequiredWithoutShrinkagePerInterval *[]float64 `json:"requiredWithoutShrinkagePerInterval,omitempty"`
		*Alias
	}{ 
		PlanningGroup: u.PlanningGroup,
		
		RequiredPerInterval: u.RequiredPerInterval,
		
		RequiredWithoutShrinkagePerInterval: u.RequiredWithoutShrinkagePerInterval,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Buplanninggroupheadcountforecast) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
