package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Visibilitycondition
type Visibilitycondition struct { 
	// CombiningOperation
	CombiningOperation *string `json:"combiningOperation,omitempty"`


	// Predicates - A list of strings, each representing the location in the form of the Answer Option to depend on. In the format of \"/form/questionGroup/{questionGroupIndex}/question/{questionIndex}/answer/{answerIndex}\" or, to assume the current question group, \"../question/{questionIndex}/answer/{answerIndex}\". Note: Indexes are zero-based
	Predicates *[]interface{} `json:"predicates,omitempty"`

}

func (o *Visibilitycondition) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Visibilitycondition
	
	return json.Marshal(&struct { 
		CombiningOperation *string `json:"combiningOperation,omitempty"`
		
		Predicates *[]interface{} `json:"predicates,omitempty"`
		*Alias
	}{ 
		CombiningOperation: o.CombiningOperation,
		
		Predicates: o.Predicates,
		Alias:    (*Alias)(o),
	})
}

func (o *Visibilitycondition) UnmarshalJSON(b []byte) error {
	var VisibilityconditionMap map[string]interface{}
	err := json.Unmarshal(b, &VisibilityconditionMap)
	if err != nil {
		return err
	}
	
	if CombiningOperation, ok := VisibilityconditionMap["combiningOperation"].(string); ok {
		o.CombiningOperation = &CombiningOperation
	}
    
	if Predicates, ok := VisibilityconditionMap["predicates"].([]interface{}); ok {
		PredicatesString, _ := json.Marshal(Predicates)
		json.Unmarshal(PredicatesString, &o.Predicates)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Visibilitycondition) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
