package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Attributefilteritem
type Attributefilteritem struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Operator
	Operator *string `json:"operator,omitempty"`


	// Values
	Values *[]string `json:"values,omitempty"`

}

func (o *Attributefilteritem) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Attributefilteritem
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Operator *string `json:"operator,omitempty"`
		
		Values *[]string `json:"values,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Operator: o.Operator,
		
		Values: o.Values,
		Alias:    (*Alias)(o),
	})
}

func (o *Attributefilteritem) UnmarshalJSON(b []byte) error {
	var AttributefilteritemMap map[string]interface{}
	err := json.Unmarshal(b, &AttributefilteritemMap)
	if err != nil {
		return err
	}
	
	if Id, ok := AttributefilteritemMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Operator, ok := AttributefilteritemMap["operator"].(string); ok {
		o.Operator = &Operator
	}
    
	if Values, ok := AttributefilteritemMap["values"].([]interface{}); ok {
		ValuesString, _ := json.Marshal(Values)
		json.Unmarshal(ValuesString, &o.Values)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Attributefilteritem) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
