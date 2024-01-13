package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Assessmentquestionscore
type Assessmentquestionscore struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// FailedKillQuestion - True if this was a failed Kill question
	FailedKillQuestion *bool `json:"failedKillQuestion,omitempty"`

	// Comments - Comments provided for the answer
	Comments *string `json:"comments,omitempty"`

	// QuestionId - The ID of the question
	QuestionId *string `json:"questionId,omitempty"`

	// AnswerId - The ID of the selected answer
	AnswerId *string `json:"answerId,omitempty"`

	// Score - The score received for this question
	Score *int `json:"score,omitempty"`

	// MarkedNA - True if this question was marked as NA
	MarkedNA *bool `json:"markedNA,omitempty"`

	// SystemMarkedNA - If markedNA is true, systemMarkedNA indicates whether it was marked by a user or by the system due to visibility conditions. Always false if markedNA is false.
	SystemMarkedNA *bool `json:"systemMarkedNA,omitempty"`

	// FreeTextAnswer - Answer for free text answer type
	FreeTextAnswer *string `json:"freeTextAnswer,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Assessmentquestionscore) SetField(field string, fieldValue interface{}) {
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

func (o Assessmentquestionscore) MarshalJSON() ([]byte, error) {
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
	type Alias Assessmentquestionscore
	
	return json.Marshal(&struct { 
		FailedKillQuestion *bool `json:"failedKillQuestion,omitempty"`
		
		Comments *string `json:"comments,omitempty"`
		
		QuestionId *string `json:"questionId,omitempty"`
		
		AnswerId *string `json:"answerId,omitempty"`
		
		Score *int `json:"score,omitempty"`
		
		MarkedNA *bool `json:"markedNA,omitempty"`
		
		SystemMarkedNA *bool `json:"systemMarkedNA,omitempty"`
		
		FreeTextAnswer *string `json:"freeTextAnswer,omitempty"`
		Alias
	}{ 
		FailedKillQuestion: o.FailedKillQuestion,
		
		Comments: o.Comments,
		
		QuestionId: o.QuestionId,
		
		AnswerId: o.AnswerId,
		
		Score: o.Score,
		
		MarkedNA: o.MarkedNA,
		
		SystemMarkedNA: o.SystemMarkedNA,
		
		FreeTextAnswer: o.FreeTextAnswer,
		Alias:    (Alias)(o),
	})
}

func (o *Assessmentquestionscore) UnmarshalJSON(b []byte) error {
	var AssessmentquestionscoreMap map[string]interface{}
	err := json.Unmarshal(b, &AssessmentquestionscoreMap)
	if err != nil {
		return err
	}
	
	if FailedKillQuestion, ok := AssessmentquestionscoreMap["failedKillQuestion"].(bool); ok {
		o.FailedKillQuestion = &FailedKillQuestion
	}
    
	if Comments, ok := AssessmentquestionscoreMap["comments"].(string); ok {
		o.Comments = &Comments
	}
    
	if QuestionId, ok := AssessmentquestionscoreMap["questionId"].(string); ok {
		o.QuestionId = &QuestionId
	}
    
	if AnswerId, ok := AssessmentquestionscoreMap["answerId"].(string); ok {
		o.AnswerId = &AnswerId
	}
    
	if Score, ok := AssessmentquestionscoreMap["score"].(float64); ok {
		ScoreInt := int(Score)
		o.Score = &ScoreInt
	}
	
	if MarkedNA, ok := AssessmentquestionscoreMap["markedNA"].(bool); ok {
		o.MarkedNA = &MarkedNA
	}
    
	if SystemMarkedNA, ok := AssessmentquestionscoreMap["systemMarkedNA"].(bool); ok {
		o.SystemMarkedNA = &SystemMarkedNA
	}
    
	if FreeTextAnswer, ok := AssessmentquestionscoreMap["freeTextAnswer"].(string); ok {
		o.FreeTextAnswer = &FreeTextAnswer
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Assessmentquestionscore) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
