package platformclientv2
import (
	"encoding/json"
)

// Trunkmetricsnetworktypeip
type Trunkmetricsnetworktypeip struct { 
	// Address - Assigned IP Address for the interface
	Address *string `json:"address,omitempty"`


	// ErrorInfo - Information about the error.
	ErrorInfo *Trunkerrorinfo `json:"errorInfo,omitempty"`

}

// String returns a JSON representation of the model
func (o *Trunkmetricsnetworktypeip) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
