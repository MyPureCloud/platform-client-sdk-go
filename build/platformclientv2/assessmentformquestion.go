package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Assessmentformquestion
type Assessmentformquestion struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// VarType
	VarType *string `json:"type,omitempty"`


	// Text
	Text *string `json:"text,omitempty"`


	// HelpText
	HelpText *string `json:"helpText,omitempty"`


	// NaEnabled
	NaEnabled *bool `json:"naEnabled,omitempty"`


	// CommentsRequired
	CommentsRequired *bool `json:"commentsRequired,omitempty"`


	// VisibilityCondition
	VisibilityCondition *Visibilitycondition `json:"visibilityCondition,omitempty"`


	// AnswerOptions - Options from which to choose an answer for this question. Only used by Multiple Choice type questions.
	AnswerOptions *[]Answeroption `json:"answerOptions,omitempty"`


	// MaxResponseCharacters - How many characters are allowed in the text response to this question. Used by Free Text question types.
	MaxResponseCharacters *int `json:"maxResponseCharacters,omitempty"`


	// IsKill - Does an incorrect answer to this question mark the form as having a failed kill question. Only used by Multiple Choice type questions.
	IsKill *bool `json:"isKill,omitempty"`


	// IsCritical - Does this question contribute to the critical score. Only used by Multiple Choice type questions.
	IsCritical *bool `json:"isCritical,omitempty"`

}

// String returns a JSON representation of the model
func (o *Assessmentformquestion) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
