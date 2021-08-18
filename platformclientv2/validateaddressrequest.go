package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Validateaddressrequest
type Validateaddressrequest struct { 
	// Address - Address schema
	Address *Streetaddress `json:"address,omitempty"`

}

func (u *Validateaddressrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Validateaddressrequest

	

	return json.Marshal(&struct { 
		Address *Streetaddress `json:"address,omitempty"`
		*Alias
	}{ 
		Address: u.Address,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Validateaddressrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
