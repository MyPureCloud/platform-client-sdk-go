package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Workplanconstraintmessage
type Workplanconstraintmessage struct { 
	// VarType - Type of the work plan constraint in this message
	VarType *string `json:"type,omitempty"`


	// Arguments - Arguments of the message that provide information about the constraint that is being conflicted with, such as the value of the constraint
	Arguments *[]Workplanvalidationmessageargument `json:"arguments,omitempty"`

}

func (u *Workplanconstraintmessage) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Workplanconstraintmessage

	

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
func (o *Workplanconstraintmessage) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
