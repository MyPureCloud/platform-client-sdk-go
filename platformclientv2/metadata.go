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

func (o *Metadata) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Metadata
	
	return json.Marshal(&struct { 
		PairingToken *string `json:"pairing-token,omitempty"`
		
		PairingTrust *[]string `json:"pairing-trust,omitempty"`
		
		PairingUrl *string `json:"pairing-url,omitempty"`
		*Alias
	}{ 
		PairingToken: o.PairingToken,
		
		PairingTrust: o.PairingTrust,
		
		PairingUrl: o.PairingUrl,
		Alias:    (*Alias)(o),
	})
}

func (o *Metadata) UnmarshalJSON(b []byte) error {
	var MetadataMap map[string]interface{}
	err := json.Unmarshal(b, &MetadataMap)
	if err != nil {
		return err
	}
	
	if PairingToken, ok := MetadataMap["pairing-token"].(string); ok {
		o.PairingToken = &PairingToken
	}
    
	if PairingTrust, ok := MetadataMap["pairing-trust"].([]interface{}); ok {
		PairingTrustString, _ := json.Marshal(PairingTrust)
		json.Unmarshal(PairingTrustString, &o.PairingTrust)
	}
	
	if PairingUrl, ok := MetadataMap["pairing-url"].(string); ok {
		o.PairingUrl = &PairingUrl
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Metadata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
