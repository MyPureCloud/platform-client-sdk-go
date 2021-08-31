package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Flowobservationquery
type Flowobservationquery struct { 
	// Filter - Filter to return a subset of observations. Expresses boolean logical predicates as well as dimensional filters
	Filter *Flowobservationqueryfilter `json:"filter,omitempty"`


	// Metrics - Behaves like a SQL SELECT clause. Only named metrics will be retrieved.
	Metrics *[]string `json:"metrics,omitempty"`


	// DetailMetrics - Metrics for which to include additional detailed observations
	DetailMetrics *[]string `json:"detailMetrics,omitempty"`

}

func (o *Flowobservationquery) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Flowobservationquery
	
	return json.Marshal(&struct { 
		Filter *Flowobservationqueryfilter `json:"filter,omitempty"`
		
		Metrics *[]string `json:"metrics,omitempty"`
		
		DetailMetrics *[]string `json:"detailMetrics,omitempty"`
		*Alias
	}{ 
		Filter: o.Filter,
		
		Metrics: o.Metrics,
		
		DetailMetrics: o.DetailMetrics,
		Alias:    (*Alias)(o),
	})
}

func (o *Flowobservationquery) UnmarshalJSON(b []byte) error {
	var FlowobservationqueryMap map[string]interface{}
	err := json.Unmarshal(b, &FlowobservationqueryMap)
	if err != nil {
		return err
	}
	
	if Filter, ok := FlowobservationqueryMap["filter"].(map[string]interface{}); ok {
		FilterString, _ := json.Marshal(Filter)
		json.Unmarshal(FilterString, &o.Filter)
	}
	
	if Metrics, ok := FlowobservationqueryMap["metrics"].([]interface{}); ok {
		MetricsString, _ := json.Marshal(Metrics)
		json.Unmarshal(MetricsString, &o.Metrics)
	}
	
	if DetailMetrics, ok := FlowobservationqueryMap["detailMetrics"].([]interface{}); ok {
		DetailMetricsString, _ := json.Marshal(DetailMetrics)
		json.Unmarshal(DetailMetricsString, &o.DetailMetrics)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Flowobservationquery) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
