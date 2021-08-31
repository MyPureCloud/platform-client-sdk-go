package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Condition
type Condition struct { 
	// VarType - The type of the condition.
	VarType *string `json:"type,omitempty"`


	// Inverted - If true, inverts the result of evaluating this Condition. Default is false.
	Inverted *bool `json:"inverted,omitempty"`


	// AttributeName - An attribute name associated with this Condition. Required for a contactAttributeCondition.
	AttributeName *string `json:"attributeName,omitempty"`


	// Value - A value associated with this Condition. This could be text, a number, or a relative time. Not used for a DataActionCondition.
	Value *string `json:"value,omitempty"`


	// ValueType - The type of the value associated with this Condition. Not used for a DataActionCondition.
	ValueType *string `json:"valueType,omitempty"`


	// Operator - An operation with which to evaluate the Condition. Not used for a DataActionCondition.
	Operator *string `json:"operator,omitempty"`


	// Codes - List of wrap-up code identifiers. Required for a wrapupCondition.
	Codes *[]string `json:"codes,omitempty"`


	// Property - A value associated with the property type of this Condition. Required for a contactPropertyCondition.
	Property *string `json:"property,omitempty"`


	// PropertyType - The type of the property associated with this Condition. Required for a contactPropertyCondition.
	PropertyType *string `json:"propertyType,omitempty"`

}

func (o *Condition) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Condition
	
	return json.Marshal(&struct { 
		VarType *string `json:"type,omitempty"`
		
		Inverted *bool `json:"inverted,omitempty"`
		
		AttributeName *string `json:"attributeName,omitempty"`
		
		Value *string `json:"value,omitempty"`
		
		ValueType *string `json:"valueType,omitempty"`
		
		Operator *string `json:"operator,omitempty"`
		
		Codes *[]string `json:"codes,omitempty"`
		
		Property *string `json:"property,omitempty"`
		
		PropertyType *string `json:"propertyType,omitempty"`
		*Alias
	}{ 
		VarType: o.VarType,
		
		Inverted: o.Inverted,
		
		AttributeName: o.AttributeName,
		
		Value: o.Value,
		
		ValueType: o.ValueType,
		
		Operator: o.Operator,
		
		Codes: o.Codes,
		
		Property: o.Property,
		
		PropertyType: o.PropertyType,
		Alias:    (*Alias)(o),
	})
}

func (o *Condition) UnmarshalJSON(b []byte) error {
	var ConditionMap map[string]interface{}
	err := json.Unmarshal(b, &ConditionMap)
	if err != nil {
		return err
	}
	
	if VarType, ok := ConditionMap["type"].(string); ok {
		o.VarType = &VarType
	}
	
	if Inverted, ok := ConditionMap["inverted"].(bool); ok {
		o.Inverted = &Inverted
	}
	
	if AttributeName, ok := ConditionMap["attributeName"].(string); ok {
		o.AttributeName = &AttributeName
	}
	
	if Value, ok := ConditionMap["value"].(string); ok {
		o.Value = &Value
	}
	
	if ValueType, ok := ConditionMap["valueType"].(string); ok {
		o.ValueType = &ValueType
	}
	
	if Operator, ok := ConditionMap["operator"].(string); ok {
		o.Operator = &Operator
	}
	
	if Codes, ok := ConditionMap["codes"].([]interface{}); ok {
		CodesString, _ := json.Marshal(Codes)
		json.Unmarshal(CodesString, &o.Codes)
	}
	
	if Property, ok := ConditionMap["property"].(string); ok {
		o.Property = &Property
	}
	
	if PropertyType, ok := ConditionMap["propertyType"].(string); ok {
		o.PropertyType = &PropertyType
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Condition) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
