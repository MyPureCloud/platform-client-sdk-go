package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Ipaddressauthentication
type Ipaddressauthentication struct { 
	// NetworkWhitelist
	NetworkWhitelist *[]string `json:"networkWhitelist,omitempty"`

}

// String returns a JSON representation of the model
func (o *Ipaddressauthentication) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
