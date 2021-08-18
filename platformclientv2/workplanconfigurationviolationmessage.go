package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Workplanconfigurationviolationmessage
type Workplanconfigurationviolationmessage struct { 
	// VarType - Type of configuration violation message for this work plan
	VarType *string `json:"type,omitempty"`


	// Arguments - Arguments of the message that provide information about the misconfigured value or the threshold that is exceeded by the misconfigured value
	Arguments *[]Workplanvalidationmessageargument `json:"arguments,omitempty"`


	// Severity - Severity of the message. A message with Error severity indicates the scheduler won't be able to produce schedules and thus the work plan is invalid.
	Severity *string `json:"severity,omitempty"`

}

func (u *Workplanconfigurationviolationmessage) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Workplanconfigurationviolationmessage

	

	return json.Marshal(&struct { 
		VarType *string `json:"type,omitempty"`
		
		Arguments *[]Workplanvalidationmessageargument `json:"arguments,omitempty"`
		
		Severity *string `json:"severity,omitempty"`
		*Alias
	}{ 
		VarType: u.VarType,
		
		Arguments: u.Arguments,
		
		Severity: u.Severity,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Workplanconfigurationviolationmessage) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
