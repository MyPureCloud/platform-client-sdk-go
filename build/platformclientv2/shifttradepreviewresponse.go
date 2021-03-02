package platformclientv2
import (
	"encoding/json"
)

// Shifttradepreviewresponse
type Shifttradepreviewresponse struct { 
	// Activities - List of activities that will make up the new shift if this shift trade is approved
	Activities *[]Shifttradeactivitypreviewresponse `json:"activities,omitempty"`

}

// String returns a JSON representation of the model
func (o *Shifttradepreviewresponse) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
