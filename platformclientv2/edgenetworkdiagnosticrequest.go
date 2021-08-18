package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Edgenetworkdiagnosticrequest
type Edgenetworkdiagnosticrequest struct { 
	// Host - IPv4/6 address or host to be probed for connectivity. No port allowed.
	Host *string `json:"host,omitempty"`

}

func (u *Edgenetworkdiagnosticrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Edgenetworkdiagnosticrequest

	

	return json.Marshal(&struct { 
		Host *string `json:"host,omitempty"`
		*Alias
	}{ 
		Host: u.Host,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Edgenetworkdiagnosticrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
