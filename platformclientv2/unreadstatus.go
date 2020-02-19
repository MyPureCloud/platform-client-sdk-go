package platformclientv2
import (
	"encoding/json"
)

// Unreadstatus
type Unreadstatus struct { 
	// Unread - Sets if the alert is read or unread.
	Unread *bool `json:"unread,omitempty"`

}

// String returns a JSON representation of the model
func (o *Unreadstatus) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
