package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Policyerrors
type Policyerrors struct { 
	// PolicyErrorMessages
	PolicyErrorMessages *[]Policyerrormessage `json:"policyErrorMessages,omitempty"`

}

// String returns a JSON representation of the model
func (o *Policyerrors) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
