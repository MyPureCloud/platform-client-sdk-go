package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
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


	// ReleaseDate - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	ReleaseDate *time.Time `json:"releaseDate,omitempty"`


	// AssignedDate - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	AssignedDate *time.Time `json:"assignedDate,omitempty"`


	// ChangedDate - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	ChangedDate *time.Time `json:"changedDate,omitempty"`


	// Queue
	Queue *Queue `json:"queue,omitempty"`


	// MediaType - List of different communication types used in conversation.
	MediaType *[]string `json:"mediaType,omitempty"`


	// Rescore - Is only true when evaluation is re-scored.
	Rescore *bool `json:"rescore,omitempty"`


	// ConversationDate - Date of conversation. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	ConversationDate *time.Time `json:"conversationDate,omitempty"`


	// ConversationEndDate - End date of conversation if it had completed before evaluation creation. Null if created before the conversation ended. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
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


	// AuthorizedActions - List of user authorized actions on evaluation. Possible values: edit, editScore, editAgentSignoff, delete, viewAudit
	AuthorizedActions *[]string `json:"authorizedActions,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (u *Evaluation) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Evaluation

	
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
	
	ConversationEndDate := new(string)
	if u.ConversationEndDate != nil {
		
		*ConversationEndDate = timeutil.Strftime(u.ConversationEndDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ConversationEndDate = nil
	}
	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Conversation *Conversation `json:"conversation,omitempty"`
		
		EvaluationForm *Evaluationform `json:"evaluationForm,omitempty"`
		
		Evaluator *User `json:"evaluator,omitempty"`
		
		Agent *User `json:"agent,omitempty"`
		
		Calibration *Calibration `json:"calibration,omitempty"`
		
		Status *string `json:"status,omitempty"`
		
		Answers *Evaluationscoringset `json:"answers,omitempty"`
		
		AgentHasRead *bool `json:"agentHasRead,omitempty"`
		
		ReleaseDate *string `json:"releaseDate,omitempty"`
		
		AssignedDate *string `json:"assignedDate,omitempty"`
		
		ChangedDate *string `json:"changedDate,omitempty"`
		
		Queue *Queue `json:"queue,omitempty"`
		
		MediaType *[]string `json:"mediaType,omitempty"`
		
		Rescore *bool `json:"rescore,omitempty"`
		
		ConversationDate *string `json:"conversationDate,omitempty"`
		
		ConversationEndDate *string `json:"conversationEndDate,omitempty"`
		
		NeverRelease *bool `json:"neverRelease,omitempty"`
		
		ResourceId *string `json:"resourceId,omitempty"`
		
		ResourceType *string `json:"resourceType,omitempty"`
		
		Redacted *bool `json:"redacted,omitempty"`
		
		IsScoringIndex *bool `json:"isScoringIndex,omitempty"`
		
		AuthorizedActions *[]string `json:"authorizedActions,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Name: u.Name,
		
		Conversation: u.Conversation,
		
		EvaluationForm: u.EvaluationForm,
		
		Evaluator: u.Evaluator,
		
		Agent: u.Agent,
		
		Calibration: u.Calibration,
		
		Status: u.Status,
		
		Answers: u.Answers,
		
		AgentHasRead: u.AgentHasRead,
		
		ReleaseDate: ReleaseDate,
		
		AssignedDate: AssignedDate,
		
		ChangedDate: ChangedDate,
		
		Queue: u.Queue,
		
		MediaType: u.MediaType,
		
		Rescore: u.Rescore,
		
		ConversationDate: ConversationDate,
		
		ConversationEndDate: ConversationEndDate,
		
		NeverRelease: u.NeverRelease,
		
		ResourceId: u.ResourceId,
		
		ResourceType: u.ResourceType,
		
		Redacted: u.Redacted,
		
		IsScoringIndex: u.IsScoringIndex,
		
		AuthorizedActions: u.AuthorizedActions,
		
		SelfUri: u.SelfUri,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Evaluation) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
