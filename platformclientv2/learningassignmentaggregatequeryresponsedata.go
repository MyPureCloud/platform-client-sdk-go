package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Learningassignmentaggregatequeryresponsedata
type Learningassignmentaggregatequeryresponsedata struct { 
	// Interval - Specifies the range of due dates to be used for filtering. A maximum of 1 year can be specified in the range. Intervals are represented as an ISO-8601 string. For example: YYYY-MM-DDThh:mm:ss/YYYY-MM-DDThh:mm:ss
	Interval *string `json:"interval,omitempty"`


	// Metrics - The list of aggregated metrics
	Metrics *[]Learningassignmentaggregatequeryresponsemetric `json:"metrics,omitempty"`

}

func (o *Learningassignmentaggregatequeryresponsedata) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Learningassignmentaggregatequeryresponsedata
	
	return json.Marshal(&struct { 
		Interval *string `json:"interval,omitempty"`
		
		Metrics *[]Learningassignmentaggregatequeryresponsemetric `json:"metrics,omitempty"`
		*Alias
	}{ 
		Interval: o.Interval,
		
		Metrics: o.Metrics,
		Alias:    (*Alias)(o),
	})
}

func (o *Learningassignmentaggregatequeryresponsedata) UnmarshalJSON(b []byte) error {
	var LearningassignmentaggregatequeryresponsedataMap map[string]interface{}
	err := json.Unmarshal(b, &LearningassignmentaggregatequeryresponsedataMap)
	if err != nil {
		return err
	}
	
	if Interval, ok := LearningassignmentaggregatequeryresponsedataMap["interval"].(string); ok {
		o.Interval = &Interval
	}
	
	if Metrics, ok := LearningassignmentaggregatequeryresponsedataMap["metrics"].([]interface{}); ok {
		MetricsString, _ := json.Marshal(Metrics)
		json.Unmarshal(MetricsString, &o.Metrics)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Learningassignmentaggregatequeryresponsedata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
