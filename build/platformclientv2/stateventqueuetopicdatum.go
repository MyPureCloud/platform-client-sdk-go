package platformclientv2
import (
	"encoding/json"
)

// Stateventqueuetopicdatum
type Stateventqueuetopicdatum struct { 
	// Interval
	Interval *string `json:"interval,omitempty"`


	// Metrics
	Metrics *[]Stateventqueuetopicmetric `json:"metrics,omitempty"`

}

// String returns a JSON representation of the model
func (o *Stateventqueuetopicdatum) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
