package platformclientv2
import (
	"encoding/json"
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

// String returns a JSON representation of the model
func (o *Condition) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
