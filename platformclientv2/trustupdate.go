package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Trustupdate
type Trustupdate struct { 
	// Enabled - If disabled no trustee user will have access, even if they were previously added.
	Enabled *bool `json:"enabled,omitempty"`


	// DateExpired - The expiration date of the trust. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateExpired *time.Time `json:"dateExpired,omitempty"`

}

func (u *Trustupdate) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Trustupdate

	
	DateExpired := new(string)
	if u.DateExpired != nil {
		
		*DateExpired = timeutil.Strftime(u.DateExpired, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateExpired = nil
	}
	

	return json.Marshal(&struct { 
		Enabled *bool `json:"enabled,omitempty"`
		
		DateExpired *string `json:"dateExpired,omitempty"`
		*Alias
	}{ 
		Enabled: u.Enabled,
		
		DateExpired: DateExpired,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Trustupdate) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
