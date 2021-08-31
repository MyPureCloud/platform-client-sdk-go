package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Stateventqueuetopicdatum
type Stateventqueuetopicdatum struct { 
	// Interval
	Interval *string `json:"interval,omitempty"`


	// Metrics
	Metrics *[]Stateventqueuetopicmetric `json:"metrics,omitempty"`

}

func (o *Stateventqueuetopicdatum) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Stateventqueuetopicdatum
	
	return json.Marshal(&struct { 
		Interval *string `json:"interval,omitempty"`
		
		Metrics *[]Stateventqueuetopicmetric `json:"metrics,omitempty"`
		*Alias
	}{ 
		Interval: o.Interval,
		
		Metrics: o.Metrics,
		Alias:    (*Alias)(o),
	})
}

func (o *Stateventqueuetopicdatum) UnmarshalJSON(b []byte) error {
	var StateventqueuetopicdatumMap map[string]interface{}
	err := json.Unmarshal(b, &StateventqueuetopicdatumMap)
	if err != nil {
		return err
	}
	
	if Interval, ok := StateventqueuetopicdatumMap["interval"].(string); ok {
		o.Interval = &Interval
	}
	
	if Metrics, ok := StateventqueuetopicdatumMap["metrics"].([]interface{}); ok {
		MetricsString, _ := json.Marshal(Metrics)
		json.Unmarshal(MetricsString, &o.Metrics)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Stateventqueuetopicdatum) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
