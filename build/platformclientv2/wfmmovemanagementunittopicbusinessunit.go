package platformclientv2
import (
	"encoding/json"
)

// Wfmmovemanagementunittopicbusinessunit
type Wfmmovemanagementunittopicbusinessunit struct { 
	// Id
	Id *string `json:"id,omitempty"`

}

// String returns a JSON representation of the model
func (o *Wfmmovemanagementunittopicbusinessunit) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
