package platformclientv2
import (
	"encoding/json"
)

// Wfmmoveagentscompletetopicmanagementunit
type Wfmmoveagentscompletetopicmanagementunit struct { 
	// Id
	Id *string `json:"id,omitempty"`

}

// String returns a JSON representation of the model
func (o *Wfmmoveagentscompletetopicmanagementunit) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
