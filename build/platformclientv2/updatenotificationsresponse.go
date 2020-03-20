package platformclientv2
import (
	"encoding/json"
)

// Updatenotificationsresponse
type Updatenotificationsresponse struct { 
	// Entities
	Entities *[]Updatenotificationresponse `json:"entities,omitempty"`

}

// String returns a JSON representation of the model
func (o *Updatenotificationsresponse) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
