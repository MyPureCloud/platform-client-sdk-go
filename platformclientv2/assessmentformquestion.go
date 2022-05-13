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

func (o *Assessmentformquestion) MarshalJSON() ([]byte, error) {
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
		Id: o.Id,
		
		VarType: o.VarType,
		
		Text: o.Text,
		
		HelpText: o.HelpText,
		
		NaEnabled: o.NaEnabled,
		
		CommentsRequired: o.CommentsRequired,
		
		VisibilityCondition: o.VisibilityCondition,
		
		AnswerOptions: o.AnswerOptions,
		
		MaxResponseCharacters: o.MaxResponseCharacters,
		
		IsKill: o.IsKill,
		
		IsCritical: o.IsCritical,
		Alias:    (*Alias)(o),
	})
}

func (o *Assessmentformquestion) UnmarshalJSON(b []byte) error {
	var AssessmentformquestionMap map[string]interface{}
	err := json.Unmarshal(b, &AssessmentformquestionMap)
	if err != nil {
		return err
	}
	
	if Id, ok := AssessmentformquestionMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if VarType, ok := AssessmentformquestionMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if Text, ok := AssessmentformquestionMap["text"].(string); ok {
		o.Text = &Text
	}
    
	if HelpText, ok := AssessmentformquestionMap["helpText"].(string); ok {
		o.HelpText = &HelpText
	}
    
	if NaEnabled, ok := AssessmentformquestionMap["naEnabled"].(bool); ok {
		o.NaEnabled = &NaEnabled
	}
    
	if CommentsRequired, ok := AssessmentformquestionMap["commentsRequired"].(bool); ok {
		o.CommentsRequired = &CommentsRequired
	}
    
	if VisibilityCondition, ok := AssessmentformquestionMap["visibilityCondition"].(map[string]interface{}); ok {
		VisibilityConditionString, _ := json.Marshal(VisibilityCondition)
		json.Unmarshal(VisibilityConditionString, &o.VisibilityCondition)
	}
	
	if AnswerOptions, ok := AssessmentformquestionMap["answerOptions"].([]interface{}); ok {
		AnswerOptionsString, _ := json.Marshal(AnswerOptions)
		json.Unmarshal(AnswerOptionsString, &o.AnswerOptions)
	}
	
	if MaxResponseCharacters, ok := AssessmentformquestionMap["maxResponseCharacters"].(float64); ok {
		MaxResponseCharactersInt := int(MaxResponseCharacters)
		o.MaxResponseCharacters = &MaxResponseCharactersInt
	}
	
	if IsKill, ok := AssessmentformquestionMap["isKill"].(bool); ok {
		o.IsKill = &IsKill
	}
    
	if IsCritical, ok := AssessmentformquestionMap["isCritical"].(bool); ok {
		o.IsCritical = &IsCritical
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Assessmentformquestion) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
