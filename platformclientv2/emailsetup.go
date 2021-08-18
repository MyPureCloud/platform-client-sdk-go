package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Emailsetup
type Emailsetup struct { 
	// RootDomain - The root PureCloud domain that all sub-domains are created from.
	RootDomain *string `json:"rootDomain,omitempty"`

}

func (u *Emailsetup) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Emailsetup

	

	return json.Marshal(&struct { 
		RootDomain *string `json:"rootDomain,omitempty"`
		*Alias
	}{ 
		RootDomain: u.RootDomain,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Emailsetup) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
