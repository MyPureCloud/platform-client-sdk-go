package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Evaluationqualityv2topicevaluationv2
type Evaluationqualityv2topicevaluationv2 struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id
	Id *string `json:"id,omitempty"`

	// ConversationId
	ConversationId *string `json:"conversationId,omitempty"`

	// Agent
	Agent *Evaluationqualityv2topicuser `json:"agent,omitempty"`

	// Evaluator
	Evaluator *Evaluationqualityv2topicuser `json:"evaluator,omitempty"`

	// EventTime
	EventTime *time.Time `json:"eventTime,omitempty"`

	// EvaluationFormId
	EvaluationFormId *string `json:"evaluationFormId,omitempty"`

	// FormName
	FormName *string `json:"formName,omitempty"`

	// ScoringSet
	ScoringSet *Evaluationqualityv2topicevaluationscoringset `json:"scoringSet,omitempty"`

	// ContextId
	ContextId *string `json:"contextId,omitempty"`

	// Status
	Status *string `json:"status,omitempty"`

	// AgentHasRead
	AgentHasRead *bool `json:"agentHasRead,omitempty"`

	// ReleaseDate
	ReleaseDate *time.Time `json:"releaseDate,omitempty"`

	// AssignedDate
	AssignedDate *time.Time `json:"assignedDate,omitempty"`

	// ChangedDate
	ChangedDate *time.Time `json:"changedDate,omitempty"`

	// EventType
	EventType *string `json:"eventType,omitempty"`

	// ResourceId
	ResourceId *string `json:"resourceId,omitempty"`

	// ResourceType
	ResourceType *string `json:"resourceType,omitempty"`

	// DivisionIds
	DivisionIds *[]string `json:"divisionIds,omitempty"`

	// Rescore
	Rescore *bool `json:"rescore,omitempty"`

	// ConversationDate
	ConversationDate *time.Time `json:"conversationDate,omitempty"`

	// MediaType
	MediaType *[]string `json:"mediaType,omitempty"`

	// Calibration
	Calibration *Evaluationqualityv2topiccalibration `json:"calibration,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Evaluationqualityv2topicevaluationv2) SetField(field string, fieldValue interface{}) {
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

func (o Evaluationqualityv2topicevaluationv2) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "EventTime","ReleaseDate","AssignedDate","ChangedDate","ConversationDate", }
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
	type Alias Evaluationqualityv2topicevaluationv2
	
	EventTime := new(string)
	if o.EventTime != nil {
		
		*EventTime = timeutil.Strftime(o.EventTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		EventTime = nil
	}
	
	ReleaseDate := new(string)
	if o.ReleaseDate != nil {
		
		*ReleaseDate = timeutil.Strftime(o.ReleaseDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ReleaseDate = nil
	}
	
	AssignedDate := new(string)
	if o.AssignedDate != nil {
		
		*AssignedDate = timeutil.Strftime(o.AssignedDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		AssignedDate = nil
	}
	
	ChangedDate := new(string)
	if o.ChangedDate != nil {
		
		*ChangedDate = timeutil.Strftime(o.ChangedDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ChangedDate = nil
	}
	
	ConversationDate := new(string)
	if o.ConversationDate != nil {
		
		*ConversationDate = timeutil.Strftime(o.ConversationDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ConversationDate = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		ConversationId *string `json:"conversationId,omitempty"`
		
		Agent *Evaluationqualityv2topicuser `json:"agent,omitempty"`
		
		Evaluator *Evaluationqualityv2topicuser `json:"evaluator,omitempty"`
		
		EventTime *string `json:"eventTime,omitempty"`
		
		EvaluationFormId *string `json:"evaluationFormId,omitempty"`
		
		FormName *string `json:"formName,omitempty"`
		
		ScoringSet *Evaluationqualityv2topicevaluationscoringset `json:"scoringSet,omitempty"`
		
		ContextId *string `json:"contextId,omitempty"`
		
		Status *string `json:"status,omitempty"`
		
		AgentHasRead *bool `json:"agentHasRead,omitempty"`
		
		ReleaseDate *string `json:"releaseDate,omitempty"`
		
		AssignedDate *string `json:"assignedDate,omitempty"`
		
		ChangedDate *string `json:"changedDate,omitempty"`
		
		EventType *string `json:"eventType,omitempty"`
		
		ResourceId *string `json:"resourceId,omitempty"`
		
		ResourceType *string `json:"resourceType,omitempty"`
		
		DivisionIds *[]string `json:"divisionIds,omitempty"`
		
		Rescore *bool `json:"rescore,omitempty"`
		
		ConversationDate *string `json:"conversationDate,omitempty"`
		
		MediaType *[]string `json:"mediaType,omitempty"`
		
		Calibration *Evaluationqualityv2topiccalibration `json:"calibration,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		ConversationId: o.ConversationId,
		
		Agent: o.Agent,
		
		Evaluator: o.Evaluator,
		
		EventTime: EventTime,
		
		EvaluationFormId: o.EvaluationFormId,
		
		FormName: o.FormName,
		
		ScoringSet: o.ScoringSet,
		
		ContextId: o.ContextId,
		
		Status: o.Status,
		
		AgentHasRead: o.AgentHasRead,
		
		ReleaseDate: ReleaseDate,
		
		AssignedDate: AssignedDate,
		
		ChangedDate: ChangedDate,
		
		EventType: o.EventType,
		
		ResourceId: o.ResourceId,
		
		ResourceType: o.ResourceType,
		
		DivisionIds: o.DivisionIds,
		
		Rescore: o.Rescore,
		
		ConversationDate: ConversationDate,
		
		MediaType: o.MediaType,
		
		Calibration: o.Calibration,
		Alias:    (Alias)(o),
	})
}

