package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Replacementterm
type Replacementterm struct { 
	// VarType
	VarType *string `json:"type,omitempty"`


	// ExistingValue
	ExistingValue *string `json:"existingValue,omitempty"`


	// UpdatedValue
	UpdatedValue *string `json:"updatedValue,omitempty"`

}

func (u *Replacementterm) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Replacementterm

	

	return json.Marshal(&struct { 
		VarType *string `json:"type,omitempty"`
		
		ExistingValue *string `json:"existingValue,omitempty"`
		
		UpdatedValue *string `json:"updatedValue,omitempty"`
		*Alias
	}{ 
		VarType: u.VarType,
		
		ExistingValue: u.ExistingValue,
		
		UpdatedValue: u.UpdatedValue,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Replacementterm) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
