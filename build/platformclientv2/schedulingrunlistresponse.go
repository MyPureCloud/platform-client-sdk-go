package platformclientv2
import (
	"encoding/json"
)

// Schedulingrunlistresponse
type Schedulingrunlistresponse struct { 
	// Entities
	Entities *[]Schedulingrunresponse `json:"entities,omitempty"`

}

// String returns a JSON representation of the model
func (o *Schedulingrunlistresponse) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
