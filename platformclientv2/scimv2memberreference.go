package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Scimv2memberreference - Defines a reference to SCIM group members.
type Scimv2memberreference struct { 
	// VarType - The type of SCIM resource.
	VarType *string `json:"type,omitempty"`


	// Value - The ID of the group member. Can be \"userId\" or \"groupId\".
	Value *string `json:"value,omitempty"`


	// Ref - The reference URI of the SCIM resource.
	Ref *string `json:"$ref,omitempty"`

}

func (o *Scimv2memberreference) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Scimv2memberreference
	
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

func (o *Scimv2memberreference) UnmarshalJSON(b []byte) error {
	var Scimv2memberreferenceMap map[string]interface{}
	err := json.Unmarshal(b, &Scimv2memberreferenceMap)
	if err != nil {
		return err
	}
	
	if VarType, ok := Scimv2memberreferenceMap["type"].(string); ok {
		o.VarType = &VarType
	}
	
	if Value, ok := Scimv2memberreferenceMap["value"].(string); ok {
		o.Value = &Value
	}
	
	if Ref, ok := Scimv2memberreferenceMap["$ref"].(string); ok {
		o.Ref = &Ref
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Scimv2memberreference) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
