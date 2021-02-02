package platformclientv2
import (
	"encoding/json"
)

// Surveyquestion
type Surveyquestion struct { 
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


	// VisibilityCondition
	VisibilityCondition *Visibilitycondition `json:"visibilityCondition,omitempty"`


	// AnswerOptions - Options from which to choose an answer for this question. Only used by Multiple Choice type questions.
	AnswerOptions *[]Answeroption `json:"answerOptions,omitempty"`


	// MaxResponseCharacters - How many characters are allowed in the text response to this question. Used by NPS and Free Text question types.
	MaxResponseCharacters *int `json:"maxResponseCharacters,omitempty"`


	// ExplanationPrompt - Prompt for details explaining the chosen NPS score. Used by NPS questions.
	ExplanationPrompt *string `json:"explanationPrompt,omitempty"`

}

// String returns a JSON representation of the model
func (o *Surveyquestion) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
