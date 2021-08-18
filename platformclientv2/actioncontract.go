package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Actioncontract - This resource contains all of the schemas needed to define the inputs and outputs, of a single Action.
type Actioncontract struct { 
	// Output - The output to expect when executing this action.
	Output *Actionoutput `json:"output,omitempty"`


	// Input - The input required when executing this action.
	Input *Actioninput `json:"input,omitempty"`

}

func (u *Actioncontract) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Actioncontract

	

	return json.Marshal(&struct { 
		Output *Actionoutput `json:"output,omitempty"`
		
		Input *Actioninput `json:"input,omitempty"`
		*Alias
	}{ 
		Output: u.Output,
		
		Input: u.Input,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Actioncontract) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
