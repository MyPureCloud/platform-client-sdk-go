package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Workdayvaluesmetricitem
type Workdayvaluesmetricitem struct { 
	// MetricDefinition - Gamification metric for the average and the trend
	MetricDefinition *Domainentityref `json:"metricDefinition,omitempty"`


	// Average - The average value of the metric
	Average *float64 `json:"average,omitempty"`


	// UnitType - The unit type of the metric value
	UnitType *string `json:"unitType,omitempty"`


	// Trend - The metric value trend
	Trend *[]Workdayvaluestrenditem `json:"trend,omitempty"`

}

func (u *Workdayvaluesmetricitem) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Workdayvaluesmetricitem

	

	return json.Marshal(&struct { 
		MetricDefinition *Domainentityref `json:"metricDefinition,omitempty"`
		
		Average *float64 `json:"average,omitempty"`
		
		UnitType *string `json:"unitType,omitempty"`
		
		Trend *[]Workdayvaluestrenditem `json:"trend,omitempty"`
		*Alias
	}{ 
		MetricDefinition: u.MetricDefinition,
		
		Average: u.Average,
		
		UnitType: u.UnitType,
		
		Trend: u.Trend,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Workdayvaluesmetricitem) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
