package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Contentattributefilteritem
type Contentattributefilteritem struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Operator
	Operator *string `json:"operator,omitempty"`


	// Values
	Values *[]string `json:"values,omitempty"`

}

func (o *Contentattributefilteritem) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Contentattributefilteritem
	
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

func (o *Contentattributefilteritem) UnmarshalJSON(b []byte) error {
	var ContentattributefilteritemMap map[string]interface{}
	err := json.Unmarshal(b, &ContentattributefilteritemMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ContentattributefilteritemMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Operator, ok := ContentattributefilteritemMap["operator"].(string); ok {
		o.Operator = &Operator
	}
    
	if Values, ok := ContentattributefilteritemMap["values"].([]interface{}); ok {
		ValuesString, _ := json.Marshal(Values)
		json.Unmarshal(ValuesString, &o.Values)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Contentattributefilteritem) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
