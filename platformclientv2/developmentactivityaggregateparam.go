package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Developmentactivityaggregateparam
type Developmentactivityaggregateparam struct { 
	// Interval - Specifies the range of due dates to be used for filtering. Milliseconds will be truncated. A maximum of 1 year can be specified in the range. End date is not inclusive. Intervals are represented as an ISO-8601 string. For example: YYYY-MM-DDThh:mm:ss/YYYY-MM-DDThh:mm:ss
	Interval *string `json:"interval,omitempty"`


	// Metrics - The list of metrics to be returned. If omitted, all metrics are returned.
	Metrics *[]string `json:"metrics,omitempty"`


	// GroupBy - Specifies if the aggregated data is combined into a single set of metrics (groupBy is empty or not specified), or contains an element per attendeeId (groupBy is \"attendeeId\")
	GroupBy *[]string `json:"groupBy,omitempty"`


	// Filter - The filter applied to the data. This is ANDed with the interval parameter.
	Filter *Developmentactivityaggregatequeryrequestfilter `json:"filter,omitempty"`

}

func (o *Developmentactivityaggregateparam) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Developmentactivityaggregateparam
	
	return json.Marshal(&struct { 
		Interval *string `json:"interval,omitempty"`
		
		Metrics *[]string `json:"metrics,omitempty"`
		
		GroupBy *[]string `json:"groupBy,omitempty"`
		
		Filter *Developmentactivityaggregatequeryrequestfilter `json:"filter,omitempty"`
		*Alias
	}{ 
		Interval: o.Interval,
		
		Metrics: o.Metrics,
		
		GroupBy: o.GroupBy,
		
		Filter: o.Filter,
		Alias:    (*Alias)(o),
	})
}

func (o *Developmentactivityaggregateparam) UnmarshalJSON(b []byte) error {
	var DevelopmentactivityaggregateparamMap map[string]interface{}
	err := json.Unmarshal(b, &DevelopmentactivityaggregateparamMap)
	if err != nil {
		return err
	}
	
	if Interval, ok := DevelopmentactivityaggregateparamMap["interval"].(string); ok {
		o.Interval = &Interval
	}
    
	if Metrics, ok := DevelopmentactivityaggregateparamMap["metrics"].([]interface{}); ok {
		MetricsString, _ := json.Marshal(Metrics)
		json.Unmarshal(MetricsString, &o.Metrics)
	}
	
	if GroupBy, ok := DevelopmentactivityaggregateparamMap["groupBy"].([]interface{}); ok {
		GroupByString, _ := json.Marshal(GroupBy)
		json.Unmarshal(GroupByString, &o.GroupBy)
	}
	
	if Filter, ok := DevelopmentactivityaggregateparamMap["filter"].(map[string]interface{}); ok {
		FilterString, _ := json.Marshal(Filter)
		json.Unmarshal(FilterString, &o.Filter)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Developmentactivityaggregateparam) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
