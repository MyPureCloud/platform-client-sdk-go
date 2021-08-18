package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Expansioncriterium
type Expansioncriterium struct { 
	// VarType
	VarType *string `json:"type,omitempty"`


	// Threshold
	Threshold *float64 `json:"threshold,omitempty"`

}

func (u *Expansioncriterium) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Expansioncriterium

	

	return json.Marshal(&struct { 
		VarType *string `json:"type,omitempty"`
		
		Threshold *float64 `json:"threshold,omitempty"`
		*Alias
	}{ 
		VarType: u.VarType,
		
		Threshold: u.Threshold,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Expansioncriterium) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
