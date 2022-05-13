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

func (o *Resolutiondetailquerypredicate) MarshalJSON() ([]byte, error) {
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
		VarType: o.VarType,
		
		Metric: o.Metric,
		
		Operator: o.Operator,
		
		Value: o.Value,
		
		VarRange: o.VarRange,
		Alias:    (*Alias)(o),
	})
}

func (o *Resolutiondetailquerypredicate) UnmarshalJSON(b []byte) error {
	var ResolutiondetailquerypredicateMap map[string]interface{}
	err := json.Unmarshal(b, &ResolutiondetailquerypredicateMap)
	if err != nil {
		return err
	}
	
	if VarType, ok := ResolutiondetailquerypredicateMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if Metric, ok := ResolutiondetailquerypredicateMap["metric"].(string); ok {
		o.Metric = &Metric
	}
    
	if Operator, ok := ResolutiondetailquerypredicateMap["operator"].(string); ok {
		o.Operator = &Operator
	}
    
	if Value, ok := ResolutiondetailquerypredicateMap["value"].(string); ok {
		o.Value = &Value
	}
    
	if VarRange, ok := ResolutiondetailquerypredicateMap["range"].(map[string]interface{}); ok {
		VarRangeString, _ := json.Marshal(VarRange)
		json.Unmarshal(VarRangeString, &o.VarRange)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Resolutiondetailquerypredicate) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
