package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Assessmentformquestion
type Assessmentformquestion struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id
	Id *string `json:"id,omitempty"`

	// ContextId - An identifier for this question that stays the same across versions of the form.
	ContextId *string `json:"contextId,omitempty"`

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

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Assessmentformquestion) SetField(field string, fieldValue interface{}) {
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

func (o Assessmentformquestion) MarshalJSON() ([]byte, error) {
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
	type Alias Assessmentformquestion
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		ContextId *string `json:"contextId,omitempty"`
		
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
		Alias
	}{ 
		Id: o.Id,
		
		ContextId: o.ContextId,
		
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
		Alias:    (Alias)(o),
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
    
	if ContextId, ok := AssessmentformquestionMap["contextId"].(string); ok {
		o.ContextId = &ContextId
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
