package platformclientv2
import (
	"encoding/json"
)

// Queueusereventtopicuserreference
type Queueusereventtopicuserreference struct { 
	// Id
	Id *string `json:"id,omitempty"`

}

// String returns a JSON representation of the model
func (o *Queueusereventtopicuserreference) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
