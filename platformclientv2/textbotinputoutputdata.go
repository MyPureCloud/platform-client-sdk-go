package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Textbotinputoutputdata - Input/Output data related to a bot flow which is exiting gracefully.
type Textbotinputoutputdata struct { 
	// Variables - The input/output variables using the format as appropriate for the variable data type in the flow definition.
	Variables *map[string]interface{} `json:"variables,omitempty"`

}

func (o *Textbotinputoutputdata) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Textbotinputoutputdata
	
	return json.Marshal(&struct { 
		Variables *map[string]interface{} `json:"variables,omitempty"`
		*Alias
	}{ 
		Variables: o.Variables,
		Alias:    (*Alias)(o),
	})
}

func (o *Textbotinputoutputdata) UnmarshalJSON(b []byte) error {
	var TextbotinputoutputdataMap map[string]interface{}
	err := json.Unmarshal(b, &TextbotinputoutputdataMap)
	if err != nil {
		return err
	}
	
	if Variables, ok := TextbotinputoutputdataMap["variables"].(map[string]interface{}); ok {
		VariablesString, _ := json.Marshal(Variables)
		json.Unmarshal(VariablesString, &o.Variables)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Textbotinputoutputdata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
