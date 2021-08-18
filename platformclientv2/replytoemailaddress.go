package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Replytoemailaddress
type Replytoemailaddress struct { 
	// Domain - The InboundDomain used for the email address.
	Domain *Domainentityref `json:"domain,omitempty"`


	// Route - The InboundRoute used for the email address.
	Route *Domainentityref `json:"route,omitempty"`

}

func (u *Replytoemailaddress) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Replytoemailaddress

	

	return json.Marshal(&struct { 
		Domain *Domainentityref `json:"domain,omitempty"`
		
		Route *Domainentityref `json:"route,omitempty"`
		*Alias
	}{ 
		Domain: u.Domain,
		
		Route: u.Route,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Replytoemailaddress) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
