package platformclientv2
import (
	"encoding/json"
)

// Notificationtemplatefooter - Template footer object
type Notificationtemplatefooter struct { 
	// Text - Footer text. For WhatsApp, ignored
	Text *string `json:"text,omitempty"`

}

// String returns a JSON representation of the model
func (o *Notificationtemplatefooter) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
