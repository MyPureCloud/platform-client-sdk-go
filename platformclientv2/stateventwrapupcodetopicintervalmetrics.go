package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Stateventwrapupcodetopicintervalmetrics
type Stateventwrapupcodetopicintervalmetrics struct { 
	// Interval
	Interval *string `json:"interval,omitempty"`


	// Metrics
	Metrics *[]Stateventwrapupcodetopicmetricstats `json:"metrics,omitempty"`

}

func (o *Stateventwrapupcodetopicintervalmetrics) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Stateventwrapupcodetopicintervalmetrics
	
	return json.Marshal(&struct { 
		Interval *string `json:"interval,omitempty"`
		
		Metrics *[]Stateventwrapupcodetopicmetricstats `json:"metrics,omitempty"`
		*Alias
	}{ 
		Interval: o.Interval,
		
		Metrics: o.Metrics,
		Alias:    (*Alias)(o),
	})
}

func (o *Stateventwrapupcodetopicintervalmetrics) UnmarshalJSON(b []byte) error {
	var StateventwrapupcodetopicintervalmetricsMap map[string]interface{}
	err := json.Unmarshal(b, &StateventwrapupcodetopicintervalmetricsMap)
	if err != nil {
		return err
	}
	
	if Interval, ok := StateventwrapupcodetopicintervalmetricsMap["interval"].(string); ok {
		o.Interval = &Interval
	}
	
	if Metrics, ok := StateventwrapupcodetopicintervalmetricsMap["metrics"].([]interface{}); ok {
		MetricsString, _ := json.Marshal(Metrics)
		json.Unmarshal(MetricsString, &o.Metrics)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Stateventwrapupcodetopicintervalmetrics) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
