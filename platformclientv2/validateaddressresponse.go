package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Validateaddressresponse
type Validateaddressresponse struct { 
	// Valid - Was the passed in address valid
	Valid *bool `json:"valid,omitempty"`


	// Response - Subscriber schema
	Response *Subscriberresponse `json:"response,omitempty"`

}

func (u *Validateaddressresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Validateaddressresponse

	

	return json.Marshal(&struct { 
		Valid *bool `json:"valid,omitempty"`
		
		Response *Subscriberresponse `json:"response,omitempty"`
		*Alias
	}{ 
		Valid: u.Valid,
		
		Response: u.Response,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Validateaddressresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
