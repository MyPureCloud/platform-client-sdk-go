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

func (u *Condition) MarshalJSON() ([]byte, error) {
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
		VarType: u.VarType,
		
		Inverted: u.Inverted,
		
		AttributeName: u.AttributeName,
		
		Value: u.Value,
		
		ValueType: u.ValueType,
		
		Operator: u.Operator,
		
		Codes: u.Codes,
		
		Property: u.Property,
		
		PropertyType: u.PropertyType,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Condition) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
