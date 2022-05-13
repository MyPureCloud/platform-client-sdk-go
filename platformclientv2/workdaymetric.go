package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Workdaymetric
type Workdaymetric struct { 
	// Metric - Gamification metric
	Metric *Metric `json:"metric,omitempty"`


	// Objective - Current objective for this metric
	Objective *Objective `json:"objective,omitempty"`


	// Points - Gamification points earned for this metric
	Points *int `json:"points,omitempty"`


	// Value - Value of this metric
	Value *float64 `json:"value,omitempty"`


	// PunctualityEvents - List of schedule activity events for punctuality metrics
	PunctualityEvents *[]Punctualityevent `json:"punctualityEvents,omitempty"`

}

func (o *Workdaymetric) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Workdaymetric
	
	return json.Marshal(&struct { 
		Metric *Metric `json:"metric,omitempty"`
		
		Objective *Objective `json:"objective,omitempty"`
		
		Points *int `json:"points,omitempty"`
		
		Value *float64 `json:"value,omitempty"`
		
		PunctualityEvents *[]Punctualityevent `json:"punctualityEvents,omitempty"`
		*Alias
	}{ 
		Metric: o.Metric,
		
		Objective: o.Objective,
		
		Points: o.Points,
		
		Value: o.Value,
		
		PunctualityEvents: o.PunctualityEvents,
		Alias:    (*Alias)(o),
	})
}

func (o *Workdaymetric) UnmarshalJSON(b []byte) error {
	var WorkdaymetricMap map[string]interface{}
	err := json.Unmarshal(b, &WorkdaymetricMap)
	if err != nil {
		return err
	}
	
	if Metric, ok := WorkdaymetricMap["metric"].(map[string]interface{}); ok {
		MetricString, _ := json.Marshal(Metric)
		json.Unmarshal(MetricString, &o.Metric)
	}
	
	if Objective, ok := WorkdaymetricMap["objective"].(map[string]interface{}); ok {
		ObjectiveString, _ := json.Marshal(Objective)
		json.Unmarshal(ObjectiveString, &o.Objective)
	}
	
	if Points, ok := WorkdaymetricMap["points"].(float64); ok {
		PointsInt := int(Points)
		o.Points = &PointsInt
	}
	
	if Value, ok := WorkdaymetricMap["value"].(float64); ok {
		o.Value = &Value
	}
    
	if PunctualityEvents, ok := WorkdaymetricMap["punctualityEvents"].([]interface{}); ok {
		PunctualityEventsString, _ := json.Marshal(PunctualityEvents)
		json.Unmarshal(PunctualityEventsString, &o.PunctualityEvents)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Workdaymetric) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
