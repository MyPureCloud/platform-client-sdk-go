package platformclientv2
import (
	"encoding/json"
)

// Developmentactivityaggregatequeryresponsedata
type Developmentactivityaggregatequeryresponsedata struct { 
	// Interval - Specifies the range of due dates to be used for filtering. A maximum of 1 year can be specified in the range. Intervals are represented as an ISO-8601 string. For example: YYYY-MM-DDThh:mm:ss/YYYY-MM-DDThh:mm:ss
	Interval *string `json:"interval,omitempty"`


	// Metrics - The list of aggregated metrics
	Metrics *[]Developmentactivityaggregatequeryresponsemetric `json:"metrics,omitempty"`

}

// String returns a JSON representation of the model
func (o *Developmentactivityaggregatequeryresponsedata) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
