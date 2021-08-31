package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Stateventflowtopicdatum
type Stateventflowtopicdatum struct { 
	// Interval
	Interval *string `json:"interval,omitempty"`


	// Metrics
	Metrics *[]Stateventflowtopicmetric `json:"metrics,omitempty"`

}

func (o *Stateventflowtopicdatum) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Stateventflowtopicdatum
	
	return json.Marshal(&struct { 
		Interval *string `json:"interval,omitempty"`
		
		Metrics *[]Stateventflowtopicmetric `json:"metrics,omitempty"`
		*Alias
	}{ 
		Interval: o.Interval,
		
		Metrics: o.Metrics,
		Alias:    (*Alias)(o),
	})
}

func (o *Stateventflowtopicdatum) UnmarshalJSON(b []byte) error {
	var StateventflowtopicdatumMap map[string]interface{}
	err := json.Unmarshal(b, &StateventflowtopicdatumMap)
	if err != nil {
		return err
	}
	
	if Interval, ok := StateventflowtopicdatumMap["interval"].(string); ok {
		o.Interval = &Interval
	}
	
	if Metrics, ok := StateventflowtopicdatumMap["metrics"].([]interface{}); ok {
		MetricsString, _ := json.Marshal(Metrics)
		json.Unmarshal(MetricsString, &o.Metrics)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Stateventflowtopicdatum) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
