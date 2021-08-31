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

func (o *Surveyquestion) MarshalJSON() ([]byte, error) {
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
		Id: o.Id,
		
		Text: o.Text,
		
		HelpText: o.HelpText,
		
		VarType: o.VarType,
		
		NaEnabled: o.NaEnabled,
		
		VisibilityCondition: o.VisibilityCondition,
		
		AnswerOptions: o.AnswerOptions,
		
		MaxResponseCharacters: o.MaxResponseCharacters,
		
		ExplanationPrompt: o.ExplanationPrompt,
		Alias:    (*Alias)(o),
	})
}

func (o *Surveyquestion) UnmarshalJSON(b []byte) error {
	var SurveyquestionMap map[string]interface{}
	err := json.Unmarshal(b, &SurveyquestionMap)
	if err != nil {
		return err
	}
	
	if Id, ok := SurveyquestionMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Text, ok := SurveyquestionMap["text"].(string); ok {
		o.Text = &Text
	}
	
	if HelpText, ok := SurveyquestionMap["helpText"].(string); ok {
		o.HelpText = &HelpText
	}
	
	if VarType, ok := SurveyquestionMap["type"].(string); ok {
		o.VarType = &VarType
	}
	
	if NaEnabled, ok := SurveyquestionMap["naEnabled"].(bool); ok {
		o.NaEnabled = &NaEnabled
	}
	
	if VisibilityCondition, ok := SurveyquestionMap["visibilityCondition"].(map[string]interface{}); ok {
		VisibilityConditionString, _ := json.Marshal(VisibilityCondition)
		json.Unmarshal(VisibilityConditionString, &o.VisibilityCondition)
	}
	
	if AnswerOptions, ok := SurveyquestionMap["answerOptions"].([]interface{}); ok {
		AnswerOptionsString, _ := json.Marshal(AnswerOptions)
		json.Unmarshal(AnswerOptionsString, &o.AnswerOptions)
	}
	
	if MaxResponseCharacters, ok := SurveyquestionMap["maxResponseCharacters"].(float64); ok {
		MaxResponseCharactersInt := int(MaxResponseCharacters)
		o.MaxResponseCharacters = &MaxResponseCharactersInt
	}
	
	if ExplanationPrompt, ok := SurveyquestionMap["explanationPrompt"].(string); ok {
		o.ExplanationPrompt = &ExplanationPrompt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Surveyquestion) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
