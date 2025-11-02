package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Evaluationquestion
type Evaluationquestion struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id
	Id *string `json:"id,omitempty"`

	// ContextId - An identifier for this question that stays the same across versions of the form.
	ContextId *string `json:"contextId,omitempty"`

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

	// MultipleSelectOptionQuestions - Only used by Multiple Select type questions. A list of multiple choice questions representing selectable options.
	MultipleSelectOptionQuestions *[]Evaluationquestion `json:"multipleSelectOptionQuestions,omitempty"`

	// DefaultAnswer - The default selected answer for the question
	DefaultAnswer *Defaultanswer `json:"defaultAnswer,omitempty"`

	// IsKill
	IsKill *bool `json:"isKill,omitempty"`

	// IsCritical
	IsCritical *bool `json:"isCritical,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Evaluationquestion) SetField(field string, fieldValue interface{}) {
	// Get Value object for field
	target := reflect.ValueOf(o)
	targetField := reflect.Indirect(target).FieldByName(field)

	// Set value
	if fieldValue != nil {
		targetField.Set(reflect.ValueOf(fieldValue))
	} else {
		// Must create a new Value (creates **type) then get its element (*type), which will be nil pointer of the appropriate type
		x := reflect.Indirect(reflect.New(targetField.Type()))
		targetField.Set(x)
	}

	// Add field to set field names list
	if o.SetFieldNames == nil {
		o.SetFieldNames = make(map[string]bool)
	}
	o.SetFieldNames[field] = true
}

func (o Evaluationquestion) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{  }
		localDateTimeFields := []string{  }
		dateFields := []string{  }

		// Construct object
		newObj := make(map[string]interface{})
		for fieldName := range o.SetFieldNames {
			// Get initial field value
			fieldValue := val.FieldByName(fieldName).Interface()

			// Apply value formatting overrides
			if fieldValue == nil || reflect.ValueOf(fieldValue).IsNil()  {
				// Do nothing. Just catching this case to avoid trying to custom serialize a nil value
			} else if contains(dateTimeFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%dT%H:%M:%S.%fZ")
			} else if contains(localDateTimeFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%dT%H:%M:%S.%f")
			} else if contains(dateFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%d")
			}

			// Assign value to field using JSON tag name
			newObj[getFieldName(reflect.TypeOf(&o), fieldName)] = fieldValue
		}

		// Marshal and return dynamically constructed interface
		return json.Marshal(newObj)
	}

	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Evaluationquestion
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		ContextId *string `json:"contextId,omitempty"`
		
		Text *string `json:"text,omitempty"`
		
		HelpText *string `json:"helpText,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		
		NaEnabled *bool `json:"naEnabled,omitempty"`
		
		CommentsRequired *bool `json:"commentsRequired,omitempty"`
		
		VisibilityCondition *Visibilitycondition `json:"visibilityCondition,omitempty"`
		
		AnswerOptions *[]Answeroption `json:"answerOptions,omitempty"`
		
		MultipleSelectOptionQuestions *[]Evaluationquestion `json:"multipleSelectOptionQuestions,omitempty"`
		
		DefaultAnswer *Defaultanswer `json:"defaultAnswer,omitempty"`
		
		IsKill *bool `json:"isKill,omitempty"`
		
		IsCritical *bool `json:"isCritical,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		ContextId: o.ContextId,
		
		Text: o.Text,
		
		HelpText: o.HelpText,
		
		VarType: o.VarType,
		
		NaEnabled: o.NaEnabled,
		
		CommentsRequired: o.CommentsRequired,
		
		VisibilityCondition: o.VisibilityCondition,
		
		AnswerOptions: o.AnswerOptions,
		
		MultipleSelectOptionQuestions: o.MultipleSelectOptionQuestions,
		
		DefaultAnswer: o.DefaultAnswer,
		
		IsKill: o.IsKill,
		
		IsCritical: o.IsCritical,
		Alias:    (Alias)(o),
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
    
	if ContextId, ok := EvaluationquestionMap["contextId"].(string); ok {
		o.ContextId = &ContextId
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
	
	if MultipleSelectOptionQuestions, ok := EvaluationquestionMap["multipleSelectOptionQuestions"].([]interface{}); ok {
		MultipleSelectOptionQuestionsString, _ := json.Marshal(MultipleSelectOptionQuestions)
		json.Unmarshal(MultipleSelectOptionQuestionsString, &o.MultipleSelectOptionQuestions)
	}
	
	if DefaultAnswer, ok := EvaluationquestionMap["defaultAnswer"].(map[string]interface{}); ok {
		DefaultAnswerString, _ := json.Marshal(DefaultAnswer)
		json.Unmarshal(DefaultAnswerString, &o.DefaultAnswer)
	}
	
	if IsKill, ok := EvaluationquestionMap["isKill"].(bool); ok {
		o.IsKill = &IsKill
	}
    
	if IsCritical, ok := EvaluationquestionMap["isCritical"].(bool); ok {
		o.IsCritical = &IsCritical
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Evaluationquestion) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
