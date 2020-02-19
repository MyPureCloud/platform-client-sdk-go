package platformclientv2
import (
	"encoding/json"
)

// Notificationsresponse
type Notificationsresponse struct { 
	// Entities
	Entities *[]Wfmusernotification `json:"entities,omitempty"`

}

// String returns a JSON representation of the model
func (o *Notificationsresponse) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
