package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Stateventusertopicintervalmetrics
type Stateventusertopicintervalmetrics struct { 
	// Interval
	Interval *string `json:"interval,omitempty"`


	// Metrics
	Metrics *[]Stateventusertopicmetricstats `json:"metrics,omitempty"`

}

func (o *Stateventusertopicintervalmetrics) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Stateventusertopicintervalmetrics
	
	return json.Marshal(&struct { 
		Interval *string `json:"interval,omitempty"`
		
		Metrics *[]Stateventusertopicmetricstats `json:"metrics,omitempty"`
		*Alias
	}{ 
		Interval: o.Interval,
		
		Metrics: o.Metrics,
		Alias:    (*Alias)(o),
	})
}

func (o *Stateventusertopicintervalmetrics) UnmarshalJSON(b []byte) error {
	var StateventusertopicintervalmetricsMap map[string]interface{}
	err := json.Unmarshal(b, &StateventusertopicintervalmetricsMap)
	if err != nil {
		return err
	}
	
	if Interval, ok := StateventusertopicintervalmetricsMap["interval"].(string); ok {
		o.Interval = &Interval
	}
    
	if Metrics, ok := StateventusertopicintervalmetricsMap["metrics"].([]interface{}); ok {
		MetricsString, _ := json.Marshal(Metrics)
		json.Unmarshal(MetricsString, &o.Metrics)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Stateventusertopicintervalmetrics) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
