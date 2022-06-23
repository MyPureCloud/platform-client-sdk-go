package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Wrapupcodeconfig
type Wrapupcodeconfig struct { 
	// Values - A set of valid Wrap Up Code UUIDs used to optimize a KPI.
	Values *[]string `json:"values,omitempty"`

}

func (o *Wrapupcodeconfig) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Wrapupcodeconfig
	
	return json.Marshal(&struct { 
		Values *[]string `json:"values,omitempty"`
		*Alias
	}{ 
		Values: o.Values,
		Alias:    (*Alias)(o),
	})
}

func (o *Wrapupcodeconfig) UnmarshalJSON(b []byte) error {
	var WrapupcodeconfigMap map[string]interface{}
	err := json.Unmarshal(b, &WrapupcodeconfigMap)
	if err != nil {
		return err
	}
	
	if Values, ok := WrapupcodeconfigMap["values"].([]interface{}); ok {
		ValuesString, _ := json.Marshal(Values)
		json.Unmarshal(ValuesString, &o.Values)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Wrapupcodeconfig) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
