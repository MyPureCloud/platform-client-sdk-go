package platformclientv2
import (
	"encoding/json"
)

// Statisticalresponse
type Statisticalresponse struct { 
	// Interval
	Interval *string `json:"interval,omitempty"`


	// Metrics
	Metrics *[]Aggregatemetricdata `json:"metrics,omitempty"`


	// Views
	Views *[]Aggregateviewdata `json:"views,omitempty"`

}

// String returns a JSON representation of the model
func (o *Statisticalresponse) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
