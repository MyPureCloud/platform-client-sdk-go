package platformclientv2
import (
	"encoding/json"
)

// Wfmmoveagentscompletetopicuserreference
type Wfmmoveagentscompletetopicuserreference struct { 
	// Id
	Id *string `json:"id,omitempty"`

}

// String returns a JSON representation of the model
func (o *Wfmmoveagentscompletetopicuserreference) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
