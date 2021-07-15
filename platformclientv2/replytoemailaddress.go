package platformclientv2
import (
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

// String returns a JSON representation of the model
func (o *Replytoemailaddress) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
