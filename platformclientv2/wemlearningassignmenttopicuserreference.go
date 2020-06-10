package platformclientv2
import (
	"encoding/json"
)

// Wemlearningassignmenttopicuserreference
type Wemlearningassignmenttopicuserreference struct { 
	// Id
	Id *string `json:"id,omitempty"`

}

// String returns a JSON representation of the model
func (o *Wemlearningassignmenttopicuserreference) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
