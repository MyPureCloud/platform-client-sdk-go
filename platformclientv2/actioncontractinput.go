package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Actioncontractinput - Contract definition.
type Actioncontractinput struct { 
	// Input - Execution input contract
	Input *Postinputcontract `json:"input,omitempty"`


	// Output - Execution output contract
	Output *Postoutputcontract `json:"output,omitempty"`

}

func (u *Actioncontractinput) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Actioncontractinput

	

	return json.Marshal(&struct { 
		Input *Postinputcontract `json:"input,omitempty"`
		
		Output *Postoutputcontract `json:"output,omitempty"`
		*Alias
	}{ 
		Input: u.Input,
		
		Output: u.Output,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Actioncontractinput) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
