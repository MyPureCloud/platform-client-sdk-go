package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Predictorworkloadbalancing
type Predictorworkloadbalancing struct { 
	// Enabled - Flag to activate and deactivate workload balancing.
	Enabled *bool `json:"enabled,omitempty"`


	// MinimumOccupancy - Desired minimum occupancy threshold of agents. Must be between 0 and 100.
	MinimumOccupancy *int `json:"minimumOccupancy,omitempty"`


	// MaximumOccupancy - Desired maximum occupancy threshold of agents. Must be between 0 and 100.
	MaximumOccupancy *int `json:"maximumOccupancy,omitempty"`

}

func (o *Predictorworkloadbalancing) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Predictorworkloadbalancing
	
	return json.Marshal(&struct { 
		Enabled *bool `json:"enabled,omitempty"`
		
		MinimumOccupancy *int `json:"minimumOccupancy,omitempty"`
		
		MaximumOccupancy *int `json:"maximumOccupancy,omitempty"`
		*Alias
	}{ 
		Enabled: o.Enabled,
		
		MinimumOccupancy: o.MinimumOccupancy,
		
		MaximumOccupancy: o.MaximumOccupancy,
		Alias:    (*Alias)(o),
	})
}

func (o *Predictorworkloadbalancing) UnmarshalJSON(b []byte) error {
	var PredictorworkloadbalancingMap map[string]interface{}
	err := json.Unmarshal(b, &PredictorworkloadbalancingMap)
	if err != nil {
		return err
	}
	
	if Enabled, ok := PredictorworkloadbalancingMap["enabled"].(bool); ok {
		o.Enabled = &Enabled
	}
    
	if MinimumOccupancy, ok := PredictorworkloadbalancingMap["minimumOccupancy"].(float64); ok {
		MinimumOccupancyInt := int(MinimumOccupancy)
		o.MinimumOccupancy = &MinimumOccupancyInt
	}
	
	if MaximumOccupancy, ok := PredictorworkloadbalancingMap["maximumOccupancy"].(float64); ok {
		MaximumOccupancyInt := int(MaximumOccupancy)
		o.MaximumOccupancy = &MaximumOccupancyInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Predictorworkloadbalancing) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
