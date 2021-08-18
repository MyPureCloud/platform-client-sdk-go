package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Evaluationqualityv2topicevaluationv2
type Evaluationqualityv2topicevaluationv2 struct { 
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

func (u *Evaluationqualityv2topicevaluationv2) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Evaluationqualityv2topicevaluationv2

	
	EventTime := new(string)
	if u.EventTime != nil {
		
		*EventTime = timeutil.Strftime(u.EventTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		EventTime = nil
	}
	
	ReleaseDate := new(string)
	if u.ReleaseDate != nil {
		
		*ReleaseDate = timeutil.Strftime(u.ReleaseDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ReleaseDate = nil
	}
	
	AssignedDate := new(string)
	if u.AssignedDate != nil {
		
		*AssignedDate = timeutil.Strftime(u.AssignedDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		AssignedDate = nil
	}
	
	ChangedDate := new(string)
	if u.ChangedDate != nil {
		
		*ChangedDate = timeutil.Strftime(u.ChangedDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ChangedDate = nil
	}
	
	ConversationDate := new(string)
	if u.ConversationDate != nil {
		
		*ConversationDate = timeutil.Strftime(u.ConversationDate, "%Y-%m-%dT%H:%M:%S.%fZ")
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
		*Alias
	}{ 
		Id: u.Id,
		
		ConversationId: u.ConversationId,
		
		Agent: u.Agent,
		
		Evaluator: u.Evaluator,
		
		EventTime: EventTime,
		
		EvaluationFormId: u.EvaluationFormId,
		
		FormName: u.FormName,
		
		ScoringSet: u.ScoringSet,
		
		ContextId: u.ContextId,
		
		Status: u.Status,
		
		AgentHasRead: u.AgentHasRead,
		
		ReleaseDate: ReleaseDate,
		
		AssignedDate: AssignedDate,
		
		ChangedDate: ChangedDate,
		
		EventType: u.EventType,
		
		ResourceId: u.ResourceId,
		
		ResourceType: u.ResourceType,
		
		DivisionIds: u.DivisionIds,
		
		Rescore: u.Rescore,
		
		ConversationDate: ConversationDate,
		
		MediaType: u.MediaType,
		
		Calibration: u.Calibration,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Evaluationqualityv2topicevaluationv2) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
