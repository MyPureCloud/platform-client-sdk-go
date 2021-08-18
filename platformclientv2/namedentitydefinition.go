package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Namedentitydefinition
type Namedentitydefinition struct { 
	// Name - The name of the entity.
	Name *string `json:"name,omitempty"`


	// VarType - The name of the entity type.
	VarType *string `json:"type,omitempty"`

}

func (u *Namedentitydefinition) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Namedentitydefinition

	

	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		*Alias
	}{ 
		Name: u.Name,
		
		VarType: u.VarType,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Namedentitydefinition) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
