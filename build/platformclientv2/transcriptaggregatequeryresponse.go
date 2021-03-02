package platformclientv2
import (
	"encoding/json"
)

// Transcriptaggregatequeryresponse
type Transcriptaggregatequeryresponse struct { 
	// Results
	Results *[]Transcriptaggregatedatacontainer `json:"results,omitempty"`

}

// String returns a JSON representation of the model
func (o *Transcriptaggregatequeryresponse) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
