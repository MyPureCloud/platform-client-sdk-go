package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Mediatypes - Media types
type Mediatypes struct { 
	// Allow - Specify allowed media types for inbound and outbound messages. If this field is empty, all inbound and outbound media will be blocked.
	Allow *Mediatypeaccess `json:"allow,omitempty"`

}

func (u *Mediatypes) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Mediatypes

	

	return json.Marshal(&struct { 
		Allow *Mediatypeaccess `json:"allow,omitempty"`
		*Alias
	}{ 
		Allow: u.Allow,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Mediatypes) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
