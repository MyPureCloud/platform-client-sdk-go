package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Urlcondition
type Urlcondition struct { 
	// Values - The URL condition value.
	Values *[]string `json:"values,omitempty"`


	// Operator - The comparison operator.
	Operator *string `json:"operator,omitempty"`

}

func (o *Urlcondition) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Urlcondition
	
	return json.Marshal(&struct { 
		Values *[]string `json:"values,omitempty"`
		
		Operator *string `json:"operator,omitempty"`
		*Alias
	}{ 
		Values: o.Values,
		
		Operator: o.Operator,
		Alias:    (*Alias)(o),
	})
}

func (o *Urlcondition) UnmarshalJSON(b []byte) error {
	var UrlconditionMap map[string]interface{}
	err := json.Unmarshal(b, &UrlconditionMap)
	if err != nil {
		return err
	}
	
	if Values, ok := UrlconditionMap["values"].([]interface{}); ok {
		ValuesString, _ := json.Marshal(Values)
		json.Unmarshal(ValuesString, &o.Values)
	}
	
	if Operator, ok := UrlconditionMap["operator"].(string); ok {
		o.Operator = &Operator
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Urlcondition) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
