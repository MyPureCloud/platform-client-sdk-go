package platformclientv2
import (
	"encoding/json"
)

// Trunkinstancetopictrunkerrorinfodetails
type Trunkinstancetopictrunkerrorinfodetails struct { 
	// Code
	Code *string `json:"code,omitempty"`


	// Message
	Message *string `json:"message,omitempty"`


	// Hostname
	Hostname *string `json:"hostname,omitempty"`

}

// String returns a JSON representation of the model
func (o *Trunkinstancetopictrunkerrorinfodetails) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
