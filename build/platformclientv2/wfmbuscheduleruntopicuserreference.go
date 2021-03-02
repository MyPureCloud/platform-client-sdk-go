package platformclientv2
import (
	"encoding/json"
)

// Wfmbuscheduleruntopicuserreference
type Wfmbuscheduleruntopicuserreference struct { 
	// Id
	Id *string `json:"id,omitempty"`

}

// String returns a JSON representation of the model
func (o *Wfmbuscheduleruntopicuserreference) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
