package platformclientv2
import (
	"encoding/json"
)

// Errorinfo
type Errorinfo struct { 
	// Message
	Message *string `json:"message,omitempty"`


	// Code
	Code *string `json:"code,omitempty"`

}

// String returns a JSON representation of the model
func (o *Errorinfo) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
