package platformclientv2
import (
	"github.com/leekchan/timeutil"
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


	// Text - The question text
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

func (u *Assessmentformquestion) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Assessmentformquestion

	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		
		Text *string `json:"text,omitempty"`
		
		HelpText *string `json:"helpText,omitempty"`
		
		NaEnabled *bool `json:"naEnabled,omitempty"`
		
		CommentsRequired *bool `json:"commentsRequired,omitempty"`
		
		VisibilityCondition *Visibilitycondition `json:"visibilityCondition,omitempty"`
		
		AnswerOptions *[]Answeroption `json:"answerOptions,omitempty"`
		
		MaxResponseCharacters *int `json:"maxResponseCharacters,omitempty"`
		
		IsKill *bool `json:"isKill,omitempty"`
		
		IsCritical *bool `json:"isCritical,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		VarType: u.VarType,
		
		Text: u.Text,
		
		HelpText: u.HelpText,
		
		NaEnabled: u.NaEnabled,
		
		CommentsRequired: u.CommentsRequired,
		
		VisibilityCondition: u.VisibilityCondition,
		
		AnswerOptions: u.AnswerOptions,
		
		MaxResponseCharacters: u.MaxResponseCharacters,
		
		IsKill: u.IsKill,
		
		IsCritical: u.IsCritical,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Assessmentformquestion) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
