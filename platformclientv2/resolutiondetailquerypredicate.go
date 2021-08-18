package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Resolutiondetailquerypredicate
type Resolutiondetailquerypredicate struct { 
	// VarType - Optional type, can usually be inferred
	VarType *string `json:"type,omitempty"`


	// Metric - Left hand side for metric predicates
	Metric *string `json:"metric,omitempty"`


	// Operator - Optional operator, default is matches
	Operator *string `json:"operator,omitempty"`


	// Value - Right hand side for metric predicates
	Value *string `json:"value,omitempty"`


	// VarRange - Right hand side for metric predicates
	VarRange *Numericrange `json:"range,omitempty"`

}

func (u *Resolutiondetailquerypredicate) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Resolutiondetailquerypredicate

	

	return json.Marshal(&struct { 
		VarType *string `json:"type,omitempty"`
		
		Metric *string `json:"metric,omitempty"`
		
		Operator *string `json:"operator,omitempty"`
		
		Value *string `json:"value,omitempty"`
		
		VarRange *Numericrange `json:"range,omitempty"`
		*Alias
	}{ 
		VarType: u.VarType,
		
		Metric: u.Metric,
		
		Operator: u.Operator,
		
		Value: u.Value,
		
		VarRange: u.VarRange,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Resolutiondetailquerypredicate) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
