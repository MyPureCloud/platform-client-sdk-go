package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Coachingappointmentaggregaterequest
type Coachingappointmentaggregaterequest struct { 
	// Interval - Interval to aggregate across. End date is not inclusive. Intervals are represented as an ISO-8601 string. For example: YYYY-MM-DDThh:mm:ss/YYYY-MM-DDThh:mm:ss
	Interval *string `json:"interval,omitempty"`


	// Metrics - A list of metrics to aggregate.  If omitted, all metrics are returned.
	Metrics *[]string `json:"metrics,omitempty"`


	// GroupBy - An optional list of items by which to group the result data.
	GroupBy *[]string `json:"groupBy,omitempty"`


	// Filter - The filter applied to the data
	Filter *Queryrequestfilter `json:"filter,omitempty"`

}

// String returns a JSON representation of the model
func (o *Coachingappointmentaggregaterequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
