package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Items
type Items struct { 
	// VarType
	VarType *string `json:"type,omitempty"`


	// Pattern
	Pattern *string `json:"pattern,omitempty"`

}

func (u *Items) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Items

	

	return json.Marshal(&struct { 
		VarType *string `json:"type,omitempty"`
		
		Pattern *string `json:"pattern,omitempty"`
		*Alias
	}{ 
		VarType: u.VarType,
		
		Pattern: u.Pattern,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Items) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
