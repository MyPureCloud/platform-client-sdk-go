package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
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

func (u *Surveyquestion) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Surveyquestion

	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Text *string `json:"text,omitempty"`
		
		HelpText *string `json:"helpText,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		
		NaEnabled *bool `json:"naEnabled,omitempty"`
		
		VisibilityCondition *Visibilitycondition `json:"visibilityCondition,omitempty"`
		
		AnswerOptions *[]Answeroption `json:"answerOptions,omitempty"`
		
		MaxResponseCharacters *int `json:"maxResponseCharacters,omitempty"`
		
		ExplanationPrompt *string `json:"explanationPrompt,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Text: u.Text,
		
		HelpText: u.HelpText,
		
		VarType: u.VarType,
		
		NaEnabled: u.NaEnabled,
		
		VisibilityCondition: u.VisibilityCondition,
		
		AnswerOptions: u.AnswerOptions,
		
		MaxResponseCharacters: u.MaxResponseCharacters,
		
		ExplanationPrompt: u.ExplanationPrompt,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Surveyquestion) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
