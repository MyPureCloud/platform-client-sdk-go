package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Auditqueryservicemapping
type Auditqueryservicemapping struct { 
	// Services - List of Services
	Services *[]Auditqueryservice `json:"services,omitempty"`

}

func (u *Auditqueryservicemapping) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Auditqueryservicemapping

	

	return json.Marshal(&struct { 
		Services *[]Auditqueryservice `json:"services,omitempty"`
		*Alias
	}{ 
		Services: u.Services,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Auditqueryservicemapping) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
