package platformclientv2
import (
	"encoding/json"
)

// Wfmbuscheduletopicmanagementunit
type Wfmbuscheduletopicmanagementunit struct { 
	// Id
	Id *string `json:"id,omitempty"`

}

// String returns a JSON representation of the model
func (o *Wfmbuscheduletopicmanagementunit) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
