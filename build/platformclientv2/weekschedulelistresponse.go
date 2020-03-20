package platformclientv2
import (
	"encoding/json"
)

// Weekschedulelistresponse - Week schedule list
type Weekschedulelistresponse struct { 
	// Entities
	Entities *[]Weekschedulelistitemresponse `json:"entities,omitempty"`

}

// String returns a JSON representation of the model
func (o *Weekschedulelistresponse) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
