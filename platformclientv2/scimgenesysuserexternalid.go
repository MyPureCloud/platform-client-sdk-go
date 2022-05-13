package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Scimgenesysuserexternalid - External Identifiers of user. The external identifier must be unique within the organization and the 'authority'
type Scimgenesysuserexternalid struct { 
	// Authority - Authority, or scope, of \"externalId\". Allows multiple external identifiers to be defined. Represents the source of the external identifier.
	Authority *string `json:"authority,omitempty"`


	// Value - Identifier of the user in an external system.
	Value *string `json:"value,omitempty"`

}

func (o *Scimgenesysuserexternalid) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Scimgenesysuserexternalid
	
	return json.Marshal(&struct { 
		Authority *string `json:"authority,omitempty"`
		
		Value *string `json:"value,omitempty"`
		*Alias
	}{ 
		Authority: o.Authority,
		
		Value: o.Value,
		Alias:    (*Alias)(o),
	})
}

func (o *Scimgenesysuserexternalid) UnmarshalJSON(b []byte) error {
	var ScimgenesysuserexternalidMap map[string]interface{}
	err := json.Unmarshal(b, &ScimgenesysuserexternalidMap)
	if err != nil {
		return err
	}
	
	if Authority, ok := ScimgenesysuserexternalidMap["authority"].(string); ok {
		o.Authority = &Authority
	}
    
	if Value, ok := ScimgenesysuserexternalidMap["value"].(string); ok {
		o.Value = &Value
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Scimgenesysuserexternalid) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
