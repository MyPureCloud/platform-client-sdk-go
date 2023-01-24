package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Assessmentscoringset
type Assessmentscoringset struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// TotalScore - The total score of the answers
	TotalScore *float32 `json:"totalScore,omitempty"`

	// TotalCriticalScore - The total score for the critical questions
	TotalCriticalScore *float32 `json:"totalCriticalScore,omitempty"`

	// TotalNonCriticalScore - The total score for the non-critical questions
	TotalNonCriticalScore *float32 `json:"totalNonCriticalScore,omitempty"`

	// QuestionGroupScores - The individual scores for each question group
	QuestionGroupScores *[]Assessmentquestiongroupscore `json:"questionGroupScores,omitempty"`

	// FailureReasons - If the assessment was not passed, the reasons for failure.
	FailureReasons *[]string `json:"failureReasons,omitempty"`

	// Comments - Comments provided for these answers.
	Comments *string `json:"comments,omitempty"`

	// AgentComments - Comments provided by agent.
	AgentComments *string `json:"agentComments,omitempty"`

	// IsPassed - True if the assessment was passed
	IsPassed *bool `json:"isPassed,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Assessmentscoringset) SetField(field string, fieldValue interface{}) {
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

func (o Assessmentscoringset) MarshalJSON() ([]byte, error) {
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
			if contains(dateTimeFields, fieldName) {
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
	type Alias Assessmentscoringset
	
	return json.Marshal(&struct { 
		TotalScore *float32 `json:"totalScore,omitempty"`
		
		TotalCriticalScore *float32 `json:"totalCriticalScore,omitempty"`
		
		TotalNonCriticalScore *float32 `json:"totalNonCriticalScore,omitempty"`
		
		QuestionGroupScores *[]Assessmentquestiongroupscore `json:"questionGroupScores,omitempty"`
		
		FailureReasons *[]string `json:"failureReasons,omitempty"`
		
		Comments *string `json:"comments,omitempty"`
		
		AgentComments *string `json:"agentComments,omitempty"`
		
		IsPassed *bool `json:"isPassed,omitempty"`
		Alias
	}{ 
		TotalScore: o.TotalScore,
		
		TotalCriticalScore: o.TotalCriticalScore,
		
		TotalNonCriticalScore: o.TotalNonCriticalScore,
		
		QuestionGroupScores: o.QuestionGroupScores,
		
		FailureReasons: o.FailureReasons,
		
		Comments: o.Comments,
		
		AgentComments: o.AgentComments,
		
		IsPassed: o.IsPassed,
		Alias:    (Alias)(o),
	})
}

func (o *Assessmentscoringset) UnmarshalJSON(b []byte) error {
	var AssessmentscoringsetMap map[string]interface{}
	err := json.Unmarshal(b, &AssessmentscoringsetMap)
	if err != nil {
		return err
	}
	
	if TotalScore, ok := AssessmentscoringsetMap["totalScore"].(float64); ok {
		TotalScoreFloat32 := float32(TotalScore)
		o.TotalScore = &TotalScoreFloat32
	}
	
	if TotalCriticalScore, ok := AssessmentscoringsetMap["totalCriticalScore"].(float64); ok {
		TotalCriticalScoreFloat32 := float32(TotalCriticalScore)
		o.TotalCriticalScore = &TotalCriticalScoreFloat32
	}
	
	if TotalNonCriticalScore, ok := AssessmentscoringsetMap["totalNonCriticalScore"].(float64); ok {
		TotalNonCriticalScoreFloat32 := float32(TotalNonCriticalScore)
		o.TotalNonCriticalScore = &TotalNonCriticalScoreFloat32
	}
	
	if QuestionGroupScores, ok := AssessmentscoringsetMap["questionGroupScores"].([]interface{}); ok {
		QuestionGroupScoresString, _ := json.Marshal(QuestionGroupScores)
		json.Unmarshal(QuestionGroupScoresString, &o.QuestionGroupScores)
	}
	
	if FailureReasons, ok := AssessmentscoringsetMap["failureReasons"].([]interface{}); ok {
		FailureReasonsString, _ := json.Marshal(FailureReasons)
		json.Unmarshal(FailureReasonsString, &o.FailureReasons)
	}
	
	if Comments, ok := AssessmentscoringsetMap["comments"].(string); ok {
		o.Comments = &Comments
	}
    
	if AgentComments, ok := AssessmentscoringsetMap["agentComments"].(string); ok {
		o.AgentComments = &AgentComments
	}
    
	if IsPassed, ok := AssessmentscoringsetMap["isPassed"].(bool); ok {
		o.IsPassed = &IsPassed
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Assessmentscoringset) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
