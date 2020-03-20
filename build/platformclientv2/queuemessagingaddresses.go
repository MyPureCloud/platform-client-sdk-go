package platformclientv2
import (
	"encoding/json"
)

// Queuemessagingaddresses
type Queuemessagingaddresses struct { 
	// SmsAddress
	SmsAddress *Domainentityref `json:"smsAddress,omitempty"`

}

// String returns a JSON representation of the model
func (o *Queuemessagingaddresses) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
