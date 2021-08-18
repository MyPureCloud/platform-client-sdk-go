package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Certificate - Represents a certificate to parse.
type Certificate struct { 
	// Certificate - The certificate to parse.
	Certificate *string `json:"certificate,omitempty"`

}

func (u *Certificate) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Certificate

	

	return json.Marshal(&struct { 
		Certificate *string `json:"certificate,omitempty"`
		*Alias
	}{ 
		Certificate: u.Certificate,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Certificate) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
