package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Policyerrors
type Policyerrors struct { 
	// PolicyErrorMessages
	PolicyErrorMessages *[]Policyerrormessage `json:"policyErrorMessages,omitempty"`

}

func (u *Policyerrors) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Policyerrors

	

	return json.Marshal(&struct { 
		PolicyErrorMessages *[]Policyerrormessage `json:"policyErrorMessages,omitempty"`
		*Alias
	}{ 
		PolicyErrorMessages: u.PolicyErrorMessages,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Policyerrors) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
