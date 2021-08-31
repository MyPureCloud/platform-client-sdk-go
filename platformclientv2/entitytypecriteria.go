package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Entitytypecriteria
type Entitytypecriteria struct { 
	// Key - The criteria key.
	Key *string `json:"key,omitempty"`


	// Values - The criteria values.
	Values *[]string `json:"values,omitempty"`


	// ShouldIgnoreCase - Should criteria be case insensitive.
	ShouldIgnoreCase *bool `json:"shouldIgnoreCase,omitempty"`


	// Operator - The comparison operator.
	Operator *string `json:"operator,omitempty"`


	// EntityType - The entity to match the pattern against.
	EntityType *string `json:"entityType,omitempty"`

}

func (o *Entitytypecriteria) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Entitytypecriteria
	
	return json.Marshal(&struct { 
		Key *string `json:"key,omitempty"`
		
		Values *[]string `json:"values,omitempty"`
		
		ShouldIgnoreCase *bool `json:"shouldIgnoreCase,omitempty"`
		
		Operator *string `json:"operator,omitempty"`
		
		EntityType *string `json:"entityType,omitempty"`
		*Alias
	}{ 
		Key: o.Key,
		
		Values: o.Values,
		
		ShouldIgnoreCase: o.ShouldIgnoreCase,
		
		Operator: o.Operator,
		
		EntityType: o.EntityType,
		Alias:    (*Alias)(o),
	})
}

func (o *Entitytypecriteria) UnmarshalJSON(b []byte) error {
	var EntitytypecriteriaMap map[string]interface{}
	err := json.Unmarshal(b, &EntitytypecriteriaMap)
	if err != nil {
		return err
	}
	
	if Key, ok := EntitytypecriteriaMap["key"].(string); ok {
		o.Key = &Key
	}
	
	if Values, ok := EntitytypecriteriaMap["values"].([]interface{}); ok {
		ValuesString, _ := json.Marshal(Values)
		json.Unmarshal(ValuesString, &o.Values)
	}
	
	if ShouldIgnoreCase, ok := EntitytypecriteriaMap["shouldIgnoreCase"].(bool); ok {
		o.ShouldIgnoreCase = &ShouldIgnoreCase
	}
	
	if Operator, ok := EntitytypecriteriaMap["operator"].(string); ok {
		o.Operator = &Operator
	}
	
	if EntityType, ok := EntitytypecriteriaMap["entityType"].(string); ok {
		o.EntityType = &EntityType
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Entitytypecriteria) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
