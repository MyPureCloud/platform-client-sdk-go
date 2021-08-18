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

func (u *Scimv2groupreference) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Scimv2groupreference

	

	return json.Marshal(&struct { 
		VarType *string `json:"type,omitempty"`
		
		Value *string `json:"value,omitempty"`
		
		Ref *string `json:"$ref,omitempty"`
		*Alias
	}{ 
		VarType: u.VarType,
		
		Value: u.Value,
		
		Ref: u.Ref,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Scimv2groupreference) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
