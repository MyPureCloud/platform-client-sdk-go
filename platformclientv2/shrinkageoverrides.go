package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Shrinkageoverrides
type Shrinkageoverrides struct { 
	// Clear - Set true to clear the shrinkage interval overrides
	Clear *bool `json:"clear,omitempty"`


	// Values - List of interval shrinkage overrides
	Values *[]Shrinkageoverride `json:"values,omitempty"`

}

func (o *Shrinkageoverrides) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Shrinkageoverrides
	
	return json.Marshal(&struct { 
		Clear *bool `json:"clear,omitempty"`
		
		Values *[]Shrinkageoverride `json:"values,omitempty"`
		*Alias
	}{ 
		Clear: o.Clear,
		
		Values: o.Values,
		Alias:    (*Alias)(o),
	})
}

func (o *Shrinkageoverrides) UnmarshalJSON(b []byte) error {
	var ShrinkageoverridesMap map[string]interface{}
	err := json.Unmarshal(b, &ShrinkageoverridesMap)
	if err != nil {
		return err
	}
	
	if Clear, ok := ShrinkageoverridesMap["clear"].(bool); ok {
		o.Clear = &Clear
	}
	
	if Values, ok := ShrinkageoverridesMap["values"].([]interface{}); ok {
		ValuesString, _ := json.Marshal(Values)
		json.Unmarshal(ValuesString, &o.Values)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Shrinkageoverrides) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
