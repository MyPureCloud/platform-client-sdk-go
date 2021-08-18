package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Workplanconstraintconflictmessage
type Workplanconstraintconflictmessage struct { 
	// VarType - Type of constraint conflict that can be resolved by clients in order to generate agent schedules
	VarType *string `json:"type,omitempty"`


	// Arguments - The arguments to the type of the message that can help clients resolve validation issues
	Arguments *[]Workplanvalidationmessageargument `json:"arguments,omitempty"`

}

func (u *Workplanconstraintconflictmessage) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Workplanconstraintconflictmessage

	

	return json.Marshal(&struct { 
		VarType *string `json:"type,omitempty"`
		
		Arguments *[]Workplanvalidationmessageargument `json:"arguments,omitempty"`
		*Alias
	}{ 
		VarType: u.VarType,
		
		Arguments: u.Arguments,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Workplanconstraintconflictmessage) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
