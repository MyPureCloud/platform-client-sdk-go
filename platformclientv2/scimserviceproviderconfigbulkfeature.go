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

func (o *Scimserviceproviderconfigbulkfeature) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Scimserviceproviderconfigbulkfeature
	
	return json.Marshal(&struct { 
		Supported *bool `json:"supported,omitempty"`
		
		MaxOperations *int `json:"maxOperations,omitempty"`
		
		MaxPayloadSize *int `json:"maxPayloadSize,omitempty"`
		*Alias
	}{ 
		Supported: o.Supported,
		
		MaxOperations: o.MaxOperations,
		
		MaxPayloadSize: o.MaxPayloadSize,
		Alias:    (*Alias)(o),
	})
}

func (o *Scimserviceproviderconfigbulkfeature) UnmarshalJSON(b []byte) error {
	var ScimserviceproviderconfigbulkfeatureMap map[string]interface{}
	err := json.Unmarshal(b, &ScimserviceproviderconfigbulkfeatureMap)
	if err != nil {
		return err
	}
	
	if Supported, ok := ScimserviceproviderconfigbulkfeatureMap["supported"].(bool); ok {
		o.Supported = &Supported
	}
    
	if MaxOperations, ok := ScimserviceproviderconfigbulkfeatureMap["maxOperations"].(float64); ok {
		MaxOperationsInt := int(MaxOperations)
		o.MaxOperations = &MaxOperationsInt
	}
	
	if MaxPayloadSize, ok := ScimserviceproviderconfigbulkfeatureMap["maxPayloadSize"].(float64); ok {
		MaxPayloadSizeInt := int(MaxPayloadSize)
		o.MaxPayloadSize = &MaxPayloadSizeInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Scimserviceproviderconfigbulkfeature) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
