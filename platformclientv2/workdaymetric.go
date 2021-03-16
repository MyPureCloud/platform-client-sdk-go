package platformclientv2
import (
	"encoding/json"
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

// String returns a JSON representation of the model
func (o *Workdaymetric) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
