package platformclientv2
import (
	"encoding/json"
)

// Updatenotificationsrequest
type Updatenotificationsrequest struct { 
	// Entities - The notifications to update
	Entities *[]Wfmusernotification `json:"entities,omitempty"`

}

// String returns a JSON representation of the model
func (o *Updatenotificationsrequest) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
