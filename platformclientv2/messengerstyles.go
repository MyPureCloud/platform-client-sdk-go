package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Messengerstyles
type Messengerstyles struct { 
	// PrimaryColor - The primary color of messenger in hexadecimal
	PrimaryColor *string `json:"primaryColor,omitempty"`

}

func (u *Messengerstyles) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Messengerstyles

	

	return json.Marshal(&struct { 
		PrimaryColor *string `json:"primaryColor,omitempty"`
		*Alias
	}{ 
		PrimaryColor: u.PrimaryColor,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Messengerstyles) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
