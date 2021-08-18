package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Availabletranslations
type Availabletranslations struct { 
	// OrgSpecific
	OrgSpecific *[]string `json:"orgSpecific,omitempty"`


	// Builtin
	Builtin *[]string `json:"builtin,omitempty"`

}

func (u *Availabletranslations) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Availabletranslations

	

	return json.Marshal(&struct { 
		OrgSpecific *[]string `json:"orgSpecific,omitempty"`
		
		Builtin *[]string `json:"builtin,omitempty"`
		*Alias
	}{ 
		OrgSpecific: u.OrgSpecific,
		
		Builtin: u.Builtin,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Availabletranslations) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
