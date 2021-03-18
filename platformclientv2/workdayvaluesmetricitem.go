package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Workdayvaluesmetricitem
type Workdayvaluesmetricitem struct { 
	// MetricDefinition - Gamification metric for the average and the trend
	MetricDefinition *Metricdefinition `json:"metricDefinition,omitempty"`


	// Average - The average value of the metric
	Average *float64 `json:"average,omitempty"`


	// UnitType - The unit type of the metric value
	UnitType *string `json:"unitType,omitempty"`


	// Trend - The metric value trend
	Trend *[]Workdayvaluestrenditem `json:"trend,omitempty"`

}

// String returns a JSON representation of the model
func (o *Workdayvaluesmetricitem) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
