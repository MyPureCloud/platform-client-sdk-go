package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Documentquerypredicate
type Documentquerypredicate struct { 
	// Fields - Specifies the document fields to be matched against.
	Fields *[]string `json:"fields,omitempty"`


	// Values - Specifies the values of the fields to be matched against.
	Values *[]string `json:"values,omitempty"`


	// VarType - Specifies the matching criteria between the fields and values.
	VarType *string `json:"type,omitempty"`

}

func (o *Documentquerypredicate) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Documentquerypredicate
	
	return json.Marshal(&struct { 
		Fields *[]string `json:"fields,omitempty"`
		
		Values *[]string `json:"values,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		*Alias
	}{ 
		Fields: o.Fields,
		
		Values: o.Values,
		
		VarType: o.VarType,
		Alias:    (*Alias)(o),
	})
}

func (o *Documentquerypredicate) UnmarshalJSON(b []byte) error {
	var DocumentquerypredicateMap map[string]interface{}
	err := json.Unmarshal(b, &DocumentquerypredicateMap)
	if err != nil {
		return err
	}
	
	if Fields, ok := DocumentquerypredicateMap["fields"].([]interface{}); ok {
		FieldsString, _ := json.Marshal(Fields)
		json.Unmarshal(FieldsString, &o.Fields)
	}
	
	if Values, ok := DocumentquerypredicateMap["values"].([]interface{}); ok {
		ValuesString, _ := json.Marshal(Values)
		json.Unmarshal(ValuesString, &o.Values)
	}
	
	if VarType, ok := DocumentquerypredicateMap["type"].(string); ok {
		o.VarType = &VarType
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Documentquerypredicate) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
