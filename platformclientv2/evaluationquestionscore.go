package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Evaluationquestionscore
type Evaluationquestionscore struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// QuestionId
	QuestionId *string `json:"questionId,omitempty"`

	// AnswerId
	AnswerId *string `json:"answerId,omitempty"`

	// Score - Unweighted score of the question
	Score *int `json:"score,omitempty"`

	// MarkedNA - True when the evaluation is submitted with a question that does not have an answer. Only allowed when naEnabled is true or if set by the system
	MarkedNA *bool `json:"markedNA,omitempty"`

	// SystemMarkedNA - If markedNA is true, systemMarkedNA indicates whether it was marked by a user or by the system due to visibility conditions. Always false if markedNA is false.
	SystemMarkedNA *bool `json:"systemMarkedNA,omitempty"`

	// AssistedAnswerId - AnswerId found with evaluation assistance conditions
	AssistedAnswerId *string `json:"assistedAnswerId,omitempty"`

	// FailedKillQuestion - Applicable only on fatal questions. Indicates that the answer selected was not the highest score available for the question
	FailedKillQuestion *bool `json:"failedKillQuestion,omitempty"`

	// Comments - Comments from the evaluator specific to this question
	Comments *string `json:"comments,omitempty"`

	// AiAnswer - Suggested AI answer
	AiAnswer *Aianswer `json:"aiAnswer,omitempty"`

	// MultipleSelectQuestionOptionScores - Only applicable to Multiple Select questions. Scores corresponding to the options of Multiple Select questions.
	MultipleSelectQuestionOptionScores *[]Evaluationquestionscore `json:"multipleSelectQuestionOptionScores,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Evaluationquestionscore) SetField(field string, fieldValue interface{}) {
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

func (o Evaluationquestionscore) MarshalJSON() ([]byte, error) {
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
	type Alias Evaluationquestionscore
	
	return json.Marshal(&struct { 
		QuestionId *string `json:"questionId,omitempty"`
		
		AnswerId *string `json:"answerId,omitempty"`
		
		Score *int `json:"score,omitempty"`
		
		MarkedNA *bool `json:"markedNA,omitempty"`
		
		SystemMarkedNA *bool `json:"systemMarkedNA,omitempty"`
		
		AssistedAnswerId *string `json:"assistedAnswerId,omitempty"`
		
		FailedKillQuestion *bool `json:"failedKillQuestion,omitempty"`
		
		Comments *string `json:"comments,omitempty"`
		
		AiAnswer *Aianswer `json:"aiAnswer,omitempty"`
		
		MultipleSelectQuestionOptionScores *[]Evaluationquestionscore `json:"multipleSelectQuestionOptionScores,omitempty"`
		Alias
	}{ 
		QuestionId: o.QuestionId,
		
		AnswerId: o.AnswerId,
		
		Score: o.Score,
		
		MarkedNA: o.MarkedNA,
		
		SystemMarkedNA: o.SystemMarkedNA,
		
		AssistedAnswerId: o.AssistedAnswerId,
		
		FailedKillQuestion: o.FailedKillQuestion,
		
		Comments: o.Comments,
		
		AiAnswer: o.AiAnswer,
		
		MultipleSelectQuestionOptionScores: o.MultipleSelectQuestionOptionScores,
		Alias:    (Alias)(o),
	})
}

func (o *Evaluationquestionscore) UnmarshalJSON(b []byte) error {
	var EvaluationquestionscoreMap map[string]interface{}
	err := json.Unmarshal(b, &EvaluationquestionscoreMap)
	if err != nil {
		return err
	}
	
	if QuestionId, ok := EvaluationquestionscoreMap["questionId"].(string); ok {
		o.QuestionId = &QuestionId
	}
    
	if AnswerId, ok := EvaluationquestionscoreMap["answerId"].(string); ok {
		o.AnswerId = &AnswerId
	}
    
	if Score, ok := EvaluationquestionscoreMap["score"].(float64); ok {
		ScoreInt := int(Score)
		o.Score = &ScoreInt
	}
	
	if MarkedNA, ok := EvaluationquestionscoreMap["markedNA"].(bool); ok {
		o.MarkedNA = &MarkedNA
	}
    
	if SystemMarkedNA, ok := EvaluationquestionscoreMap["systemMarkedNA"].(bool); ok {
		o.SystemMarkedNA = &SystemMarkedNA
	}
    
	if AssistedAnswerId, ok := EvaluationquestionscoreMap["assistedAnswerId"].(string); ok {
		o.AssistedAnswerId = &AssistedAnswerId
	}
    
	if FailedKillQuestion, ok := EvaluationquestionscoreMap["failedKillQuestion"].(bool); ok {
		o.FailedKillQuestion = &FailedKillQuestion
	}
    
	if Comments, ok := EvaluationquestionscoreMap["comments"].(string); ok {
		o.Comments = &Comments
	}
    
	if AiAnswer, ok := EvaluationquestionscoreMap["aiAnswer"].(map[string]interface{}); ok {
		AiAnswerString, _ := json.Marshal(AiAnswer)
		json.Unmarshal(AiAnswerString, &o.AiAnswer)
	}
	
	if MultipleSelectQuestionOptionScores, ok := EvaluationquestionscoreMap["multipleSelectQuestionOptionScores"].([]interface{}); ok {
		MultipleSelectQuestionOptionScoresString, _ := json.Marshal(MultipleSelectQuestionOptionScores)
		json.Unmarshal(MultipleSelectQuestionOptionScoresString, &o.MultipleSelectQuestionOptionScores)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Evaluationquestionscore) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
