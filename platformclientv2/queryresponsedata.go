package platformclientv2
import (
	"encoding/json"
)

// Queryresponsedata
type Queryresponsedata struct { 
	// Interval - Interval with start and end represented as ISO-8601 string. i.e: yyyy-MM-dd'T'HH:mm:ss.SSS'Z'/yyyy-MM-dd'T'HH:mm:ss.SSS'Z'
	Interval *string `json:"interval,omitempty"`


	// Metrics - A list of aggregated metrics
	Metrics *[]Queryresponsemetric `json:"metrics,omitempty"`

}

// String returns a JSON representation of the model
func (o *Queryresponsedata) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
