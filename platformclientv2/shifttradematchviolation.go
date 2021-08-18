package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Shifttradematchviolation
type Shifttradematchviolation struct { 
	// VarType - The type of constraint violation
	VarType *string `json:"type,omitempty"`


	// Params - Clarifying user params for constructing helpful error messages
	Params *map[string]string `json:"params,omitempty"`

}

func (u *Shifttradematchviolation) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Shifttradematchviolation

	

	return json.Marshal(&struct { 
		VarType *string `json:"type,omitempty"`
		
		Params *map[string]string `json:"params,omitempty"`
		*Alias
	}{ 
		VarType: u.VarType,
		
		Params: u.Params,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Shifttradematchviolation) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
