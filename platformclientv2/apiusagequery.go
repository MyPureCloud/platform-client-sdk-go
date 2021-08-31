package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Apiusagequery
type Apiusagequery struct { 
	// Interval - Behaves like one clause in a SQL WHERE. Specifies the date and time range of data being queried. Intervals are represented as an ISO-8601 string. For example: YYYY-MM-DDThh:mm:ss/YYYY-MM-DDThh:mm:ss
	Interval *string `json:"interval,omitempty"`


	// Granularity - Date granularity of the results
	Granularity *string `json:"granularity,omitempty"`


	// GroupBy - Behaves like a SQL GROUPBY. Allows for multiple levels of grouping as a list of dimensions. Partitions resulting aggregate computations into distinct named subgroups rather than across the entire result set as if it were one group.
	GroupBy *[]string `json:"groupBy,omitempty"`


	// Metrics - Behaves like a SQL SELECT clause. Enables retrieving only named metrics. If omitted, all metrics that are available will be returned (like SELECT *).
	Metrics *[]string `json:"metrics,omitempty"`

}

func (o *Apiusagequery) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Apiusagequery
	
	return json.Marshal(&struct { 
		Interval *string `json:"interval,omitempty"`
		
		Granularity *string `json:"granularity,omitempty"`
		
		GroupBy *[]string `json:"groupBy,omitempty"`
		
		Metrics *[]string `json:"metrics,omitempty"`
		*Alias
	}{ 
		Interval: o.Interval,
		
		Granularity: o.Granularity,
		
		GroupBy: o.GroupBy,
		
		Metrics: o.Metrics,
		Alias:    (*Alias)(o),
	})
}

func (o *Apiusagequery) UnmarshalJSON(b []byte) error {
	var ApiusagequeryMap map[string]interface{}
	err := json.Unmarshal(b, &ApiusagequeryMap)
	if err != nil {
		return err
	}
	
	if Interval, ok := ApiusagequeryMap["interval"].(string); ok {
		o.Interval = &Interval
	}
	
	if Granularity, ok := ApiusagequeryMap["granularity"].(string); ok {
		o.Granularity = &Granularity
	}
	
	if GroupBy, ok := ApiusagequeryMap["groupBy"].([]interface{}); ok {
		GroupByString, _ := json.Marshal(GroupBy)
		json.Unmarshal(GroupByString, &o.GroupBy)
	}
	
	if Metrics, ok := ApiusagequeryMap["metrics"].([]interface{}); ok {
		MetricsString, _ := json.Marshal(Metrics)
		json.Unmarshal(MetricsString, &o.Metrics)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Apiusagequery) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
