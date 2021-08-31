package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Documentattribute
type Documentattribute struct { 
	// Attribute
	Attribute *Attribute `json:"attribute,omitempty"`


	// Values
	Values *[]string `json:"values,omitempty"`

}

func (o *Documentattribute) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Documentattribute
	
	return json.Marshal(&struct { 
		Attribute *Attribute `json:"attribute,omitempty"`
		
		Values *[]string `json:"values,omitempty"`
		*Alias
	}{ 
		Attribute: o.Attribute,
		
		Values: o.Values,
		Alias:    (*Alias)(o),
	})
}

func (o *Documentattribute) UnmarshalJSON(b []byte) error {
	var DocumentattributeMap map[string]interface{}
	err := json.Unmarshal(b, &DocumentattributeMap)
	if err != nil {
		return err
	}
	
	if Attribute, ok := DocumentattributeMap["attribute"].(map[string]interface{}); ok {
		AttributeString, _ := json.Marshal(Attribute)
		json.Unmarshal(AttributeString, &o.Attribute)
	}
	
	if Values, ok := DocumentattributeMap["values"].([]interface{}); ok {
		ValuesString, _ := json.Marshal(Values)
		json.Unmarshal(ValuesString, &o.Values)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Documentattribute) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
