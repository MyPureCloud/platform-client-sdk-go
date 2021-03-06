package platformclientv2
import (
	"encoding/json"
)

// Surveyquestiongroup
type Surveyquestiongroup struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// VarType
	VarType *string `json:"type,omitempty"`


	// NaEnabled
	NaEnabled *bool `json:"naEnabled,omitempty"`


	// Questions
	Questions *[]Surveyquestion `json:"questions,omitempty"`


	// VisibilityCondition
	VisibilityCondition *Visibilitycondition `json:"visibilityCondition,omitempty"`

}

// String returns a JSON representation of the model
func (o *Surveyquestiongroup) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
