package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Contentfacetfilteritem
type Contentfacetfilteritem struct { 
	// Name
	Name *string `json:"name,omitempty"`


	// VarType
	VarType *string `json:"type,omitempty"`


	// Operator
	Operator *string `json:"operator,omitempty"`


	// Values
	Values *[]string `json:"values,omitempty"`

}

func (o *Contentfacetfilteritem) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Contentfacetfilteritem
	
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

func (o *Contentfacetfilteritem) UnmarshalJSON(b []byte) error {
	var ContentfacetfilteritemMap map[string]interface{}
	err := json.Unmarshal(b, &ContentfacetfilteritemMap)
	if err != nil {
		return err
	}
	
	if Name, ok := ContentfacetfilteritemMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if VarType, ok := ContentfacetfilteritemMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if Operator, ok := ContentfacetfilteritemMap["operator"].(string); ok {
		o.Operator = &Operator
	}
    
	if Values, ok := ContentfacetfilteritemMap["values"].([]interface{}); ok {
		ValuesString, _ := json.Marshal(Values)
		json.Unmarshal(ValuesString, &o.Values)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Contentfacetfilteritem) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
