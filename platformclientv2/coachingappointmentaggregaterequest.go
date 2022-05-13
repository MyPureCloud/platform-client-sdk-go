package platformclientv2
import (
	"github.com/leekchan/timeutil"
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

func (o *Coachingappointmentaggregaterequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Coachingappointmentaggregaterequest
	
	return json.Marshal(&struct { 
		Interval *string `json:"interval,omitempty"`
		
		Metrics *[]string `json:"metrics,omitempty"`
		
		GroupBy *[]string `json:"groupBy,omitempty"`
		
		Filter *Queryrequestfilter `json:"filter,omitempty"`
		*Alias
	}{ 
		Interval: o.Interval,
		
		Metrics: o.Metrics,
		
		GroupBy: o.GroupBy,
		
		Filter: o.Filter,
		Alias:    (*Alias)(o),
	})
}

func (o *Coachingappointmentaggregaterequest) UnmarshalJSON(b []byte) error {
	var CoachingappointmentaggregaterequestMap map[string]interface{}
	err := json.Unmarshal(b, &CoachingappointmentaggregaterequestMap)
	if err != nil {
		return err
	}
	
	if Interval, ok := CoachingappointmentaggregaterequestMap["interval"].(string); ok {
		o.Interval = &Interval
	}
    
	if Metrics, ok := CoachingappointmentaggregaterequestMap["metrics"].([]interface{}); ok {
		MetricsString, _ := json.Marshal(Metrics)
		json.Unmarshal(MetricsString, &o.Metrics)
	}
	
	if GroupBy, ok := CoachingappointmentaggregaterequestMap["groupBy"].([]interface{}); ok {
		GroupByString, _ := json.Marshal(GroupBy)
		json.Unmarshal(GroupByString, &o.GroupBy)
	}
	
	if Filter, ok := CoachingappointmentaggregaterequestMap["filter"].(map[string]interface{}); ok {
		FilterString, _ := json.Marshal(Filter)
		json.Unmarshal(FilterString, &o.Filter)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Coachingappointmentaggregaterequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
