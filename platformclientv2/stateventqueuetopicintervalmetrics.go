package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Stateventqueuetopicintervalmetrics
type Stateventqueuetopicintervalmetrics struct { 
	// Interval
	Interval *string `json:"interval,omitempty"`


	// Metrics
	Metrics *[]Stateventqueuetopicmetricstats `json:"metrics,omitempty"`

}

func (o *Stateventqueuetopicintervalmetrics) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Stateventqueuetopicintervalmetrics
	
	return json.Marshal(&struct { 
		Interval *string `json:"interval,omitempty"`
		
		Metrics *[]Stateventqueuetopicmetricstats `json:"metrics,omitempty"`
		*Alias
	}{ 
		Interval: o.Interval,
		
		Metrics: o.Metrics,
		Alias:    (*Alias)(o),
	})
}

func (o *Stateventqueuetopicintervalmetrics) UnmarshalJSON(b []byte) error {
	var StateventqueuetopicintervalmetricsMap map[string]interface{}
	err := json.Unmarshal(b, &StateventqueuetopicintervalmetricsMap)
	if err != nil {
		return err
	}
	
	if Interval, ok := StateventqueuetopicintervalmetricsMap["interval"].(string); ok {
		o.Interval = &Interval
	}
    
	if Metrics, ok := StateventqueuetopicintervalmetricsMap["metrics"].([]interface{}); ok {
		MetricsString, _ := json.Marshal(Metrics)
		json.Unmarshal(MetricsString, &o.Metrics)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Stateventqueuetopicintervalmetrics) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
