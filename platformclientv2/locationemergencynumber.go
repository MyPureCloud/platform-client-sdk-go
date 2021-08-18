package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Locationemergencynumber
type Locationemergencynumber struct { 
	// E164
	E164 *string `json:"e164,omitempty"`


	// Number
	Number *string `json:"number,omitempty"`


	// VarType - The type of emergency number.
	VarType *string `json:"type,omitempty"`

}

func (u *Locationemergencynumber) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Locationemergencynumber

	

	return json.Marshal(&struct { 
		E164 *string `json:"e164,omitempty"`
		
		Number *string `json:"number,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		*Alias
	}{ 
		E164: u.E164,
		
		Number: u.Number,
		
		VarType: u.VarType,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Locationemergencynumber) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
