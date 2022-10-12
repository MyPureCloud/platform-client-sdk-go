package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Signature
type Signature struct { 
	// Enabled - A toggle to enable the signature on email send.
	Enabled *bool `json:"enabled,omitempty"`


	// CannedResponseId - The identifier referring to an email signature canned response.
	CannedResponseId *string `json:"cannedResponseId,omitempty"`


	// AlwaysIncluded - A toggle that defines if a signature is always included or only set on the first email in an email chain.
	AlwaysIncluded *bool `json:"alwaysIncluded,omitempty"`

}

func (o *Signature) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Signature
	
	return json.Marshal(&struct { 
		Enabled *bool `json:"enabled,omitempty"`
		
		CannedResponseId *string `json:"cannedResponseId,omitempty"`
		
		AlwaysIncluded *bool `json:"alwaysIncluded,omitempty"`
		*Alias
	}{ 
		Enabled: o.Enabled,
		
		CannedResponseId: o.CannedResponseId,
		
		AlwaysIncluded: o.AlwaysIncluded,
		Alias:    (*Alias)(o),
	})
}

func (o *Signature) UnmarshalJSON(b []byte) error {
	var SignatureMap map[string]interface{}
	err := json.Unmarshal(b, &SignatureMap)
	if err != nil {
		return err
	}
	
	if Enabled, ok := SignatureMap["enabled"].(bool); ok {
		o.Enabled = &Enabled
	}
    
	if CannedResponseId, ok := SignatureMap["cannedResponseId"].(string); ok {
		o.CannedResponseId = &CannedResponseId
	}
    
	if AlwaysIncluded, ok := SignatureMap["alwaysIncluded"].(bool); ok {
		o.AlwaysIncluded = &AlwaysIncluded
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Signature) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
