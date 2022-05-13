package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Stateventflowoutcometopicintervalmetrics
type Stateventflowoutcometopicintervalmetrics struct { 
	// Interval
	Interval *string `json:"interval,omitempty"`


	// Metrics
	Metrics *[]Stateventflowoutcometopicmetricstats `json:"metrics,omitempty"`

}

func (o *Stateventflowoutcometopicintervalmetrics) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Stateventflowoutcometopicintervalmetrics
	
	return json.Marshal(&struct { 
		Interval *string `json:"interval,omitempty"`
		
		Metrics *[]Stateventflowoutcometopicmetricstats `json:"metrics,omitempty"`
		*Alias
	}{ 
		Interval: o.Interval,
		
		Metrics: o.Metrics,
		Alias:    (*Alias)(o),
	})
}

func (o *Stateventflowoutcometopicintervalmetrics) UnmarshalJSON(b []byte) error {
	var StateventflowoutcometopicintervalmetricsMap map[string]interface{}
	err := json.Unmarshal(b, &StateventflowoutcometopicintervalmetricsMap)
	if err != nil {
		return err
	}
	
	if Interval, ok := StateventflowoutcometopicintervalmetricsMap["interval"].(string); ok {
		o.Interval = &Interval
	}
    
	if Metrics, ok := StateventflowoutcometopicintervalmetricsMap["metrics"].([]interface{}); ok {
		MetricsString, _ := json.Marshal(Metrics)
		json.Unmarshal(MetricsString, &o.Metrics)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Stateventflowoutcometopicintervalmetrics) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
