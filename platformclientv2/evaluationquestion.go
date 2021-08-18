package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
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


	// IsKill
	IsKill *bool `json:"isKill,omitempty"`


	// IsCritical
	IsCritical *bool `json:"isCritical,omitempty"`

}

func (u *Evaluationquestion) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Evaluationquestion

	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Text *string `json:"text,omitempty"`
		
		HelpText *string `json:"helpText,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		
		NaEnabled *bool `json:"naEnabled,omitempty"`
		
		CommentsRequired *bool `json:"commentsRequired,omitempty"`
		
		VisibilityCondition *Visibilitycondition `json:"visibilityCondition,omitempty"`
		
		AnswerOptions *[]Answeroption `json:"answerOptions,omitempty"`
		
		IsKill *bool `json:"isKill,omitempty"`
		
		IsCritical *bool `json:"isCritical,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Text: u.Text,
		
		HelpText: u.HelpText,
		
		VarType: u.VarType,
		
		NaEnabled: u.NaEnabled,
		
		CommentsRequired: u.CommentsRequired,
		
		VisibilityCondition: u.VisibilityCondition,
		
		AnswerOptions: u.AnswerOptions,
		
		IsKill: u.IsKill,
		
		IsCritical: u.IsCritical,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Evaluationquestion) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
