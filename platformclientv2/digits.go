package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Digits
type Digits struct { 
	// Digits - A string representing the digits pressed on phone.
	Digits *string `json:"digits,omitempty"`

}

func (u *Digits) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Digits

	

	return json.Marshal(&struct { 
		Digits *string `json:"digits,omitempty"`
		*Alias
	}{ 
		Digits: u.Digits,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Digits) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
