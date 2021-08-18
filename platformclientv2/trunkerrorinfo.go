package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Trunkerrorinfo
type Trunkerrorinfo struct { 
	// Text
	Text *string `json:"text,omitempty"`


	// Code
	Code *string `json:"code,omitempty"`


	// Details
	Details *Trunkerrorinfodetails `json:"details,omitempty"`

}

func (u *Trunkerrorinfo) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Trunkerrorinfo

	

	return json.Marshal(&struct { 
		Text *string `json:"text,omitempty"`
		
		Code *string `json:"code,omitempty"`
		
		Details *Trunkerrorinfodetails `json:"details,omitempty"`
		*Alias
	}{ 
		Text: u.Text,
		
		Code: u.Code,
		
		Details: u.Details,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Trunkerrorinfo) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
