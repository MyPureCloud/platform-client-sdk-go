package platformclientv2
import (
	"encoding/json"
)

// Edgeversioninformation
type Edgeversioninformation struct { 
	// SoftwareVersion
	SoftwareVersion *string `json:"softwareVersion,omitempty"`

}

// String returns a JSON representation of the model
func (o *Edgeversioninformation) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
