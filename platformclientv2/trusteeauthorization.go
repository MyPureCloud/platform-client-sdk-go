package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Trusteeauthorization
type Trusteeauthorization struct { 
	// Permissions - Permissions that the trustee user has in the trustor organization
	Permissions *[]string `json:"permissions,omitempty"`

}

func (u *Trusteeauthorization) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Trusteeauthorization

	

	return json.Marshal(&struct { 
		Permissions *[]string `json:"permissions,omitempty"`
		*Alias
	}{ 
		Permissions: u.Permissions,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Trusteeauthorization) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
