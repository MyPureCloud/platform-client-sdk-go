package platformclientv2
import (
	"encoding/json"
)

// Wfmagentscheduleupdatetopicuserreference
type Wfmagentscheduleupdatetopicuserreference struct { 
	// Id
	Id *string `json:"id,omitempty"`

}

// String returns a JSON representation of the model
func (o *Wfmagentscheduleupdatetopicuserreference) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
