package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Queueemailaddress
type Queueemailaddress struct { 
	// Domain
	Domain *Domainentityref `json:"domain,omitempty"`


	// Route
	Route *Inboundroute `json:"route,omitempty"`

}

func (u *Queueemailaddress) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Queueemailaddress

	

	return json.Marshal(&struct { 
		Domain *Domainentityref `json:"domain,omitempty"`
		
		Route *Inboundroute `json:"route,omitempty"`
		*Alias
	}{ 
		Domain: u.Domain,
		
		Route: u.Route,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Queueemailaddress) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
