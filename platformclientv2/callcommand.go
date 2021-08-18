package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Callcommand
type Callcommand struct { 
	// CallNumber - The phone number to dial for this call.
	CallNumber *string `json:"callNumber,omitempty"`


	// PhoneColumn - For a dialer preview or scheduled callback, the phone column associated with the phone number
	PhoneColumn *string `json:"phoneColumn,omitempty"`

}

func (u *Callcommand) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Callcommand

	

	return json.Marshal(&struct { 
		CallNumber *string `json:"callNumber,omitempty"`
		
		PhoneColumn *string `json:"phoneColumn,omitempty"`
		*Alias
	}{ 
		CallNumber: u.CallNumber,
		
		PhoneColumn: u.PhoneColumn,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Callcommand) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
