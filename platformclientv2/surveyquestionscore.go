package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Surveyquestionscore
type Surveyquestionscore struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// QuestionId
	QuestionId *string `json:"questionId,omitempty"`

	// AnswerId
	AnswerId *string `json:"answerId,omitempty"`

	// Score - Unweighted score of the question
	Score *int `json:"score,omitempty"`

	// MarkedNA
	MarkedNA *bool `json:"markedNA,omitempty"`

	// AssistedAnswerId - AnswerId found with evaluation assistance conditions
	AssistedAnswerId *string `json:"assistedAnswerId,omitempty"`

	// NpsScore
	NpsScore *int `json:"npsScore,omitempty"`

	// NpsTextAnswer
	NpsTextAnswer *string `json:"npsTextAnswer,omitempty"`

	// FreeTextAnswer
	FreeTextAnswer *string `json:"freeTextAnswer,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Surveyquestionscore) SetField(field string, fieldValue interface{}) {
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

func (o Surveyquestionscore) MarshalJSON() ([]byte, error) {
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
	type Alias Surveyquestionscore
	
	return json.Marshal(&struct { 
		QuestionId *string `json:"questionId,omitempty"`
		
		AnswerId *string `json:"answerId,omitempty"`
		
		Score *int `json:"score,omitempty"`
		
		MarkedNA *bool `json:"markedNA,omitempty"`
		
		AssistedAnswerId *string `json:"assistedAnswerId,omitempty"`
		
		NpsScore *int `json:"npsScore,omitempty"`
		
		NpsTextAnswer *string `json:"npsTextAnswer,omitempty"`
		
		FreeTextAnswer *string `json:"freeTextAnswer,omitempty"`
		Alias
	}{ 
		QuestionId: o.QuestionId,
		
		AnswerId: o.AnswerId,
		
		Score: o.Score,
		
		MarkedNA: o.MarkedNA,
		
		AssistedAnswerId: o.AssistedAnswerId,
		
		NpsScore: o.NpsScore,
		
		NpsTextAnswer: o.NpsTextAnswer,
		
		FreeTextAnswer: o.FreeTextAnswer,
		Alias:    (Alias)(o),
	})
}

func (o *Surveyquestionscore) UnmarshalJSON(b []byte) error {
	var SurveyquestionscoreMap map[string]interface{}
	err := json.Unmarshal(b, &SurveyquestionscoreMap)
	if err != nil {
		return err
	}
	
	if QuestionId, ok := SurveyquestionscoreMap["questionId"].(string); ok {
		o.QuestionId = &QuestionId
	}
    
	if AnswerId, ok := SurveyquestionscoreMap["answerId"].(string); ok {
		o.AnswerId = &AnswerId
	}
    
	if Score, ok := SurveyquestionscoreMap["score"].(float64); ok {
		ScoreInt := int(Score)
		o.Score = &ScoreInt
	}
	
	if MarkedNA, ok := SurveyquestionscoreMap["markedNA"].(bool); ok {
		o.MarkedNA = &MarkedNA
	}
    
	if AssistedAnswerId, ok := SurveyquestionscoreMap["assistedAnswerId"].(string); ok {
		o.AssistedAnswerId = &AssistedAnswerId
	}
    
	if NpsScore, ok := SurveyquestionscoreMap["npsScore"].(float64); ok {
		NpsScoreInt := int(NpsScore)
		o.NpsScore = &NpsScoreInt
	}
	
	if NpsTextAnswer, ok := SurveyquestionscoreMap["npsTextAnswer"].(string); ok {
		o.NpsTextAnswer = &NpsTextAnswer
	}
    
	if FreeTextAnswer, ok := SurveyquestionscoreMap["freeTextAnswer"].(string); ok {
		o.FreeTextAnswer = &FreeTextAnswer
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Surveyquestionscore) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
