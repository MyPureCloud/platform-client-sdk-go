package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Matchcriteria - Defines a simple matching condition
type Matchcriteria struct { 
	// JsonPath - The Goessner json path of the field to match
	JsonPath *string `json:"jsonPath,omitempty"`


	// Operator - The type of operation to perform for matching check
	Operator *string `json:"operator,omitempty"`


	// Value - The value to match on. Only one of value and values can be included
	Value *interface{} `json:"value,omitempty"`


	// Values - The list of values to match on. Only one of value and values can be included
	Values *[]map[string]interface{} `json:"values,omitempty"`

}

func (o *Matchcriteria) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Matchcriteria
	
	return json.Marshal(&struct { 
		JsonPath *string `json:"jsonPath,omitempty"`
		
		Operator *string `json:"operator,omitempty"`
		
		Value *interface{} `json:"value,omitempty"`
		
		Values *[]map[string]interface{} `json:"values,omitempty"`
		*Alias
	}{ 
		JsonPath: o.JsonPath,
		
		Operator: o.Operator,
		
		Value: o.Value,
		
		Values: o.Values,
		Alias:    (*Alias)(o),
	})
}

func (o *Matchcriteria) UnmarshalJSON(b []byte) error {
	var MatchcriteriaMap map[string]interface{}
	err := json.Unmarshal(b, &MatchcriteriaMap)
	if err != nil {
		return err
	}
	
	if JsonPath, ok := MatchcriteriaMap["jsonPath"].(string); ok {
		o.JsonPath = &JsonPath
	}
    
	if Operator, ok := MatchcriteriaMap["operator"].(string); ok {
		o.Operator = &Operator
	}
    
	if Value, ok := MatchcriteriaMap["value"].(map[string]interface{}); ok {
		ValueString, _ := json.Marshal(Value)
		json.Unmarshal(ValueString, &o.Value)
	}
	
	if Values, ok := MatchcriteriaMap["values"].([]interface{}); ok {
		ValuesString, _ := json.Marshal(Values)
		json.Unmarshal(ValuesString, &o.Values)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Matchcriteria) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
