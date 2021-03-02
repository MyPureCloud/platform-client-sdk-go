package platformclientv2
import (
	"encoding/json"
)

// Userobservationqueryresponse
type Userobservationqueryresponse struct { 
	// Results
	Results *[]Userobservationdatacontainer `json:"results,omitempty"`

}

// String returns a JSON representation of the model
func (o *Userobservationqueryresponse) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
