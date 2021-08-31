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

func (o *Actioncontractinput) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Actioncontractinput
	
	return json.Marshal(&struct { 
		Input *Postinputcontract `json:"input,omitempty"`
		
		Output *Postoutputcontract `json:"output,omitempty"`
		*Alias
	}{ 
		Input: o.Input,
		
		Output: o.Output,
		Alias:    (*Alias)(o),
	})
}

func (o *Actioncontractinput) UnmarshalJSON(b []byte) error {
	var ActioncontractinputMap map[string]interface{}
	err := json.Unmarshal(b, &ActioncontractinputMap)
	if err != nil {
		return err
	}
	
	if Input, ok := ActioncontractinputMap["input"].(map[string]interface{}); ok {
		InputString, _ := json.Marshal(Input)
		json.Unmarshal(InputString, &o.Input)
	}
	
	if Output, ok := ActioncontractinputMap["output"].(map[string]interface{}); ok {
		OutputString, _ := json.Marshal(Output)
		json.Unmarshal(OutputString, &o.Output)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Actioncontractinput) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
