package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Gamificationscorecardchangetopicperformancemetric
type Gamificationscorecardchangetopicperformancemetric struct { 
	// Metric
	Metric *Gamificationscorecardchangetopicmetric `json:"metric,omitempty"`


	// Points
	Points *int `json:"points,omitempty"`


	// Value
	Value *float32 `json:"value,omitempty"`


	// PunctualityEvents
	PunctualityEvents *[]Gamificationscorecardchangetopicpunctualityevent `json:"punctualityEvents,omitempty"`

}

func (o *Gamificationscorecardchangetopicperformancemetric) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Gamificationscorecardchangetopicperformancemetric
	
	return json.Marshal(&struct { 
		Metric *Gamificationscorecardchangetopicmetric `json:"metric,omitempty"`
		
		Points *int `json:"points,omitempty"`
		
		Value *float32 `json:"value,omitempty"`
		
		PunctualityEvents *[]Gamificationscorecardchangetopicpunctualityevent `json:"punctualityEvents,omitempty"`
		*Alias
	}{ 
		Metric: o.Metric,
		
		Points: o.Points,
		
		Value: o.Value,
		
		PunctualityEvents: o.PunctualityEvents,
		Alias:    (*Alias)(o),
	})
}

func (o *Gamificationscorecardchangetopicperformancemetric) UnmarshalJSON(b []byte) error {
	var GamificationscorecardchangetopicperformancemetricMap map[string]interface{}
	err := json.Unmarshal(b, &GamificationscorecardchangetopicperformancemetricMap)
	if err != nil {
		return err
	}
	
	if Metric, ok := GamificationscorecardchangetopicperformancemetricMap["metric"].(map[string]interface{}); ok {
		MetricString, _ := json.Marshal(Metric)
		json.Unmarshal(MetricString, &o.Metric)
	}
	
	if Points, ok := GamificationscorecardchangetopicperformancemetricMap["points"].(float64); ok {
		PointsInt := int(Points)
		o.Points = &PointsInt
	}
	
	if Value, ok := GamificationscorecardchangetopicperformancemetricMap["value"].(float64); ok {
		ValueFloat32 := float32(Value)
		o.Value = &ValueFloat32
	}
    
	if PunctualityEvents, ok := GamificationscorecardchangetopicperformancemetricMap["punctualityEvents"].([]interface{}); ok {
		PunctualityEventsString, _ := json.Marshal(PunctualityEvents)
		json.Unmarshal(PunctualityEventsString, &o.PunctualityEvents)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Gamificationscorecardchangetopicperformancemetric) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
