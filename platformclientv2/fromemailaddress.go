package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Fromemailaddress
type Fromemailaddress struct { 
	// Domain - The OutboundDomain used for the email address.
	Domain *Domainentityref `json:"domain,omitempty"`


	// FriendlyName - The friendly name of the email address.
	FriendlyName *string `json:"friendlyName,omitempty"`


	// LocalPart - The local part of the email address.
	LocalPart *string `json:"localPart,omitempty"`

}

// String returns a JSON representation of the model
func (o *Fromemailaddress) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
