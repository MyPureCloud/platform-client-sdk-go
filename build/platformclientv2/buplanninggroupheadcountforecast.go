package platformclientv2
import (
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

// String returns a JSON representation of the model
func (o *Buplanninggroupheadcountforecast) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
