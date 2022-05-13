package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Workdayvaluesmetricitem
type Workdayvaluesmetricitem struct { 
	// Metric - Gamification metric for the average and the trend
	Metric *Addressableentityref `json:"metric,omitempty"`


	// MetricDefinition - Gamification metric definition for the average and the trend
	MetricDefinition *Domainentityref `json:"metricDefinition,omitempty"`


	// Average - The average value of the metric
	Average *float64 `json:"average,omitempty"`


	// UnitType - The unit type of the metric value
	UnitType *string `json:"unitType,omitempty"`


	// Trend - The metric value trend
	Trend *[]Workdayvaluestrenditem `json:"trend,omitempty"`

}

func (o *Workdayvaluesmetricitem) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Workdayvaluesmetricitem
	
	return json.Marshal(&struct { 
		Metric *Addressableentityref `json:"metric,omitempty"`
		
		MetricDefinition *Domainentityref `json:"metricDefinition,omitempty"`
		
		Average *float64 `json:"average,omitempty"`
		
		UnitType *string `json:"unitType,omitempty"`
		
		Trend *[]Workdayvaluestrenditem `json:"trend,omitempty"`
		*Alias
	}{ 
		Metric: o.Metric,
		
		MetricDefinition: o.MetricDefinition,
		
		Average: o.Average,
		
		UnitType: o.UnitType,
		
		Trend: o.Trend,
		Alias:    (*Alias)(o),
	})
}

func (o *Workdayvaluesmetricitem) UnmarshalJSON(b []byte) error {
	var WorkdayvaluesmetricitemMap map[string]interface{}
	err := json.Unmarshal(b, &WorkdayvaluesmetricitemMap)
	if err != nil {
		return err
	}
	
	if Metric, ok := WorkdayvaluesmetricitemMap["metric"].(map[string]interface{}); ok {
		MetricString, _ := json.Marshal(Metric)
		json.Unmarshal(MetricString, &o.Metric)
	}
	
	if MetricDefinition, ok := WorkdayvaluesmetricitemMap["metricDefinition"].(map[string]interface{}); ok {
		MetricDefinitionString, _ := json.Marshal(MetricDefinition)
		json.Unmarshal(MetricDefinitionString, &o.MetricDefinition)
	}
	
	if Average, ok := WorkdayvaluesmetricitemMap["average"].(float64); ok {
		o.Average = &Average
	}
    
	if UnitType, ok := WorkdayvaluesmetricitemMap["unitType"].(string); ok {
		o.UnitType = &UnitType
	}
    
	if Trend, ok := WorkdayvaluesmetricitemMap["trend"].([]interface{}); ok {
		TrendString, _ := json.Marshal(Trend)
		json.Unmarshal(TrendString, &o.Trend)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Workdayvaluesmetricitem) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
