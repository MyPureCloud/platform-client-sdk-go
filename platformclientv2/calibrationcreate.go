package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Calibrationcreate
type Calibrationcreate struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// Name
	Name *string `json:"name,omitempty"`

	// Calibrator
	Calibrator *User `json:"calibrator,omitempty"`

	// Agent
	Agent *User `json:"agent,omitempty"`

	// Conversation - The conversation to use for the calibration.
	Conversation *Conversationreference `json:"conversation,omitempty"`

	// EvaluationForm
	EvaluationForm *Evaluationform `json:"evaluationForm,omitempty"`

	// ContextId
	ContextId *string `json:"contextId,omitempty"`

	// AverageScore
	AverageScore *int `json:"averageScore,omitempty"`

	// HighScore
	HighScore *int `json:"highScore,omitempty"`

	// LowScore
	LowScore *int `json:"lowScore,omitempty"`

	// CreatedDate - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	CreatedDate *time.Time `json:"createdDate,omitempty"`

	// Evaluations
	Evaluations *[]Evaluation `json:"evaluations,omitempty"`

	// Evaluators
	Evaluators *[]User `json:"evaluators,omitempty"`

	// ScoringIndex
	ScoringIndex *Evaluation `json:"scoringIndex,omitempty"`

	// ExpertEvaluator
	ExpertEvaluator *User `json:"expertEvaluator,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Calibrationcreate) SetField(field string, fieldValue interface{}) {
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

func (o Calibrationcreate) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "CreatedDate", }
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
	type Alias Calibrationcreate
	
	CreatedDate := new(string)
	if o.CreatedDate != nil {
		
		*CreatedDate = timeutil.Strftime(o.CreatedDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		CreatedDate = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Calibrator *User `json:"calibrator,omitempty"`
		
		Agent *User `json:"agent,omitempty"`
		
		Conversation *Conversationreference `json:"conversation,omitempty"`
		
		EvaluationForm *Evaluationform `json:"evaluationForm,omitempty"`
		
		ContextId *string `json:"contextId,omitempty"`
		
		AverageScore *int `json:"averageScore,omitempty"`
		
		HighScore *int `json:"highScore,omitempty"`
		
		LowScore *int `json:"lowScore,omitempty"`
		
		CreatedDate *string `json:"createdDate,omitempty"`
		
		Evaluations *[]Evaluation `json:"evaluations,omitempty"`
		
		Evaluators *[]User `json:"evaluators,omitempty"`
		
		ScoringIndex *Evaluation `json:"scoringIndex,omitempty"`
		
		ExpertEvaluator *User `json:"expertEvaluator,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Calibrator: o.Calibrator,
		
		Agent: o.Agent,
		
		Conversation: o.Conversation,
		
		EvaluationForm: o.EvaluationForm,
		
		ContextId: o.ContextId,
		
		AverageScore: o.AverageScore,
		
		HighScore: o.HighScore,
		
		LowScore: o.LowScore,
		
		CreatedDate: CreatedDate,
		
		Evaluations: o.Evaluations,
		
		Evaluators: o.Evaluators,
		
		ScoringIndex: o.ScoringIndex,
		
		ExpertEvaluator: o.ExpertEvaluator,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Calibrationcreate) UnmarshalJSON(b []byte) error {
	var CalibrationcreateMap map[string]interface{}
	err := json.Unmarshal(b, &CalibrationcreateMap)
	if err != nil {
		return err
	}
	
	if Id, ok := CalibrationcreateMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := CalibrationcreateMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Calibrator, ok := CalibrationcreateMap["calibrator"].(map[string]interface{}); ok {
		CalibratorString, _ := json.Marshal(Calibrator)
		json.Unmarshal(CalibratorString, &o.Calibrator)
	}
	
	if Agent, ok := CalibrationcreateMap["agent"].(map[string]interface{}); ok {
		AgentString, _ := json.Marshal(Agent)
		json.Unmarshal(AgentString, &o.Agent)
	}
	
	if Conversation, ok := CalibrationcreateMap["conversation"].(map[string]interface{}); ok {
		ConversationString, _ := json.Marshal(Conversation)
		json.Unmarshal(ConversationString, &o.Conversation)
	}
	
	if EvaluationForm, ok := CalibrationcreateMap["evaluationForm"].(map[string]interface{}); ok {
		EvaluationFormString, _ := json.Marshal(EvaluationForm)
		json.Unmarshal(EvaluationFormString, &o.EvaluationForm)
	}
	
	if ContextId, ok := CalibrationcreateMap["contextId"].(string); ok {
		o.ContextId = &ContextId
	}
    
	if AverageScore, ok := CalibrationcreateMap["averageScore"].(float64); ok {
		AverageScoreInt := int(AverageScore)
		o.AverageScore = &AverageScoreInt
	}
	
	if HighScore, ok := CalibrationcreateMap["highScore"].(float64); ok {
		HighScoreInt := int(HighScore)
		o.HighScore = &HighScoreInt
	}
	
	if LowScore, ok := CalibrationcreateMap["lowScore"].(float64); ok {
		LowScoreInt := int(LowScore)
		o.LowScore = &LowScoreInt
	}
	
	if createdDateString, ok := CalibrationcreateMap["createdDate"].(string); ok {
		CreatedDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", createdDateString)
		o.CreatedDate = &CreatedDate
	}
	
	if Evaluations, ok := CalibrationcreateMap["evaluations"].([]interface{}); ok {
		EvaluationsString, _ := json.Marshal(Evaluations)
		json.Unmarshal(EvaluationsString, &o.Evaluations)
	}
	
	if Evaluators, ok := CalibrationcreateMap["evaluators"].([]interface{}); ok {
		EvaluatorsString, _ := json.Marshal(Evaluators)
		json.Unmarshal(EvaluatorsString, &o.Evaluators)
	}
	
	if ScoringIndex, ok := CalibrationcreateMap["scoringIndex"].(map[string]interface{}); ok {
		ScoringIndexString, _ := json.Marshal(ScoringIndex)
		json.Unmarshal(ScoringIndexString, &o.ScoringIndex)
	}
	
	if ExpertEvaluator, ok := CalibrationcreateMap["expertEvaluator"].(map[string]interface{}); ok {
		ExpertEvaluatorString, _ := json.Marshal(ExpertEvaluator)
		json.Unmarshal(ExpertEvaluatorString, &o.ExpertEvaluator)
	}
	
	if SelfUri, ok := CalibrationcreateMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Calibrationcreate) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
