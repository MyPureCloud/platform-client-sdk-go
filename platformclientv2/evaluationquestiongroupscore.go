package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Evaluationquestiongroupscore
type Evaluationquestiongroupscore struct { 
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

	// TotalCriticalScore - Score of only the critical questions in the group
	TotalCriticalScore *float32 `json:"totalCriticalScore,omitempty"`

	// MaxTotalCriticalScore - Maximum possible score of only the critical questions in the group
	MaxTotalCriticalScore *float32 `json:"maxTotalCriticalScore,omitempty"`

	// TotalNonCriticalScore - Score of only the non critical questions in the group
	TotalNonCriticalScore *float32 `json:"totalNonCriticalScore,omitempty"`

	// MaxTotalNonCriticalScore - Maximum possible score of only the non critical questions in the group
	MaxTotalNonCriticalScore *float32 `json:"maxTotalNonCriticalScore,omitempty"`

	// TotalScoreUnweighted - Unweighted score of all questions in the group
	TotalScoreUnweighted *float32 `json:"totalScoreUnweighted,omitempty"`

	// MaxTotalScoreUnweighted - Maximum possible unweighted score of all questions in the group
	MaxTotalScoreUnweighted *float32 `json:"maxTotalScoreUnweighted,omitempty"`

	// TotalCriticalScoreUnweighted - Unweighted score of only the critical questions in the group
	TotalCriticalScoreUnweighted *float32 `json:"totalCriticalScoreUnweighted,omitempty"`

	// MaxTotalCriticalScoreUnweighted - Maximum possible unweighted score of only the critical questions in the group
	MaxTotalCriticalScoreUnweighted *float32 `json:"maxTotalCriticalScoreUnweighted,omitempty"`

	// TotalNonCriticalScoreUnweighted - Unweighted score of only the non critical questions in the group
	TotalNonCriticalScoreUnweighted *float32 `json:"totalNonCriticalScoreUnweighted,omitempty"`

	// MaxTotalNonCriticalScoreUnweighted - Maximum possible unweighted score of only the non critical questions in the group
	MaxTotalNonCriticalScoreUnweighted *float32 `json:"maxTotalNonCriticalScoreUnweighted,omitempty"`

	// QuestionScores
	QuestionScores *[]Evaluationquestionscore `json:"questionScores,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Evaluationquestiongroupscore) SetField(field string, fieldValue interface{}) {
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

func (o Evaluationquestiongroupscore) MarshalJSON() ([]byte, error) {
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
	type Alias Evaluationquestiongroupscore
	
	return json.Marshal(&struct { 
		QuestionGroupId *string `json:"questionGroupId,omitempty"`
		
		TotalScore *float32 `json:"totalScore,omitempty"`
		
		MaxTotalScore *float32 `json:"maxTotalScore,omitempty"`
		
		MarkedNA *bool `json:"markedNA,omitempty"`
		
		SystemMarkedNA *bool `json:"systemMarkedNA,omitempty"`
		
		TotalCriticalScore *float32 `json:"totalCriticalScore,omitempty"`
		
		MaxTotalCriticalScore *float32 `json:"maxTotalCriticalScore,omitempty"`
		
		TotalNonCriticalScore *float32 `json:"totalNonCriticalScore,omitempty"`
		
		MaxTotalNonCriticalScore *float32 `json:"maxTotalNonCriticalScore,omitempty"`
		
		TotalScoreUnweighted *float32 `json:"totalScoreUnweighted,omitempty"`
		
		MaxTotalScoreUnweighted *float32 `json:"maxTotalScoreUnweighted,omitempty"`
		
		TotalCriticalScoreUnweighted *float32 `json:"totalCriticalScoreUnweighted,omitempty"`
		
		MaxTotalCriticalScoreUnweighted *float32 `json:"maxTotalCriticalScoreUnweighted,omitempty"`
		
		TotalNonCriticalScoreUnweighted *float32 `json:"totalNonCriticalScoreUnweighted,omitempty"`
		
		MaxTotalNonCriticalScoreUnweighted *float32 `json:"maxTotalNonCriticalScoreUnweighted,omitempty"`
		
		QuestionScores *[]Evaluationquestionscore `json:"questionScores,omitempty"`
		Alias
	}{ 
		QuestionGroupId: o.QuestionGroupId,
		
		TotalScore: o.TotalScore,
		
		MaxTotalScore: o.MaxTotalScore,
		
		MarkedNA: o.MarkedNA,
		
		SystemMarkedNA: o.SystemMarkedNA,
		
		TotalCriticalScore: o.TotalCriticalScore,
		
		MaxTotalCriticalScore: o.MaxTotalCriticalScore,
		
		TotalNonCriticalScore: o.TotalNonCriticalScore,
		
		MaxTotalNonCriticalScore: o.MaxTotalNonCriticalScore,
		
		TotalScoreUnweighted: o.TotalScoreUnweighted,
		
		MaxTotalScoreUnweighted: o.MaxTotalScoreUnweighted,
		
		TotalCriticalScoreUnweighted: o.TotalCriticalScoreUnweighted,
		
		MaxTotalCriticalScoreUnweighted: o.MaxTotalCriticalScoreUnweighted,
		
		TotalNonCriticalScoreUnweighted: o.TotalNonCriticalScoreUnweighted,
		
		MaxTotalNonCriticalScoreUnweighted: o.MaxTotalNonCriticalScoreUnweighted,
		
		QuestionScores: o.QuestionScores,
		Alias:    (Alias)(o),
	})
}

func (o *Evaluationquestiongroupscore) UnmarshalJSON(b []byte) error {
	var EvaluationquestiongroupscoreMap map[string]interface{}
	err := json.Unmarshal(b, &EvaluationquestiongroupscoreMap)
	if err != nil {
		return err
	}
	
	if QuestionGroupId, ok := EvaluationquestiongroupscoreMap["questionGroupId"].(string); ok {
		o.QuestionGroupId = &QuestionGroupId
	}
    
	if TotalScore, ok := EvaluationquestiongroupscoreMap["totalScore"].(float64); ok {
		TotalScoreFloat32 := float32(TotalScore)
		o.TotalScore = &TotalScoreFloat32
	}
	
	if MaxTotalScore, ok := EvaluationquestiongroupscoreMap["maxTotalScore"].(float64); ok {
		MaxTotalScoreFloat32 := float32(MaxTotalScore)
		o.MaxTotalScore = &MaxTotalScoreFloat32
	}
	
	if MarkedNA, ok := EvaluationquestiongroupscoreMap["markedNA"].(bool); ok {
		o.MarkedNA = &MarkedNA
	}
    
	if SystemMarkedNA, ok := EvaluationquestiongroupscoreMap["systemMarkedNA"].(bool); ok {
		o.SystemMarkedNA = &SystemMarkedNA
	}
    
	if TotalCriticalScore, ok := EvaluationquestiongroupscoreMap["totalCriticalScore"].(float64); ok {
		TotalCriticalScoreFloat32 := float32(TotalCriticalScore)
		o.TotalCriticalScore = &TotalCriticalScoreFloat32
	}
	
	if MaxTotalCriticalScore, ok := EvaluationquestiongroupscoreMap["maxTotalCriticalScore"].(float64); ok {
		MaxTotalCriticalScoreFloat32 := float32(MaxTotalCriticalScore)
		o.MaxTotalCriticalScore = &MaxTotalCriticalScoreFloat32
	}
	
	if TotalNonCriticalScore, ok := EvaluationquestiongroupscoreMap["totalNonCriticalScore"].(float64); ok {
		TotalNonCriticalScoreFloat32 := float32(TotalNonCriticalScore)
		o.TotalNonCriticalScore = &TotalNonCriticalScoreFloat32
	}
	
	if MaxTotalNonCriticalScore, ok := EvaluationquestiongroupscoreMap["maxTotalNonCriticalScore"].(float64); ok {
		MaxTotalNonCriticalScoreFloat32 := float32(MaxTotalNonCriticalScore)
		o.MaxTotalNonCriticalScore = &MaxTotalNonCriticalScoreFloat32
	}
	
	if TotalScoreUnweighted, ok := EvaluationquestiongroupscoreMap["totalScoreUnweighted"].(float64); ok {
		TotalScoreUnweightedFloat32 := float32(TotalScoreUnweighted)
		o.TotalScoreUnweighted = &TotalScoreUnweightedFloat32
	}
	
	if MaxTotalScoreUnweighted, ok := EvaluationquestiongroupscoreMap["maxTotalScoreUnweighted"].(float64); ok {
		MaxTotalScoreUnweightedFloat32 := float32(MaxTotalScoreUnweighted)
		o.MaxTotalScoreUnweighted = &MaxTotalScoreUnweightedFloat32
	}
	
	if TotalCriticalScoreUnweighted, ok := EvaluationquestiongroupscoreMap["totalCriticalScoreUnweighted"].(float64); ok {
		TotalCriticalScoreUnweightedFloat32 := float32(TotalCriticalScoreUnweighted)
		o.TotalCriticalScoreUnweighted = &TotalCriticalScoreUnweightedFloat32
	}
	
	if MaxTotalCriticalScoreUnweighted, ok := EvaluationquestiongroupscoreMap["maxTotalCriticalScoreUnweighted"].(float64); ok {
		MaxTotalCriticalScoreUnweightedFloat32 := float32(MaxTotalCriticalScoreUnweighted)
		o.MaxTotalCriticalScoreUnweighted = &MaxTotalCriticalScoreUnweightedFloat32
	}
	
	if TotalNonCriticalScoreUnweighted, ok := EvaluationquestiongroupscoreMap["totalNonCriticalScoreUnweighted"].(float64); ok {
		TotalNonCriticalScoreUnweightedFloat32 := float32(TotalNonCriticalScoreUnweighted)
		o.TotalNonCriticalScoreUnweighted = &TotalNonCriticalScoreUnweightedFloat32
	}
	
	if MaxTotalNonCriticalScoreUnweighted, ok := EvaluationquestiongroupscoreMap["maxTotalNonCriticalScoreUnweighted"].(float64); ok {
		MaxTotalNonCriticalScoreUnweightedFloat32 := float32(MaxTotalNonCriticalScoreUnweighted)
		o.MaxTotalNonCriticalScoreUnweighted = &MaxTotalNonCriticalScoreUnweightedFloat32
	}
	
	if QuestionScores, ok := EvaluationquestiongroupscoreMap["questionScores"].([]interface{}); ok {
		QuestionScoresString, _ := json.Marshal(QuestionScores)
		json.Unmarshal(QuestionScoresString, &o.QuestionScores)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Evaluationquestiongroupscore) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
