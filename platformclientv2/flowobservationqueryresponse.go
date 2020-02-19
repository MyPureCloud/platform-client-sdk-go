package platformclientv2
import (
	"encoding/json"
)

// Flowobservationqueryresponse
type Flowobservationqueryresponse struct { 
	// Results
	Results *[]Flowobservationdatacontainer `json:"results,omitempty"`

}

// String returns a JSON representation of the model
func (o *Flowobservationqueryresponse) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
