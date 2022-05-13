package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Scimv2groupreference - Defines a reference to SCIM groups.
type Scimv2groupreference struct { 
	// VarType - The type of SCIM resource.
	VarType *string `json:"type,omitempty"`


	// Value - The ID of the group member. Can be \"userId\" or \"groupId\".
	Value *string `json:"value,omitempty"`


	// Ref - The reference URI of the SCIM resource.
	Ref *string `json:"$ref,omitempty"`

}

func (o *Scimv2groupreference) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Scimv2groupreference
	
	return json.Marshal(&struct { 
		VarType *string `json:"type,omitempty"`
		
		Value *string `json:"value,omitempty"`
		
		Ref *string `json:"$ref,omitempty"`
		*Alias
	}{ 
		VarType: o.VarType,
		
		Value: o.Value,
		
		Ref: o.Ref,
		Alias:    (*Alias)(o),
	})
}

func (o *Scimv2groupreference) UnmarshalJSON(b []byte) error {
	var Scimv2groupreferenceMap map[string]interface{}
	err := json.Unmarshal(b, &Scimv2groupreferenceMap)
	if err != nil {
		return err
	}
	
	if VarType, ok := Scimv2groupreferenceMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if Value, ok := Scimv2groupreferenceMap["value"].(string); ok {
		o.Value = &Value
	}
    
	if Ref, ok := Scimv2groupreferenceMap["$ref"].(string); ok {
		o.Ref = &Ref
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Scimv2groupreference) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
