package platformclientv2
import (
	"github.com/leekchan/timeutil"
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

func (u *Metadata) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Metadata

	

	return json.Marshal(&struct { 
		PairingToken *string `json:"pairing-token,omitempty"`
		
		PairingTrust *[]string `json:"pairing-trust,omitempty"`
		
		PairingUrl *string `json:"pairing-url,omitempty"`
		*Alias
	}{ 
		PairingToken: u.PairingToken,
		
		PairingTrust: u.PairingTrust,
		
		PairingUrl: u.PairingUrl,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Metadata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
