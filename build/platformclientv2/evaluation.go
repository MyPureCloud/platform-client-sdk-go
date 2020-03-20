package platformclientv2
import (
	"time"
	"encoding/json"
)

// Evaluation
type Evaluation struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Conversation
	Conversation *Conversation `json:"conversation,omitempty"`


	// EvaluationForm - Evaluation form used for evaluation.
	EvaluationForm *Evaluationform `json:"evaluationForm,omitempty"`


	// Evaluator
	Evaluator *User `json:"evaluator,omitempty"`


	// Agent
	Agent *User `json:"agent,omitempty"`


	// Calibration
	Calibration *Calibration `json:"calibration,omitempty"`


	// Status
	Status *string `json:"status,omitempty"`


	// Answers
	Answers *Evaluationscoringset `json:"answers,omitempty"`


	// AgentHasRead
	AgentHasRead *bool `json:"agentHasRead,omitempty"`


	// ReleaseDate - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	ReleaseDate *time.Time `json:"releaseDate,omitempty"`


	// AssignedDate - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	AssignedDate *time.Time `json:"assignedDate,omitempty"`


	// ChangedDate - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	ChangedDate *time.Time `json:"changedDate,omitempty"`


	// Queue
	Queue *Queue `json:"queue,omitempty"`


	// MediaType - List of different communication types used in conversation.
	MediaType *[]string `json:"mediaType,omitempty"`


	// Rescore - Is only true when evaluation is re-scored.
	Rescore *bool `json:"rescore,omitempty"`


	// ConversationDate - Date of conversation. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	ConversationDate *time.Time `json:"conversationDate,omitempty"`


	// ConversationEndDate - End date of conversation if it had completed before evaluation creation. Null if created before the conversation ended. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	ConversationEndDate *time.Time `json:"conversationEndDate,omitempty"`


	// NeverRelease - Signifies if the evaluation is never to be released. This cannot be set true if release date is also set.
	NeverRelease *bool `json:"neverRelease,omitempty"`


	// ResourceId - Only used for email evaluations. Will be null for all other evaluations.
	ResourceId *string `json:"resourceId,omitempty"`


	// ResourceType - The type of resource. Only used for email evaluations. Will be null for evaluations on all other resources.
	ResourceType *string `json:"resourceType,omitempty"`


	// Redacted - Is only true when the user making the request does not have sufficient permissions to see evaluation
	Redacted *bool `json:"redacted,omitempty"`


	// IsScoringIndex
	IsScoringIndex *bool `json:"isScoringIndex,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Evaluation) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
