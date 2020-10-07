package platformclientv2
import (
	"encoding/json"
)

// Developmentactivityaggregateparam
type Developmentactivityaggregateparam struct { 
	// Interval - Specifies the range of due dates to be used for filtering. Milliseconds will be truncated. A maximum of 1 year can be specified in the range. Intervals are represented as an ISO-8601 string. For example: YYYY-MM-DDThh:mm:ss/YYYY-MM-DDThh:mm:ss
	Interval *string `json:"interval,omitempty"`


	// Metrics - The list of metrics to be returned. If omitted, all metrics are returned.
	Metrics *[]string `json:"metrics,omitempty"`


	// GroupBy - Specifies if the aggregated data is combined into a single set of metrics (groupBy is empty or not specified), or contains an element per attendeeId (groupBy is \"attendeeId\")
	GroupBy *[]string `json:"groupBy,omitempty"`


	// Filter - The filter applied to the data. This is ANDed with the interval parameter.
	Filter *Developmentactivityaggregatequeryrequestfilter `json:"filter,omitempty"`

}

// String returns a JSON representation of the model
func (o *Developmentactivityaggregateparam) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
