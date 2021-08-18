package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Trunkinstancetopictrunkerrorinfo
type Trunkinstancetopictrunkerrorinfo struct { 
	// Text
	Text *string `json:"text,omitempty"`


	// Code
	Code *string `json:"code,omitempty"`


	// Details
	Details *Trunkinstancetopictrunkerrorinfodetails `json:"details,omitempty"`

}

func (u *Trunkinstancetopictrunkerrorinfo) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Trunkinstancetopictrunkerrorinfo

	

	return json.Marshal(&struct { 
		Text *string `json:"text,omitempty"`
		
		Code *string `json:"code,omitempty"`
		
		Details *Trunkinstancetopictrunkerrorinfodetails `json:"details,omitempty"`
		*Alias
	}{ 
		Text: u.Text,
		
		Code: u.Code,
		
		Details: u.Details,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Trunkinstancetopictrunkerrorinfo) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
