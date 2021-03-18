package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
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
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
