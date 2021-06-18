package platformclientv2
import (
	"time"
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

// String returns a JSON representation of the model
func (o *Evaluationqualityv2topicevaluationv2) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
