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

func (o *Actioncontract) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Actioncontract
	
	return json.Marshal(&struct { 
		Output *Actionoutput `json:"output,omitempty"`
		
		Input *Actioninput `json:"input,omitempty"`
		*Alias
	}{ 
		Output: o.Output,
		
		Input: o.Input,
		Alias:    (*Alias)(o),
	})
}

func (o *Actioncontract) UnmarshalJSON(b []byte) error {
	var ActioncontractMap map[string]interface{}
	err := json.Unmarshal(b, &ActioncontractMap)
	if err != nil {
		return err
	}
	
	if Output, ok := ActioncontractMap["output"].(map[string]interface{}); ok {
		OutputString, _ := json.Marshal(Output)
		json.Unmarshal(OutputString, &o.Output)
	}
	
	if Input, ok := ActioncontractMap["input"].(map[string]interface{}); ok {
		InputString, _ := json.Marshal(Input)
		json.Unmarshal(InputString, &o.Input)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Actioncontract) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
