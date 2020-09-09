package platformclientv2
import (
	"encoding/json"
)

// Coachingappointmentaggregateresponse
type Coachingappointmentaggregateresponse struct { 
	// Results - The results of the query
	Results *[]Queryresponsegroupeddata `json:"results,omitempty"`

}

// String returns a JSON representation of the model
func (o *Coachingappointmentaggregateresponse) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
