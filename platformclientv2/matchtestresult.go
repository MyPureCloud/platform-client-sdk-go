package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Matchtestresult - Information about the results being matched by the expressions
type Matchtestresult struct { 
	// Value - The value of the field being matched
	Value *interface{} `json:"value,omitempty"`


	// Path - The json path to the json node being matched on. ex: $['things'][1]
	Path *string `json:"path,omitempty"`

}

func (o *Matchtestresult) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Matchtestresult
	
	return json.Marshal(&struct { 
		Value *interface{} `json:"value,omitempty"`
		
		Path *string `json:"path,omitempty"`
		*Alias
	}{ 
		Value: o.Value,
		
		Path: o.Path,
		Alias:    (*Alias)(o),
	})
}

func (o *Matchtestresult) UnmarshalJSON(b []byte) error {
	var MatchtestresultMap map[string]interface{}
	err := json.Unmarshal(b, &MatchtestresultMap)
	if err != nil {
		return err
	}
	
	if Value, ok := MatchtestresultMap["value"].(map[string]interface{}); ok {
		ValueString, _ := json.Marshal(Value)
		json.Unmarshal(ValueString, &o.Value)
	}
	
	if Path, ok := MatchtestresultMap["path"].(string); ok {
		o.Path = &Path
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Matchtestresult) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
