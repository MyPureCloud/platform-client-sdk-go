package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Smsavailablephonenumberentitylisting
type Smsavailablephonenumberentitylisting struct { 
	// Entities
	Entities *[]Smsavailablephonenumber `json:"entities,omitempty"`

}

func (u *Smsavailablephonenumberentitylisting) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Smsavailablephonenumberentitylisting

	

	return json.Marshal(&struct { 
		Entities *[]Smsavailablephonenumber `json:"entities,omitempty"`
		*Alias
	}{ 
		Entities: u.Entities,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Smsavailablephonenumberentitylisting) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
