package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Check
type Check struct { 
	// Result - The result of a check executed. This indicates if the check was successful or not.
	Result *string `json:"result,omitempty"`


	// VarType - The type of check executed.
	VarType *string `json:"type,omitempty"`

}

func (u *Check) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Check

	

	return json.Marshal(&struct { 
		Result *string `json:"result,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		*Alias
	}{ 
		Result: u.Result,
		
		VarType: u.VarType,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Check) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