func (o *Evaluationqualityv2topicevaluationv2) UnmarshalJSON(b []byte) error {
	var Evaluationqualityv2topicevaluationv2Map map[string]interface{}
	err := json.Unmarshal(b, &Evaluationqualityv2topicevaluationv2Map)
	if err != nil {
		return err
	}
	
	if Id, ok := Evaluationqualityv2topicevaluationv2Map["id"].(string); ok {
		o.Id = &Id
	}
    
	if ConversationId, ok := Evaluationqualityv2topicevaluationv2Map["conversationId"].(string); ok {
		o.ConversationId = &ConversationId
	}
    
	if Agent, ok := Evaluationqualityv2topicevaluationv2Map["agent"].(map[string]interface{}); ok {
		AgentString, _ := json.Marshal(Agent)
		json.Unmarshal(AgentString, &o.Agent)
	}
	
	if Evaluator, ok := Evaluationqualityv2topicevaluationv2Map["evaluator"].(map[string]interface{}); ok {
		EvaluatorString, _ := json.Marshal(Evaluator)
		json.Unmarshal(EvaluatorString, &o.Evaluator)
	}
	
	if eventTimeString, ok := Evaluationqualityv2topicevaluationv2Map["eventTime"].(string); ok {
		EventTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", eventTimeString)
		o.EventTime = &EventTime
	}
	
	if EvaluationFormId, ok := Evaluationqualityv2topicevaluationv2Map["evaluationFormId"].(string); ok {
		o.EvaluationFormId = &EvaluationFormId
	}
    
	if FormName, ok := Evaluationqualityv2topicevaluationv2Map["formName"].(string); ok {
		o.FormName = &FormName
	}
    
	if ScoringSet, ok := Evaluationqualityv2topicevaluationv2Map["scoringSet"].(map[string]interface{}); ok {
		ScoringSetString, _ := json.Marshal(ScoringSet)
		json.Unmarshal(ScoringSetString, &o.ScoringSet)
	}
	
	if ContextId, ok := Evaluationqualityv2topicevaluationv2Map["contextId"].(string); ok {
		o.ContextId = &ContextId
	}
    
	if Status, ok := Evaluationqualityv2topicevaluationv2Map["status"].(string); ok {
		o.Status = &Status
	}
    
	if AgentHasRead, ok := Evaluationqualityv2topicevaluationv2Map["agentHasRead"].(bool); ok {
		o.AgentHasRead = &AgentHasRead
	}
    
	if releaseDateString, ok := Evaluationqualityv2topicevaluationv2Map["releaseDate"].(string); ok {
		ReleaseDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", releaseDateString)
		o.ReleaseDate = &ReleaseDate
	}
	
	if assignedDateString, ok := Evaluationqualityv2topicevaluationv2Map["assignedDate"].(string); ok {
		AssignedDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", assignedDateString)
		o.AssignedDate = &AssignedDate
	}
	
	if changedDateString, ok := Evaluationqualityv2topicevaluationv2Map["changedDate"].(string); ok {
		ChangedDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", changedDateString)
		o.ChangedDate = &ChangedDate
	}
	
	if EventType, ok := Evaluationqualityv2topicevaluationv2Map["eventType"].(string); ok {
		o.EventType = &EventType
	}
    
	if ResourceId, ok := Evaluationqualityv2topicevaluationv2Map["resourceId"].(string); ok {
		o.ResourceId = &ResourceId
	}
    
	if ResourceType, ok := Evaluationqualityv2topicevaluationv2Map["resourceType"].(string); ok {
		o.ResourceType = &ResourceType
	}
    
	if DivisionIds, ok := Evaluationqualityv2topicevaluationv2Map["divisionIds"].([]interface{}); ok {
		DivisionIdsString, _ := json.Marshal(DivisionIds)
		json.Unmarshal(DivisionIdsString, &o.DivisionIds)
	}
	
	if Rescore, ok := Evaluationqualityv2topicevaluationv2Map["rescore"].(bool); ok {
		o.Rescore = &Rescore
	}
    
	if conversationDateString, ok := Evaluationqualityv2topicevaluationv2Map["conversationDate"].(string); ok {
		ConversationDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", conversationDateString)
		o.ConversationDate = &ConversationDate
	}
	
	if MediaType, ok := Evaluationqualityv2topicevaluationv2Map["mediaType"].([]interface{}); ok {
		MediaTypeString, _ := json.Marshal(MediaType)
		json.Unmarshal(MediaTypeString, &o.MediaType)
	}
	
	if Calibration, ok := Evaluationqualityv2topicevaluationv2Map["calibration"].(map[string]interface{}); ok {
		CalibrationString, _ := json.Marshal(Calibration)
		json.Unmarshal(CalibrationString, &o.Calibration)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Evaluationqualityv2topicevaluationv2) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
