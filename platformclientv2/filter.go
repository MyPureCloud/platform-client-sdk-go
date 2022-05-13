package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Filter
type Filter struct { 
	// Name - The name of the field by which to filter.
	Name *string `json:"name,omitempty"`


	// VarType - The type of the filter, DATE or STRING.
	VarType *string `json:"type,omitempty"`


	// Operator - The operation that the filter performs.
	Operator *string `json:"operator,omitempty"`


	// Values - The values to make the filter comparison against.
	Values *[]string `json:"values,omitempty"`

}

func (o *Filter) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Filter
	
	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		
		Operator *string `json:"operator,omitempty"`
		
		Values *[]string `json:"values,omitempty"`
		*Alias
	}{ 
		Name: o.Name,
		
		VarType: o.VarType,
		
		Operator: o.Operator,
		
		Values: o.Values,
		Alias:    (*Alias)(o),
	})
}

func (o *Filter) UnmarshalJSON(b []byte) error {
	var FilterMap map[string]interface{}
	err := json.Unmarshal(b, &FilterMap)
	if err != nil {
		return err
	}
	
	if Name, ok := FilterMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if VarType, ok := FilterMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if Operator, ok := FilterMap["operator"].(string); ok {
		o.Operator = &Operator
	}
    
	if Values, ok := FilterMap["values"].([]interface{}); ok {
		ValuesString, _ := json.Marshal(Values)
		json.Unmarshal(ValuesString, &o.Values)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Filter) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
