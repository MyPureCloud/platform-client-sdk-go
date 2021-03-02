package platformclientv2
import (
	"encoding/json"
)

// Metadata
type Metadata struct { 
	// PairingToken
	PairingToken *string `json:"pairing-token,omitempty"`


	// PairingTrust
	PairingTrust *[]string `json:"pairing-trust,omitempty"`


	// PairingUrl
	PairingUrl *string `json:"pairing-url,omitempty"`

}

// String returns a JSON representation of the model
func (o *Metadata) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
