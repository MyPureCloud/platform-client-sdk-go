package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Stateventwrapupcodetopicdatum
type Stateventwrapupcodetopicdatum struct { 
	// Interval
	Interval *string `json:"interval,omitempty"`


	// Metrics
	Metrics *[]Stateventwrapupcodetopicmetric `json:"metrics,omitempty"`

}

func (o *Stateventwrapupcodetopicdatum) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Stateventwrapupcodetopicdatum
	
	return json.Marshal(&struct { 
		Interval *string `json:"interval,omitempty"`
		
		Metrics *[]Stateventwrapupcodetopicmetric `json:"metrics,omitempty"`
		*Alias
	}{ 
		Interval: o.Interval,
		
		Metrics: o.Metrics,
		Alias:    (*Alias)(o),
	})
}

func (o *Stateventwrapupcodetopicdatum) UnmarshalJSON(b []byte) error {
	var StateventwrapupcodetopicdatumMap map[string]interface{}
	err := json.Unmarshal(b, &StateventwrapupcodetopicdatumMap)
	if err != nil {
		return err
	}
	
	if Interval, ok := StateventwrapupcodetopicdatumMap["interval"].(string); ok {
		o.Interval = &Interval
	}
	
	if Metrics, ok := StateventwrapupcodetopicdatumMap["metrics"].([]interface{}); ok {
		MetricsString, _ := json.Marshal(Metrics)
		json.Unmarshal(MetricsString, &o.Metrics)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Stateventwrapupcodetopicdatum) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
