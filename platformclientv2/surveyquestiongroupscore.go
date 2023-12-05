package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Surveyquestiongroupscore
type Surveyquestiongroupscore struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// QuestionGroupId
	QuestionGroupId *string `json:"questionGroupId,omitempty"`

	// TotalScore - Score of all questions in the group
	TotalScore *float32 `json:"totalScore,omitempty"`

	// MaxTotalScore - Maximum possible score of all questions in the group
	MaxTotalScore *float32 `json:"maxTotalScore,omitempty"`

	// MarkedNA - True when the evaluation is submitted with a question group that does not have any answers. Only allowed when naEnabled is true or if set by the system
	MarkedNA *bool `json:"markedNA,omitempty"`

	// SystemMarkedNA - If markedNA is true, systemMarkedNA indicates whether it was marked by a user or by the system due to visibility conditions. Always false if markedNA is false.
	SystemMarkedNA *bool `json:"systemMarkedNA,omitempty"`

	// QuestionScores
	QuestionScores *[]Surveyquestionscore `json:"questionScores,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Surveyquestiongroupscore) SetField(field string, fieldValue interface{}) {
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

func (o Surveyquestiongroupscore) MarshalJSON() ([]byte, error) {
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
	type Alias Surveyquestiongroupscore
	
	return json.Marshal(&struct { 
		QuestionGroupId *string `json:"questionGroupId,omitempty"`
		
		TotalScore *float32 `json:"totalScore,omitempty"`
		
		MaxTotalScore *float32 `json:"maxTotalScore,omitempty"`
		
		MarkedNA *bool `json:"markedNA,omitempty"`
		
		SystemMarkedNA *bool `json:"systemMarkedNA,omitempty"`
		
		QuestionScores *[]Surveyquestionscore `json:"questionScores,omitempty"`
		Alias
	}{ 
		QuestionGroupId: o.QuestionGroupId,
		
		TotalScore: o.TotalScore,
		
		MaxTotalScore: o.MaxTotalScore,
		
		MarkedNA: o.MarkedNA,
		
		SystemMarkedNA: o.SystemMarkedNA,
		
		QuestionScores: o.QuestionScores,
		Alias:    (Alias)(o),
	})
}

func (o *Surveyquestiongroupscore) UnmarshalJSON(b []byte) error {
	var SurveyquestiongroupscoreMap map[string]interface{}
	err := json.Unmarshal(b, &SurveyquestiongroupscoreMap)
	if err != nil {
		return err
	}
	
	if QuestionGroupId, ok := SurveyquestiongroupscoreMap["questionGroupId"].(string); ok {
		o.QuestionGroupId = &QuestionGroupId
	}
    
	if TotalScore, ok := SurveyquestiongroupscoreMap["totalScore"].(float64); ok {
		TotalScoreFloat32 := float32(TotalScore)
		o.TotalScore = &TotalScoreFloat32
	}
	
	if MaxTotalScore, ok := SurveyquestiongroupscoreMap["maxTotalScore"].(float64); ok {
		MaxTotalScoreFloat32 := float32(MaxTotalScore)
		o.MaxTotalScore = &MaxTotalScoreFloat32
	}
	
	if MarkedNA, ok := SurveyquestiongroupscoreMap["markedNA"].(bool); ok {
		o.MarkedNA = &MarkedNA
	}
    
	if SystemMarkedNA, ok := SurveyquestiongroupscoreMap["systemMarkedNA"].(bool); ok {
		o.SystemMarkedNA = &SystemMarkedNA
	}
    
	if QuestionScores, ok := SurveyquestiongroupscoreMap["questionScores"].([]interface{}); ok {
		QuestionScoresString, _ := json.Marshal(QuestionScores)
		json.Unmarshal(QuestionScoresString, &o.QuestionScores)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Surveyquestiongroupscore) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
