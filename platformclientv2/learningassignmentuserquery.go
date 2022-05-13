package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Learningassignmentuserquery - Learning module users query request model
type Learningassignmentuserquery struct { 
	// Rule - Learning module rule object
	Rule *Learningmodulerule `json:"rule,omitempty"`


	// SearchTerm - The user name to be searched for
	SearchTerm *string `json:"searchTerm,omitempty"`

}

func (o *Learningassignmentuserquery) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Learningassignmentuserquery
	
	return json.Marshal(&struct { 
		Rule *Learningmodulerule `json:"rule,omitempty"`
		
		SearchTerm *string `json:"searchTerm,omitempty"`
		*Alias
	}{ 
		Rule: o.Rule,
		
		SearchTerm: o.SearchTerm,
		Alias:    (*Alias)(o),
	})
}

func (o *Learningassignmentuserquery) UnmarshalJSON(b []byte) error {
	var LearningassignmentuserqueryMap map[string]interface{}
	err := json.Unmarshal(b, &LearningassignmentuserqueryMap)
	if err != nil {
		return err
	}
	
	if Rule, ok := LearningassignmentuserqueryMap["rule"].(map[string]interface{}); ok {
		RuleString, _ := json.Marshal(Rule)
		json.Unmarshal(RuleString, &o.Rule)
	}
	
	if SearchTerm, ok := LearningassignmentuserqueryMap["searchTerm"].(string); ok {
		o.SearchTerm = &SearchTerm
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Learningassignmentuserquery) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
