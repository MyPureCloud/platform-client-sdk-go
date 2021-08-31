package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Stateventflowoutcometopicdatum
type Stateventflowoutcometopicdatum struct { 
	// Interval
	Interval *string `json:"interval,omitempty"`


	// Metrics
	Metrics *[]Stateventflowoutcometopicmetric `json:"metrics,omitempty"`

}

func (o *Stateventflowoutcometopicdatum) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Stateventflowoutcometopicdatum
	
	return json.Marshal(&struct { 
		Interval *string `json:"interval,omitempty"`
		
		Metrics *[]Stateventflowoutcometopicmetric `json:"metrics,omitempty"`
		*Alias
	}{ 
		Interval: o.Interval,
		
		Metrics: o.Metrics,
		Alias:    (*Alias)(o),
	})
}

func (o *Stateventflowoutcometopicdatum) UnmarshalJSON(b []byte) error {
	var StateventflowoutcometopicdatumMap map[string]interface{}
	err := json.Unmarshal(b, &StateventflowoutcometopicdatumMap)
	if err != nil {
		return err
	}
	
	if Interval, ok := StateventflowoutcometopicdatumMap["interval"].(string); ok {
		o.Interval = &Interval
	}
	
	if Metrics, ok := StateventflowoutcometopicdatumMap["metrics"].([]interface{}); ok {
		MetricsString, _ := json.Marshal(Metrics)
		json.Unmarshal(MetricsString, &o.Metrics)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Stateventflowoutcometopicdatum) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
