package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Learningassignmentaggregatequeryrequestpredicate
type Learningassignmentaggregatequeryrequestpredicate struct { 
	// Dimension - Each predicates specifies a dimension.
	Dimension *string `json:"dimension,omitempty"`


	// Value - Corresponding value for dimensions in predicates. If the dimension is type, Valid Values: Informational, AssessedContent, Assessment
	Value *string `json:"value,omitempty"`

}

func (o *Learningassignmentaggregatequeryrequestpredicate) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Learningassignmentaggregatequeryrequestpredicate
	
	return json.Marshal(&struct { 
		Dimension *string `json:"dimension,omitempty"`
		
		Value *string `json:"value,omitempty"`
		*Alias
	}{ 
		Dimension: o.Dimension,
		
		Value: o.Value,
		Alias:    (*Alias)(o),
	})
}

func (o *Learningassignmentaggregatequeryrequestpredicate) UnmarshalJSON(b []byte) error {
	var LearningassignmentaggregatequeryrequestpredicateMap map[string]interface{}
	err := json.Unmarshal(b, &LearningassignmentaggregatequeryrequestpredicateMap)
	if err != nil {
		return err
	}
	
	if Dimension, ok := LearningassignmentaggregatequeryrequestpredicateMap["dimension"].(string); ok {
		o.Dimension = &Dimension
	}
	
	if Value, ok := LearningassignmentaggregatequeryrequestpredicateMap["value"].(string); ok {
		o.Value = &Value
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Learningassignmentaggregatequeryrequestpredicate) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
