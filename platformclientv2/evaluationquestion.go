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


	// IsCritical
	IsCritical *bool `json:"isCritical,omitempty"`


	// IsKill
	IsKill *bool `json:"isKill,omitempty"`

}

func (o *Evaluationquestion) MarshalJSON() ([]byte, error) {
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
		
		IsCritical *bool `json:"isCritical,omitempty"`
		
		IsKill *bool `json:"isKill,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Text: o.Text,
		
		HelpText: o.HelpText,
		
		VarType: o.VarType,
		
		NaEnabled: o.NaEnabled,
		
		CommentsRequired: o.CommentsRequired,
		
		VisibilityCondition: o.VisibilityCondition,
		
		AnswerOptions: o.AnswerOptions,
		
		IsCritical: o.IsCritical,
		
		IsKill: o.IsKill,
		Alias:    (*Alias)(o),
	})
}

func (o *Evaluationquestion) UnmarshalJSON(b []byte) error {
	var EvaluationquestionMap map[string]interface{}
	err := json.Unmarshal(b, &EvaluationquestionMap)
	if err != nil {
		return err
	}
	
	if Id, ok := EvaluationquestionMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Text, ok := EvaluationquestionMap["text"].(string); ok {
		o.Text = &Text
	}
	
	if HelpText, ok := EvaluationquestionMap["helpText"].(string); ok {
		o.HelpText = &HelpText
	}
	
	if VarType, ok := EvaluationquestionMap["type"].(string); ok {
		o.VarType = &VarType
	}
	
	if NaEnabled, ok := EvaluationquestionMap["naEnabled"].(bool); ok {
		o.NaEnabled = &NaEnabled
	}
	
	if CommentsRequired, ok := EvaluationquestionMap["commentsRequired"].(bool); ok {
		o.CommentsRequired = &CommentsRequired
	}
	
	if VisibilityCondition, ok := EvaluationquestionMap["visibilityCondition"].(map[string]interface{}); ok {
		VisibilityConditionString, _ := json.Marshal(VisibilityCondition)
		json.Unmarshal(VisibilityConditionString, &o.VisibilityCondition)
	}
	
	if AnswerOptions, ok := EvaluationquestionMap["answerOptions"].([]interface{}); ok {
		AnswerOptionsString, _ := json.Marshal(AnswerOptions)
		json.Unmarshal(AnswerOptionsString, &o.AnswerOptions)
	}
	
	if IsCritical, ok := EvaluationquestionMap["isCritical"].(bool); ok {
		o.IsCritical = &IsCritical
	}
	
	if IsKill, ok := EvaluationquestionMap["isKill"].(bool); ok {
		o.IsKill = &IsKill
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Evaluationquestion) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
