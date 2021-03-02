package platformclientv2
import (
	"encoding/json"
)

// Queueobservationquery
type Queueobservationquery struct { 
	// Filter - Filter to return a subset of observations. Expresses boolean logical predicates as well as dimensional filters
	Filter *Queueobservationqueryfilter `json:"filter,omitempty"`


	// Metrics - Behaves like a SQL SELECT clause. Only named metrics will be retrieved.
	Metrics *[]string `json:"metrics,omitempty"`


	// DetailMetrics - Metrics for which to include additional detailed observations
	DetailMetrics *[]string `json:"detailMetrics,omitempty"`

}

// String returns a JSON representation of the model
func (o *Queueobservationquery) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
