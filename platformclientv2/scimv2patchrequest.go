package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Scimv2patchrequest - Defines a SCIM PATCH request. See section 3.5.2 \"Modifying with PATCH\" in RFC 7644 for details.
type Scimv2patchrequest struct { 
	// Schemas - The list of schemas used in the PATCH request.
	Schemas *[]string `json:"schemas,omitempty"`


	// Operations - The list of operations to perform for the PATCH request.
	Operations *[]Scimv2patchoperation `json:"Operations,omitempty"`

}

func (o *Scimv2patchrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Scimv2patchrequest
	
	return json.Marshal(&struct { 
		Schemas *[]string `json:"schemas,omitempty"`
		
		Operations *[]Scimv2patchoperation `json:"Operations,omitempty"`
		*Alias
	}{ 
		Schemas: o.Schemas,
		
		Operations: o.Operations,
		Alias:    (*Alias)(o),
	})
}

func (o *Scimv2patchrequest) UnmarshalJSON(b []byte) error {
	var Scimv2patchrequestMap map[string]interface{}
	err := json.Unmarshal(b, &Scimv2patchrequestMap)
	if err != nil {
		return err
	}
	
	if Schemas, ok := Scimv2patchrequestMap["schemas"].([]interface{}); ok {
		SchemasString, _ := json.Marshal(Schemas)
		json.Unmarshal(SchemasString, &o.Schemas)
	}
	
	if Operations, ok := Scimv2patchrequestMap["Operations"].([]interface{}); ok {
		OperationsString, _ := json.Marshal(Operations)
		json.Unmarshal(OperationsString, &o.Operations)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Scimv2patchrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
