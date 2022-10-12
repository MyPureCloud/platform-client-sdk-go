package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Identifierclaimrequest
type Identifierclaimrequest struct { 
	// Operation - The operation to perform claim/release
	Operation *string `json:"operation,omitempty"`


	// Identifier - The identifier that should be claimed/released from a contact
	Identifier *Contactidentifier `json:"identifier,omitempty"`

}

func (o *Identifierclaimrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Identifierclaimrequest
	
	return json.Marshal(&struct { 
		Operation *string `json:"operation,omitempty"`
		
		Identifier *Contactidentifier `json:"identifier,omitempty"`
		*Alias
	}{ 
		Operation: o.Operation,
		
		Identifier: o.Identifier,
		Alias:    (*Alias)(o),
	})
}

func (o *Identifierclaimrequest) UnmarshalJSON(b []byte) error {
	var IdentifierclaimrequestMap map[string]interface{}
	err := json.Unmarshal(b, &IdentifierclaimrequestMap)
	if err != nil {
		return err
	}
	
	if Operation, ok := IdentifierclaimrequestMap["operation"].(string); ok {
		o.Operation = &Operation
	}
    
	if Identifier, ok := IdentifierclaimrequestMap["identifier"].(map[string]interface{}); ok {
		IdentifierString, _ := json.Marshal(Identifier)
		json.Unmarshal(IdentifierString, &o.Identifier)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Identifierclaimrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
