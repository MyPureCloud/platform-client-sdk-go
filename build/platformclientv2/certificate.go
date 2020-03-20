package platformclientv2
import (
	"encoding/json"
)

// Certificate - Represents a certificate to parse.
type Certificate struct { 
	// Certificate - The certificate to parse.
	Certificate *string `json:"certificate,omitempty"`

}

// String returns a JSON representation of the model
func (o *Certificate) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
