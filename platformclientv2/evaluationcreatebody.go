package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Evaluationcreatebody
type Evaluationcreatebody struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// EvaluationForm - Evaluation form used for evaluation (must be included for a successful request)
	EvaluationForm *Evaluationcreateevalform `json:"evaluationForm,omitempty"`

	// Evaluator - User ID of the evaluator (must be included for a successful request)
	Evaluator *Evaluationcreateuser `json:"evaluator,omitempty"`

	// Agent - User ID of the agent (must be included for a successful request)
	Agent *Evaluationcreateuser `json:"agent,omitempty"`

	// AgentHasRead
	AgentHasRead *bool `json:"agentHasRead,omitempty"`

	// Answers
	Answers *Evaluationscoringset `json:"answers,omitempty"`

	// Calibration
	Calibration *Evaluationcreatecalibration `json:"calibration,omitempty"`

	// EvaluationContextId
	EvaluationContextId *string `json:"evaluationContextId,omitempty"`

	// Conversation
	Conversation *Evaluationcreateconversation `json:"conversation,omitempty"`

	// ResourceType
	ResourceType *string `json:"resourceType,omitempty"`

	// EvaluationSource
	EvaluationSource *Evaluationsource `json:"evaluationSource,omitempty"`

	// Rescore
	Rescore *bool `json:"rescore,omitempty"`

	// Queue
	Queue *Evaluationcreatequeue `json:"queue,omitempty"`

	// ReleaseDate - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	ReleaseDate *time.Time `json:"releaseDate,omitempty"`

	// Status
	Status *string `json:"status,omitempty"`

	// NeverRelease
	NeverRelease *bool `json:"neverRelease,omitempty"`

	// DateAssigneeChanged - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateAssigneeChanged *time.Time `json:"dateAssigneeChanged,omitempty"`

	// Assignee
	Assignee *Evaluationcreateuser `json:"assignee,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Evaluationcreatebody) SetField(field string, fieldValue interface{}) {
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

func (o Evaluationcreatebody) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "ReleaseDate","DateAssigneeChanged", }
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
	type Alias Evaluationcreatebody
	
	ReleaseDate := new(string)
	if o.ReleaseDate != nil {
		
		*ReleaseDate = timeutil.Strftime(o.ReleaseDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ReleaseDate = nil
	}
	
	DateAssigneeChanged := new(string)
	if o.DateAssigneeChanged != nil {
		
		*DateAssigneeChanged = timeutil.Strftime(o.DateAssigneeChanged, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateAssigneeChanged = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		EvaluationForm *Evaluationcreateevalform `json:"evaluationForm,omitempty"`
		
		Evaluator *Evaluationcreateuser `json:"evaluator,omitempty"`
		
		Agent *Evaluationcreateuser `json:"agent,omitempty"`
		
		AgentHasRead *bool `json:"agentHasRead,omitempty"`
		
		Answers *Evaluationscoringset `json:"answers,omitempty"`
		
		Calibration *Evaluationcreatecalibration `json:"calibration,omitempty"`
		
		EvaluationContextId *string `json:"evaluationContextId,omitempty"`
		
		Conversation *Evaluationcreateconversation `json:"conversation,omitempty"`
		
		ResourceType *string `json:"resourceType,omitempty"`
		
		EvaluationSource *Evaluationsource `json:"evaluationSource,omitempty"`
		
		Rescore *bool `json:"rescore,omitempty"`
		
		Queue *Evaluationcreatequeue `json:"queue,omitempty"`
		
		ReleaseDate *string `json:"releaseDate,omitempty"`
		
		Status *string `json:"status,omitempty"`
		
		NeverRelease *bool `json:"neverRelease,omitempty"`
		
		DateAssigneeChanged *string `json:"dateAssigneeChanged,omitempty"`
		
		Assignee *Evaluationcreateuser `json:"assignee,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		EvaluationForm: o.EvaluationForm,
		
		Evaluator: o.Evaluator,
		
		Agent: o.Agent,
		
		AgentHasRead: o.AgentHasRead,
		
		Answers: o.Answers,
		
		Calibration: o.Calibration,
		
		EvaluationContextId: o.EvaluationContextId,
		
		Conversation: o.Conversation,
		
		ResourceType: o.ResourceType,
		
		EvaluationSource: o.EvaluationSource,
		
		Rescore: o.Rescore,
		
		Queue: o.Queue,
		
		ReleaseDate: ReleaseDate,
		
		Status: o.Status,
		
		NeverRelease: o.NeverRelease,
		
		DateAssigneeChanged: DateAssigneeChanged,
		
		Assignee: o.Assignee,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Evaluationcreatebody) UnmarshalJSON(b []byte) error {
	var EvaluationcreatebodyMap map[string]interface{}
	err := json.Unmarshal(b, &EvaluationcreatebodyMap)
	if err != nil {
		return err
	}
	
	if Id, ok := EvaluationcreatebodyMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if EvaluationForm, ok := EvaluationcreatebodyMap["evaluationForm"].(map[string]interface{}); ok {
		EvaluationFormString, _ := json.Marshal(EvaluationForm)
		json.Unmarshal(EvaluationFormString, &o.EvaluationForm)
	}
	
	if Evaluator, ok := EvaluationcreatebodyMap["evaluator"].(map[string]interface{}); ok {
		EvaluatorString, _ := json.Marshal(Evaluator)
		json.Unmarshal(EvaluatorString, &o.Evaluator)
	}
	
	if Agent, ok := EvaluationcreatebodyMap["agent"].(map[string]interface{}); ok {
		AgentString, _ := json.Marshal(Agent)
		json.Unmarshal(AgentString, &o.Agent)
	}
	
	if AgentHasRead, ok := EvaluationcreatebodyMap["agentHasRead"].(bool); ok {
		o.AgentHasRead = &AgentHasRead
	}
    
	if Answers, ok := EvaluationcreatebodyMap["answers"].(map[string]interface{}); ok {
		AnswersString, _ := json.Marshal(Answers)
		json.Unmarshal(AnswersString, &o.Answers)
	}
	
	if Calibration, ok := EvaluationcreatebodyMap["calibration"].(map[string]interface{}); ok {
		CalibrationString, _ := json.Marshal(Calibration)
		json.Unmarshal(CalibrationString, &o.Calibration)
	}
	
	if EvaluationContextId, ok := EvaluationcreatebodyMap["evaluationContextId"].(string); ok {
		o.EvaluationContextId = &EvaluationContextId
	}
    
	if Conversation, ok := EvaluationcreatebodyMap["conversation"].(map[string]interface{}); ok {
		ConversationString, _ := json.Marshal(Conversation)
		json.Unmarshal(ConversationString, &o.Conversation)
	}
	
	if ResourceType, ok := EvaluationcreatebodyMap["resourceType"].(string); ok {
		o.ResourceType = &ResourceType
	}
    
	if EvaluationSource, ok := EvaluationcreatebodyMap["evaluationSource"].(map[string]interface{}); ok {
		EvaluationSourceString, _ := json.Marshal(EvaluationSource)
		json.Unmarshal(EvaluationSourceString, &o.EvaluationSource)
	}
	
	if Rescore, ok := EvaluationcreatebodyMap["rescore"].(bool); ok {
		o.Rescore = &Rescore
	}
    
	if Queue, ok := EvaluationcreatebodyMap["queue"].(map[string]interface{}); ok {
		QueueString, _ := json.Marshal(Queue)
		json.Unmarshal(QueueString, &o.Queue)
	}
	
	if releaseDateString, ok := EvaluationcreatebodyMap["releaseDate"].(string); ok {
		ReleaseDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", releaseDateString)
		o.ReleaseDate = &ReleaseDate
	}
	
	if Status, ok := EvaluationcreatebodyMap["status"].(string); ok {
		o.Status = &Status
	}
    
	if NeverRelease, ok := EvaluationcreatebodyMap["neverRelease"].(bool); ok {
		o.NeverRelease = &NeverRelease
	}
    
	if dateAssigneeChangedString, ok := EvaluationcreatebodyMap["dateAssigneeChanged"].(string); ok {
		DateAssigneeChanged, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateAssigneeChangedString)
		o.DateAssigneeChanged = &DateAssigneeChanged
	}
	
	if Assignee, ok := EvaluationcreatebodyMap["assignee"].(map[string]interface{}); ok {
		AssigneeString, _ := json.Marshal(Assignee)
		json.Unmarshal(AssigneeString, &o.Assignee)
	}
	
	if SelfUri, ok := EvaluationcreatebodyMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Evaluationcreatebody) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
