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

func (o *Policyerrors) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Policyerrors
	
	return json.Marshal(&struct { 
		PolicyErrorMessages *[]Policyerrormessage `json:"policyErrorMessages,omitempty"`
		*Alias
	}{ 
		PolicyErrorMessages: o.PolicyErrorMessages,
		Alias:    (*Alias)(o),
	})
}

func (o *Policyerrors) UnmarshalJSON(b []byte) error {
	var PolicyerrorsMap map[string]interface{}
	err := json.Unmarshal(b, &PolicyerrorsMap)
	if err != nil {
		return err
	}
	
	if PolicyErrorMessages, ok := PolicyerrorsMap["policyErrorMessages"].([]interface{}); ok {
		PolicyErrorMessagesString, _ := json.Marshal(PolicyErrorMessages)
		json.Unmarshal(PolicyErrorMessagesString, &o.PolicyErrorMessages)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Policyerrors) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
