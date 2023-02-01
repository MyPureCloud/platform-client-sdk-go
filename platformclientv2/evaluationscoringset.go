package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Evaluationscoringset
type Evaluationscoringset struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// TotalScore - Score of all questions
	TotalScore *float32 `json:"totalScore,omitempty"`

	// TotalCriticalScore - Score of only the critical questions
	TotalCriticalScore *float32 `json:"totalCriticalScore,omitempty"`

	// TotalNonCriticalScore - Score of only the non-critical questions
	TotalNonCriticalScore *float32 `json:"totalNonCriticalScore,omitempty"`

	// QuestionGroupScores
	QuestionGroupScores *[]Evaluationquestiongroupscore `json:"questionGroupScores,omitempty"`

	// AnyFailedKillQuestions - Indicates that at least one fatal question was answered without having the highest score available for the question
	AnyFailedKillQuestions *bool `json:"anyFailedKillQuestions,omitempty"`

	// Comments - Overall comments from the evaluator
	Comments *string `json:"comments,omitempty"`

	// AgentComments - Comments from the agent while reviewing evaluation results
	AgentComments *string `json:"agentComments,omitempty"`

	// TranscriptTopics - List of topics found within the conversation's transcripts
	TranscriptTopics *[]Transcripttopic `json:"transcriptTopics,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Evaluationscoringset) SetField(field string, fieldValue interface{}) {
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

func (o Evaluationscoringset) MarshalJSON() ([]byte, error) {
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
	type Alias Evaluationscoringset
	
	return json.Marshal(&struct { 
		TotalScore *float32 `json:"totalScore,omitempty"`
		
		TotalCriticalScore *float32 `json:"totalCriticalScore,omitempty"`
		
		TotalNonCriticalScore *float32 `json:"totalNonCriticalScore,omitempty"`
		
		QuestionGroupScores *[]Evaluationquestiongroupscore `json:"questionGroupScores,omitempty"`
		
		AnyFailedKillQuestions *bool `json:"anyFailedKillQuestions,omitempty"`
		
		Comments *string `json:"comments,omitempty"`
		
		AgentComments *string `json:"agentComments,omitempty"`
		
		TranscriptTopics *[]Transcripttopic `json:"transcriptTopics,omitempty"`
		Alias
	}{ 
		TotalScore: o.TotalScore,
		
		TotalCriticalScore: o.TotalCriticalScore,
		
		TotalNonCriticalScore: o.TotalNonCriticalScore,
		
		QuestionGroupScores: o.QuestionGroupScores,
		
		AnyFailedKillQuestions: o.AnyFailedKillQuestions,
		
		Comments: o.Comments,
		
		AgentComments: o.AgentComments,
		
		TranscriptTopics: o.TranscriptTopics,
		Alias:    (Alias)(o),
	})
}

func (o *Evaluationscoringset) UnmarshalJSON(b []byte) error {
	var EvaluationscoringsetMap map[string]interface{}
	err := json.Unmarshal(b, &EvaluationscoringsetMap)
	if err != nil {
		return err
	}
	
	if TotalScore, ok := EvaluationscoringsetMap["totalScore"].(float64); ok {
		TotalScoreFloat32 := float32(TotalScore)
		o.TotalScore = &TotalScoreFloat32
	}
	
	if TotalCriticalScore, ok := EvaluationscoringsetMap["totalCriticalScore"].(float64); ok {
		TotalCriticalScoreFloat32 := float32(TotalCriticalScore)
		o.TotalCriticalScore = &TotalCriticalScoreFloat32
	}
	
	if TotalNonCriticalScore, ok := EvaluationscoringsetMap["totalNonCriticalScore"].(float64); ok {
		TotalNonCriticalScoreFloat32 := float32(TotalNonCriticalScore)
		o.TotalNonCriticalScore = &TotalNonCriticalScoreFloat32
	}
	
	if QuestionGroupScores, ok := EvaluationscoringsetMap["questionGroupScores"].([]interface{}); ok {
		QuestionGroupScoresString, _ := json.Marshal(QuestionGroupScores)
		json.Unmarshal(QuestionGroupScoresString, &o.QuestionGroupScores)
	}
	
	if AnyFailedKillQuestions, ok := EvaluationscoringsetMap["anyFailedKillQuestions"].(bool); ok {
		o.AnyFailedKillQuestions = &AnyFailedKillQuestions
	}
    
	if Comments, ok := EvaluationscoringsetMap["comments"].(string); ok {
		o.Comments = &Comments
	}
    
	if AgentComments, ok := EvaluationscoringsetMap["agentComments"].(string); ok {
		o.AgentComments = &AgentComments
	}
    
	if TranscriptTopics, ok := EvaluationscoringsetMap["transcriptTopics"].([]interface{}); ok {
		TranscriptTopicsString, _ := json.Marshal(TranscriptTopics)
		json.Unmarshal(TranscriptTopicsString, &o.TranscriptTopics)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Evaluationscoringset) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
