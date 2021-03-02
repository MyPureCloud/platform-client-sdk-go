package platformclientv2
import (
	"encoding/json"
)

// Visibilitycondition
type Visibilitycondition struct { 
	// CombiningOperation
	CombiningOperation *string `json:"combiningOperation,omitempty"`


	// Predicates - A list of strings, each representing the location in the form of the Answer Option to depend on. In the format of \"/form/questionGroup/{questionGroupIndex}/question/{questionIndex}/answer/{answerIndex}\" or, to assume the current question group, \"../question/{questionIndex}/answer/{answerIndex}\". Note: Indexes are zero-based
	Predicates *[]map[string]interface{} `json:"predicates,omitempty"`

}

// String returns a JSON representation of the model
func (o *Visibilitycondition) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
