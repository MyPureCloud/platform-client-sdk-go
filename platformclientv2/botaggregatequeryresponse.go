package platformclientv2
import (
	"encoding/json"
)

// Botaggregatequeryresponse
type Botaggregatequeryresponse struct { 
	// Results
	Results *[]Botaggregatedatacontainer `json:"results,omitempty"`

}

// String returns a JSON representation of the model
func (o *Botaggregatequeryresponse) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
