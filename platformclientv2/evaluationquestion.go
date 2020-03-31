package platformclientv2
import (
	"encoding/json"
)

// Evaluationquestion
type Evaluationquestion struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Text
	Text *string `json:"text,omitempty"`


	// HelpText
	HelpText *string `json:"helpText,omitempty"`


	// VarType
	VarType *string `json:"type,omitempty"`


	// NaEnabled
	NaEnabled *bool `json:"naEnabled,omitempty"`


	// CommentsRequired
	CommentsRequired *bool `json:"commentsRequired,omitempty"`


	// VisibilityCondition
	VisibilityCondition *Visibilitycondition `json:"visibilityCondition,omitempty"`


	// AnswerOptions - Options from which to choose an answer for this question. Only used by Multiple Choice type questions.
	AnswerOptions *[]Answeroption `json:"answerOptions,omitempty"`


	// IsCritical
	IsCritical *bool `json:"isCritical,omitempty"`


	// IsKill
	IsKill *bool `json:"isKill,omitempty"`

}

// String returns a JSON representation of the model
func (o *Evaluationquestion) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
