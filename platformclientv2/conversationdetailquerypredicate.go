package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationdetailquerypredicate
type Conversationdetailquerypredicate struct { 
	// VarType - Optional type, can usually be inferred
	VarType *string `json:"type,omitempty"`


	// Dimension - Left hand side for dimension predicates
	Dimension *string `json:"dimension,omitempty"`


	// Metric - Left hand side for metric predicates
	Metric *string `json:"metric,omitempty"`


	// Operator - Optional operator, default is matches
	Operator *string `json:"operator,omitempty"`


	// Value - Right hand side for dimension or metric predicates
	Value *string `json:"value,omitempty"`


	// VarRange - Right hand side for dimension or metric predicates
	VarRange *Numericrange `json:"range,omitempty"`

}

func (o *Conversationdetailquerypredicate) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Conversationdetailquerypredicate
	
	return json.Marshal(&struct { 
		VarType *string `json:"type,omitempty"`
		
		Dimension *string `json:"dimension,omitempty"`
		
		Metric *string `json:"metric,omitempty"`
		
		Operator *string `json:"operator,omitempty"`
		
		Value *string `json:"value,omitempty"`
		
		VarRange *Numericrange `json:"range,omitempty"`
		*Alias
	}{ 
		VarType: o.VarType,
		
		Dimension: o.Dimension,
		
		Metric: o.Metric,
		
		Operator: o.Operator,
		
		Value: o.Value,
		
		VarRange: o.VarRange,
		Alias:    (*Alias)(o),
	})
}

func (o *Conversationdetailquerypredicate) UnmarshalJSON(b []byte) error {
	var ConversationdetailquerypredicateMap map[string]interface{}
	err := json.Unmarshal(b, &ConversationdetailquerypredicateMap)
	if err != nil {
		return err
	}
	
	if VarType, ok := ConversationdetailquerypredicateMap["type"].(string); ok {
		o.VarType = &VarType
	}
	
	if Dimension, ok := ConversationdetailquerypredicateMap["dimension"].(string); ok {
		o.Dimension = &Dimension
	}
	
	if Metric, ok := ConversationdetailquerypredicateMap["metric"].(string); ok {
		o.Metric = &Metric
	}
	
	if Operator, ok := ConversationdetailquerypredicateMap["operator"].(string); ok {
		o.Operator = &Operator
	}
	
	if Value, ok := ConversationdetailquerypredicateMap["value"].(string); ok {
		o.Value = &Value
	}
	
	if VarRange, ok := ConversationdetailquerypredicateMap["range"].(map[string]interface{}); ok {
		VarRangeString, _ := json.Marshal(VarRange)
		json.Unmarshal(VarRangeString, &o.VarRange)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Conversationdetailquerypredicate) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
