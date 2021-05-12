package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Trusteeauthorization
type Trusteeauthorization struct { 
	// Permissions - Permissions that the trustee user has in the trustor organization
	Permissions *[]string `json:"permissions,omitempty"`

}

// String returns a JSON representation of the model
func (o *Trusteeauthorization) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
