package platformclientv2
import (
	"encoding/json"
)

// Chatsettings
type Chatsettings struct { 
	// MessageRetentionPeriodDays - Retention time for messages in days
	MessageRetentionPeriodDays *int `json:"messageRetentionPeriodDays,omitempty"`

}

// String returns a JSON representation of the model
func (o *Chatsettings) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
