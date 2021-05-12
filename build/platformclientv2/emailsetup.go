package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Emailsetup
type Emailsetup struct { 
	// RootDomain - The root PureCloud domain that all sub-domains are created from.
	RootDomain *string `json:"rootDomain,omitempty"`

}

// String returns a JSON representation of the model
func (o *Emailsetup) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
