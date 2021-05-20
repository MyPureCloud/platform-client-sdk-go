package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Assessmentformquestiongroup
type Assessmentformquestiongroup struct { 
	// Id - The ID of the question group,
	Id *string `json:"id,omitempty"`


	// Name - The question group name
	Name *string `json:"name,omitempty"`


	// VarType - The question group type
	VarType *string `json:"type,omitempty"`


	// DefaultAnswersToHighest
	DefaultAnswersToHighest *bool `json:"defaultAnswersToHighest,omitempty"`


	// DefaultAnswersToNA
	DefaultAnswersToNA *bool `json:"defaultAnswersToNA,omitempty"`


	// NaEnabled
	NaEnabled *bool `json:"naEnabled,omitempty"`


	// Weight
	Weight *float32 `json:"weight,omitempty"`


	// ManualWeight
	ManualWeight *bool `json:"manualWeight,omitempty"`


	// Questions - The list of questions for this question group
	Questions *[]Assessmentformquestion `json:"questions,omitempty"`


	// VisibilityCondition
	VisibilityCondition *Visibilitycondition `json:"visibilityCondition,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Assessmentformquestiongroup) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
