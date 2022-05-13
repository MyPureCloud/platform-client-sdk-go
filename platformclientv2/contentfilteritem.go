package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Contentfilteritem
type Contentfilteritem struct { 
	// Name
	Name *string `json:"name,omitempty"`


	// VarType
	VarType *string `json:"type,omitempty"`


	// Operator
	Operator *string `json:"operator,omitempty"`


	// Values
	Values *[]string `json:"values,omitempty"`

}

func (o *Contentfilteritem) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Contentfilteritem
	
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

func (o *Contentfilteritem) UnmarshalJSON(b []byte) error {
	var ContentfilteritemMap map[string]interface{}
	err := json.Unmarshal(b, &ContentfilteritemMap)
	if err != nil {
		return err
	}
	
	if Name, ok := ContentfilteritemMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if VarType, ok := ContentfilteritemMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if Operator, ok := ContentfilteritemMap["operator"].(string); ok {
		o.Operator = &Operator
	}
    
	if Values, ok := ContentfilteritemMap["values"].([]interface{}); ok {
		ValuesString, _ := json.Marshal(Values)
		json.Unmarshal(ValuesString, &o.Values)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Contentfilteritem) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
