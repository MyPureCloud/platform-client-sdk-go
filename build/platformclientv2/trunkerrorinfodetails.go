package platformclientv2
import (
	"encoding/json"
)

// Trunkerrorinfodetails
type Trunkerrorinfodetails struct { 
	// Code
	Code *string `json:"code,omitempty"`


	// Message
	Message *string `json:"message,omitempty"`


	// Hostname
	Hostname *string `json:"hostname,omitempty"`

}

// String returns a JSON representation of the model
func (o *Trunkerrorinfodetails) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
