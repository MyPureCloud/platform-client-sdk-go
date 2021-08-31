package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Queryresponsemetric
type Queryresponsemetric struct { 
	// Metric - The metric this applies to
	Metric *string `json:"metric,omitempty"`


	// Stats - The aggregated values for this metric
	Stats *Queryresponsestats `json:"stats,omitempty"`

}

func (o *Queryresponsemetric) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Queryresponsemetric
	
	return json.Marshal(&struct { 
		Metric *string `json:"metric,omitempty"`
		
		Stats *Queryresponsestats `json:"stats,omitempty"`
		*Alias
	}{ 
		Metric: o.Metric,
		
		Stats: o.Stats,
		Alias:    (*Alias)(o),
	})
}

func (o *Queryresponsemetric) UnmarshalJSON(b []byte) error {
	var QueryresponsemetricMap map[string]interface{}
	err := json.Unmarshal(b, &QueryresponsemetricMap)
	if err != nil {
		return err
	}
	
	if Metric, ok := QueryresponsemetricMap["metric"].(string); ok {
		o.Metric = &Metric
	}
	
	if Stats, ok := QueryresponsemetricMap["stats"].(map[string]interface{}); ok {
		StatsString, _ := json.Marshal(Stats)
		json.Unmarshal(StatsString, &o.Stats)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Queryresponsemetric) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
