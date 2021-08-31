package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Criteria
type Criteria struct { 
	// Key - The criteria key.
	Key *string `json:"key,omitempty"`


	// Values - The criteria values.
	Values *[]string `json:"values,omitempty"`


	// ShouldIgnoreCase - Should criteria be case insensitive.
	ShouldIgnoreCase *bool `json:"shouldIgnoreCase,omitempty"`


	// Operator - The comparison operator.
	Operator *string `json:"operator,omitempty"`

}

func (o *Criteria) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Criteria
	
	return json.Marshal(&struct { 
		Key *string `json:"key,omitempty"`
		
		Values *[]string `json:"values,omitempty"`
		
		ShouldIgnoreCase *bool `json:"shouldIgnoreCase,omitempty"`
		
		Operator *string `json:"operator,omitempty"`
		*Alias
	}{ 
		Key: o.Key,
		
		Values: o.Values,
		
		ShouldIgnoreCase: o.ShouldIgnoreCase,
		
		Operator: o.Operator,
		Alias:    (*Alias)(o),
	})
}

func (o *Criteria) UnmarshalJSON(b []byte) error {
	var CriteriaMap map[string]interface{}
	err := json.Unmarshal(b, &CriteriaMap)
	if err != nil {
		return err
	}
	
	if Key, ok := CriteriaMap["key"].(string); ok {
		o.Key = &Key
	}
	
	if Values, ok := CriteriaMap["values"].([]interface{}); ok {
		ValuesString, _ := json.Marshal(Values)
		json.Unmarshal(ValuesString, &o.Values)
	}
	
	if ShouldIgnoreCase, ok := CriteriaMap["shouldIgnoreCase"].(bool); ok {
		o.ShouldIgnoreCase = &ShouldIgnoreCase
	}
	
	if Operator, ok := CriteriaMap["operator"].(string); ok {
		o.Operator = &Operator
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Criteria) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
