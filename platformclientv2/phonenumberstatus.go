package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Phonenumberstatus
type Phonenumberstatus struct { 
	// Callable - Indicates whether or not a phone number is callable.
	Callable *bool `json:"callable,omitempty"`

}

func (u *Phonenumberstatus) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Phonenumberstatus

	

	return json.Marshal(&struct { 
		Callable *bool `json:"callable,omitempty"`
		*Alias
	}{ 
		Callable: u.Callable,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Phonenumberstatus) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
