package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Developmentactivityaggregatequeryresponsemetric
type Developmentactivityaggregatequeryresponsemetric struct { 
	// Metric - The metric this applies to
	Metric *string `json:"metric,omitempty"`


	// Stats - The aggregated values for this metric
	Stats *Developmentactivityaggregatequeryresponsestatistics `json:"stats,omitempty"`

}

func (o *Developmentactivityaggregatequeryresponsemetric) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Developmentactivityaggregatequeryresponsemetric
	
	return json.Marshal(&struct { 
		Metric *string `json:"metric,omitempty"`
		
		Stats *Developmentactivityaggregatequeryresponsestatistics `json:"stats,omitempty"`
		*Alias
	}{ 
		Metric: o.Metric,
		
		Stats: o.Stats,
		Alias:    (*Alias)(o),
	})
}

func (o *Developmentactivityaggregatequeryresponsemetric) UnmarshalJSON(b []byte) error {
	var DevelopmentactivityaggregatequeryresponsemetricMap map[string]interface{}
	err := json.Unmarshal(b, &DevelopmentactivityaggregatequeryresponsemetricMap)
	if err != nil {
		return err
	}
	
	if Metric, ok := DevelopmentactivityaggregatequeryresponsemetricMap["metric"].(string); ok {
		o.Metric = &Metric
	}
	
	if Stats, ok := DevelopmentactivityaggregatequeryresponsemetricMap["stats"].(map[string]interface{}); ok {
		StatsString, _ := json.Marshal(Stats)
		json.Unmarshal(StatsString, &o.Stats)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Developmentactivityaggregatequeryresponsemetric) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
