package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Developmentactivityaggregatequeryresponsedata
type Developmentactivityaggregatequeryresponsedata struct { 
	// Interval - Specifies the range of due dates to be used for filtering. A maximum of 1 year can be specified in the range. Intervals are represented as an ISO-8601 string. For example: YYYY-MM-DDThh:mm:ss/YYYY-MM-DDThh:mm:ss
	Interval *string `json:"interval,omitempty"`


	// Metrics - The list of aggregated metrics
	Metrics *[]Developmentactivityaggregatequeryresponsemetric `json:"metrics,omitempty"`

}

func (o *Developmentactivityaggregatequeryresponsedata) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Developmentactivityaggregatequeryresponsedata
	
	return json.Marshal(&struct { 
		Interval *string `json:"interval,omitempty"`
		
		Metrics *[]Developmentactivityaggregatequeryresponsemetric `json:"metrics,omitempty"`
		*Alias
	}{ 
		Interval: o.Interval,
		
		Metrics: o.Metrics,
		Alias:    (*Alias)(o),
	})
}

func (o *Developmentactivityaggregatequeryresponsedata) UnmarshalJSON(b []byte) error {
	var DevelopmentactivityaggregatequeryresponsedataMap map[string]interface{}
	err := json.Unmarshal(b, &DevelopmentactivityaggregatequeryresponsedataMap)
	if err != nil {
		return err
	}
	
	if Interval, ok := DevelopmentactivityaggregatequeryresponsedataMap["interval"].(string); ok {
		o.Interval = &Interval
	}
	
	if Metrics, ok := DevelopmentactivityaggregatequeryresponsedataMap["metrics"].([]interface{}); ok {
		MetricsString, _ := json.Marshal(Metrics)
		json.Unmarshal(MetricsString, &o.Metrics)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Developmentactivityaggregatequeryresponsedata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
