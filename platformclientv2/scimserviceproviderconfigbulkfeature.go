package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Scimserviceproviderconfigbulkfeature - Defines a \"bulk\" request in the SCIM service provider's configuration.
type Scimserviceproviderconfigbulkfeature struct { 
	// Supported - Indicates whether configuration options are supported.
	Supported *bool `json:"supported,omitempty"`


	// MaxOperations - The maximum number of operations for each bulk request.
	MaxOperations *int `json:"maxOperations,omitempty"`


	// MaxPayloadSize - The maximum payload size.
	MaxPayloadSize *int `json:"maxPayloadSize,omitempty"`

}

func (u *Scimserviceproviderconfigbulkfeature) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Scimserviceproviderconfigbulkfeature

	

	return json.Marshal(&struct { 
		Supported *bool `json:"supported,omitempty"`
		
		MaxOperations *int `json:"maxOperations,omitempty"`
		
		MaxPayloadSize *int `json:"maxPayloadSize,omitempty"`
		*Alias
	}{ 
		Supported: u.Supported,
		
		MaxOperations: u.MaxOperations,
		
		MaxPayloadSize: u.MaxPayloadSize,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Scimserviceproviderconfigbulkfeature) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
