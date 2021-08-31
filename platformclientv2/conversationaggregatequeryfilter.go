package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationaggregatequeryfilter
type Conversationaggregatequeryfilter struct { 
	// VarType - Boolean operation to apply to the provided predicates and clauses
	VarType *string `json:"type,omitempty"`


	// Clauses - Boolean 'and/or' logic with up to two-levels of nesting
	Clauses *[]Conversationaggregatequeryclause `json:"clauses,omitempty"`


	// Predicates - Like a three-word sentence: (attribute-name) (operator) (target-value).
	Predicates *[]Conversationaggregatequerypredicate `json:"predicates,omitempty"`

}

func (o *Conversationaggregatequeryfilter) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Conversationaggregatequeryfilter
	
	return json.Marshal(&struct { 
		VarType *string `json:"type,omitempty"`
		
		Clauses *[]Conversationaggregatequeryclause `json:"clauses,omitempty"`
		
		Predicates *[]Conversationaggregatequerypredicate `json:"predicates,omitempty"`
		*Alias
	}{ 
		VarType: o.VarType,
		
		Clauses: o.Clauses,
		
		Predicates: o.Predicates,
		Alias:    (*Alias)(o),
	})
}

func (o *Conversationaggregatequeryfilter) UnmarshalJSON(b []byte) error {
	var ConversationaggregatequeryfilterMap map[string]interface{}
	err := json.Unmarshal(b, &ConversationaggregatequeryfilterMap)
	if err != nil {
		return err
	}
	
	if VarType, ok := ConversationaggregatequeryfilterMap["type"].(string); ok {
		o.VarType = &VarType
	}
	
	if Clauses, ok := ConversationaggregatequeryfilterMap["clauses"].([]interface{}); ok {
		ClausesString, _ := json.Marshal(Clauses)
		json.Unmarshal(ClausesString, &o.Clauses)
	}
	
	if Predicates, ok := ConversationaggregatequeryfilterMap["predicates"].([]interface{}); ok {
		PredicatesString, _ := json.Marshal(Predicates)
		json.Unmarshal(PredicatesString, &o.Predicates)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Conversationaggregatequeryfilter) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
