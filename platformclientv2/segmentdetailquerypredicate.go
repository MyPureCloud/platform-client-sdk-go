package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Segmentdetailquerypredicate
type Segmentdetailquerypredicate struct { 
	// VarType - Optional type, can usually be inferred
	VarType *string `json:"type,omitempty"`


	// Dimension - Left hand side for dimension predicates
	Dimension *string `json:"dimension,omitempty"`


	// PropertyType - Left hand side for property predicates
	PropertyType *string `json:"propertyType,omitempty"`


	// Property - Left hand side for property predicates
	Property *string `json:"property,omitempty"`


	// Metric - Left hand side for metric predicates
	Metric *string `json:"metric,omitempty"`


	// Operator - Optional operator, default is matches
	Operator *string `json:"operator,omitempty"`


	// Value - Right hand side for dimension, metric, or property predicates
	Value *string `json:"value,omitempty"`


	// VarRange - Right hand side for dimension, metric, or property predicates
	VarRange *Numericrange `json:"range,omitempty"`

}

func (u *Segmentdetailquerypredicate) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Segmentdetailquerypredicate

	

	return json.Marshal(&struct { 
		VarType *string `json:"type,omitempty"`
		
		Dimension *string `json:"dimension,omitempty"`
		
		PropertyType *string `json:"propertyType,omitempty"`
		
		Property *string `json:"property,omitempty"`
		
		Metric *string `json:"metric,omitempty"`
		
		Operator *string `json:"operator,omitempty"`
		
		Value *string `json:"value,omitempty"`
		
		VarRange *Numericrange `json:"range,omitempty"`
		*Alias
	}{ 
		VarType: u.VarType,
		
		Dimension: u.Dimension,
		
		PropertyType: u.PropertyType,
		
		Property: u.Property,
		
		Metric: u.Metric,
		
		Operator: u.Operator,
		
		Value: u.Value,
		
		VarRange: u.VarRange,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Segmentdetailquerypredicate) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
