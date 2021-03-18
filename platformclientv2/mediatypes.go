package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Mediatypes - Media types
type Mediatypes struct { 
	// Allow - Specify allowed media types for inbound and outbound messages. If this field is empty, all inbound and outbound media will be blocked.
	Allow *Mediatypeaccess `json:"allow,omitempty"`

}

// String returns a JSON representation of the model
func (o *Mediatypes) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
