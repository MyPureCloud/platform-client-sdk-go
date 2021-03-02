package platformclientv2
import (
	"encoding/json"
)

// Wfmbuscheduletopicuserreference
type Wfmbuscheduletopicuserreference struct { 
	// Id
	Id *string `json:"id,omitempty"`

}

// String returns a JSON representation of the model
func (o *Wfmbuscheduletopicuserreference) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
