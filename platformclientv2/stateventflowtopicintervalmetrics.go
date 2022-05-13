package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Stateventflowtopicintervalmetrics
type Stateventflowtopicintervalmetrics struct { 
	// Interval
	Interval *string `json:"interval,omitempty"`


	// Metrics
	Metrics *[]Stateventflowtopicmetricstats `json:"metrics,omitempty"`

}

func (o *Stateventflowtopicintervalmetrics) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Stateventflowtopicintervalmetrics
	
	return json.Marshal(&struct { 
		Interval *string `json:"interval,omitempty"`
		
		Metrics *[]Stateventflowtopicmetricstats `json:"metrics,omitempty"`
		*Alias
	}{ 
		Interval: o.Interval,
		
		Metrics: o.Metrics,
		Alias:    (*Alias)(o),
	})
}

func (o *Stateventflowtopicintervalmetrics) UnmarshalJSON(b []byte) error {
	var StateventflowtopicintervalmetricsMap map[string]interface{}
	err := json.Unmarshal(b, &StateventflowtopicintervalmetricsMap)
	if err != nil {
		return err
	}
	
	if Interval, ok := StateventflowtopicintervalmetricsMap["interval"].(string); ok {
		o.Interval = &Interval
	}
    
	if Metrics, ok := StateventflowtopicintervalmetricsMap["metrics"].([]interface{}); ok {
		MetricsString, _ := json.Marshal(Metrics)
		json.Unmarshal(MetricsString, &o.Metrics)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Stateventflowtopicintervalmetrics) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
