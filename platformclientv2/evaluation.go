package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Evaluation
type Evaluation struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// Name
	Name *string `json:"name,omitempty"`

	// Conversation
	Conversation *Conversationreference `json:"conversation,omitempty"`

	// EvaluationForm - Evaluation form used for evaluation.
	EvaluationForm *Evaluationform `json:"evaluationForm,omitempty"`

	// Evaluator
	Evaluator *User `json:"evaluator,omitempty"`

	// Agent
	Agent *User `json:"agent,omitempty"`

	// Calibration
	Calibration **Calibration `json:"calibration,omitempty"`

	// Status
	Status *string `json:"status,omitempty"`

	// Answers
	Answers *Evaluationscoringset `json:"answers,omitempty"`

	// AgentHasRead
	AgentHasRead *bool `json:"agentHasRead,omitempty"`

	// Assignee
	Assignee *User `json:"assignee,omitempty"`

	// AssigneeApplicable - Indicates whether an assignee is applicable for the evaluation. Set to false when assignee is not applicable.
	AssigneeApplicable *bool `json:"assigneeApplicable,omitempty"`

	// ReleaseDate - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	ReleaseDate *time.Time `json:"releaseDate,omitempty"`

	// AssignedDate - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	AssignedDate *time.Time `json:"assignedDate,omitempty"`

	// CreatedDate - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	CreatedDate *time.Time `json:"createdDate,omitempty"`

	// ChangedDate - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	ChangedDate *time.Time `json:"changedDate,omitempty"`

	// SubmittedDate - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	SubmittedDate *time.Time `json:"submittedDate,omitempty"`

	// RevisionCreatedDate - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	RevisionCreatedDate *time.Time `json:"revisionCreatedDate,omitempty"`

	// Queue
	Queue *Queue `json:"queue,omitempty"`

	// MediaType - List of different communication types used in conversation.
	MediaType *[]string `json:"mediaType,omitempty"`

	// Rescore - Is only true when evaluation is re-scored.
	Rescore *bool `json:"rescore,omitempty"`

	// ConversationDate - Creation date of the conversation. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	ConversationDate *time.Time `json:"conversationDate,omitempty"`

	// ConversationEndDate - End date of conversation if it had completed before evaluation creation. Null if created before the conversation ended. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	ConversationEndDate *time.Time `json:"conversationEndDate,omitempty"`

	// NeverRelease - Signifies if the evaluation is never to be released. This cannot be set true if release date is also set.
	NeverRelease *bool `json:"neverRelease,omitempty"`

	// Assigned - Set to false to unassign the evaluation. This cannot be set to false when assignee is also set.
	Assigned *bool `json:"assigned,omitempty"`

	// DateAssigneeChanged - Date when the assignee was last changed. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateAssigneeChanged *time.Time `json:"dateAssigneeChanged,omitempty"`

	// ResourceId - Only used for email evaluations. Will be null for all other evaluations.
	ResourceId *string `json:"resourceId,omitempty"`

	// ResourceType - The type of resource. Only used for email evaluations. Will be null for evaluations on all other resources.
	ResourceType *string `json:"resourceType,omitempty"`

	// Redacted - Is only true when the user making the request does not have sufficient permissions to see evaluation
	Redacted *bool `json:"redacted,omitempty"`

	// AgentTeam - Team of the evaluation agent
	AgentTeam *Team `json:"agentTeam,omitempty"`

	// IsScoringIndex
	IsScoringIndex *bool `json:"isScoringIndex,omitempty"`

	// AuthorizedActions - List of user authorized actions on evaluation. Possible values: assign, edit, editScore, editAgentSignoff, delete, release, viewAudit
	AuthorizedActions *[]string `json:"authorizedActions,omitempty"`

	// HasAssistanceFailed - Is true when evaluation assistance didn't execute successfully
	HasAssistanceFailed *bool `json:"hasAssistanceFailed,omitempty"`

	// EvaluationSource - The source that created the evaluation.
	EvaluationSource *Evaluationsource `json:"evaluationSource,omitempty"`

	// AiScoring - AI scoring details for the evaluation.
	AiScoring *Aiscoring `json:"aiScoring,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Evaluation) SetField(field string, fieldValue interface{}) {
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

func (o Evaluation) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "ReleaseDate","AssignedDate","CreatedDate","ChangedDate","SubmittedDate","RevisionCreatedDate","ConversationDate","ConversationEndDate","DateAssigneeChanged", }
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
	type Alias Evaluation
	
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
	
	CreatedDate := new(string)
	if o.CreatedDate != nil {
		
		*CreatedDate = timeutil.Strftime(o.CreatedDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		CreatedDate = nil
	}
	
	ChangedDate := new(string)
	if o.ChangedDate != nil {
		
		*ChangedDate = timeutil.Strftime(o.ChangedDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ChangedDate = nil
	}
	
	SubmittedDate := new(string)
	if o.SubmittedDate != nil {
		
		*SubmittedDate = timeutil.Strftime(o.SubmittedDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		SubmittedDate = nil
	}
	
	RevisionCreatedDate := new(string)
	if o.RevisionCreatedDate != nil {
		
		*RevisionCreatedDate = timeutil.Strftime(o.RevisionCreatedDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		RevisionCreatedDate = nil
	}
	
	ConversationDate := new(string)
	if o.ConversationDate != nil {
		
		*ConversationDate = timeutil.Strftime(o.ConversationDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ConversationDate = nil
	}
	
	ConversationEndDate := new(string)
	if o.ConversationEndDate != nil {
		
		*ConversationEndDate = timeutil.Strftime(o.ConversationEndDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ConversationEndDate = nil
	}
	
	DateAssigneeChanged := new(string)
	if o.DateAssigneeChanged != nil {
		
		*DateAssigneeChanged = timeutil.Strftime(o.DateAssigneeChanged, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateAssigneeChanged = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Conversation *Conversationreference `json:"conversation,omitempty"`
		
		EvaluationForm *Evaluationform `json:"evaluationForm,omitempty"`
		
		Evaluator *User `json:"evaluator,omitempty"`
		
		Agent *User `json:"agent,omitempty"`
		
		Calibration **Calibration `json:"calibration,omitempty"`
		
		Status *string `json:"status,omitempty"`
		
		Answers *Evaluationscoringset `json:"answers,omitempty"`
		
		AgentHasRead *bool `json:"agentHasRead,omitempty"`
		
		Assignee *User `json:"assignee,omitempty"`
		
		AssigneeApplicable *bool `json:"assigneeApplicable,omitempty"`
		
		ReleaseDate *string `json:"releaseDate,omitempty"`
		
		AssignedDate *string `json:"assignedDate,omitempty"`
		
		CreatedDate *string `json:"createdDate,omitempty"`
		
		ChangedDate *string `json:"changedDate,omitempty"`
		
		SubmittedDate *string `json:"submittedDate,omitempty"`
		
		RevisionCreatedDate *string `json:"revisionCreatedDate,omitempty"`
		
		Queue *Queue `json:"queue,omitempty"`
		
		MediaType *[]string `json:"mediaType,omitempty"`
		
		Rescore *bool `json:"rescore,omitempty"`
		
		ConversationDate *string `json:"conversationDate,omitempty"`
		
		ConversationEndDate *string `json:"conversationEndDate,omitempty"`
		
		NeverRelease *bool `json:"neverRelease,omitempty"`
		
		Assigned *bool `json:"assigned,omitempty"`
		
		DateAssigneeChanged *string `json:"dateAssigneeChanged,omitempty"`
		
		ResourceId *string `json:"resourceId,omitempty"`
		
		ResourceType *string `json:"resourceType,omitempty"`
		
		Redacted *bool `json:"redacted,omitempty"`
		
		AgentTeam *Team `json:"agentTeam,omitempty"`
		
		IsScoringIndex *bool `json:"isScoringIndex,omitempty"`
		
		AuthorizedActions *[]string `json:"authorizedActions,omitempty"`
		
		HasAssistanceFailed *bool `json:"hasAssistanceFailed,omitempty"`
		
		EvaluationSource *Evaluationsource `json:"evaluationSource,omitempty"`
		
		AiScoring *Aiscoring `json:"aiScoring,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Conversation: o.Conversation,
		
		EvaluationForm: o.EvaluationForm,
		
		Evaluator: o.Evaluator,
		
		Agent: o.Agent,
		
		Calibration: o.Calibration,
		
		Status: o.Status,
		
		Answers: o.Answers,
		
		AgentHasRead: o.AgentHasRead,
		
		Assignee: o.Assignee,
		
		AssigneeApplicable: o.AssigneeApplicable,
		
		ReleaseDate: ReleaseDate,
		
		AssignedDate: AssignedDate,
		
		CreatedDate: CreatedDate,
		
		ChangedDate: ChangedDate,
		
		SubmittedDate: SubmittedDate,
		
		RevisionCreatedDate: RevisionCreatedDate,
		
		Queue: o.Queue,
		
		MediaType: o.MediaType,
		
		Rescore: o.Rescore,
		
		ConversationDate: ConversationDate,
		
		ConversationEndDate: ConversationEndDate,
		
		NeverRelease: o.NeverRelease,
		
		Assigned: o.Assigned,
		
		DateAssigneeChanged: DateAssigneeChanged,
		
		ResourceId: o.ResourceId,
		
		ResourceType: o.ResourceType,
		
		Redacted: o.Redacted,
		
		AgentTeam: o.AgentTeam,
		
		IsScoringIndex: o.IsScoringIndex,
		
		AuthorizedActions: o.AuthorizedActions,
		
		HasAssistanceFailed: o.HasAssistanceFailed,
		
		EvaluationSource: o.EvaluationSource,
		
		AiScoring: o.AiScoring,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Evaluation) UnmarshalJSON(b []byte) error {
	var EvaluationMap map[string]interface{}
	err := json.Unmarshal(b, &EvaluationMap)
	if err != nil {
		return err
	}
	
	if Id, ok := EvaluationMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := EvaluationMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Conversation, ok := EvaluationMap["conversation"].(map[string]interface{}); ok {
		ConversationString, _ := json.Marshal(Conversation)
		json.Unmarshal(ConversationString, &o.Conversation)
	}
	
	if EvaluationForm, ok := EvaluationMap["evaluationForm"].(map[string]interface{}); ok {
		EvaluationFormString, _ := json.Marshal(EvaluationForm)
		json.Unmarshal(EvaluationFormString, &o.EvaluationForm)
	}
	
	if Evaluator, ok := EvaluationMap["evaluator"].(map[string]interface{}); ok {
		EvaluatorString, _ := json.Marshal(Evaluator)
		json.Unmarshal(EvaluatorString, &o.Evaluator)
	}
	
	if Agent, ok := EvaluationMap["agent"].(map[string]interface{}); ok {
		AgentString, _ := json.Marshal(Agent)
		json.Unmarshal(AgentString, &o.Agent)
	}
	
	if Calibration, ok := EvaluationMap["calibration"].(map[string]interface{}); ok {
		CalibrationString, _ := json.Marshal(Calibration)
		json.Unmarshal(CalibrationString, &o.Calibration)
	}
	
	if Status, ok := EvaluationMap["status"].(string); ok {
		o.Status = &Status
	}
    
	if Answers, ok := EvaluationMap["answers"].(map[string]interface{}); ok {
		AnswersString, _ := json.Marshal(Answers)
		json.Unmarshal(AnswersString, &o.Answers)
	}
	
	if AgentHasRead, ok := EvaluationMap["agentHasRead"].(bool); ok {
		o.AgentHasRead = &AgentHasRead
	}
    
	if Assignee, ok := EvaluationMap["assignee"].(map[string]interface{}); ok {
		AssigneeString, _ := json.Marshal(Assignee)
		json.Unmarshal(AssigneeString, &o.Assignee)
	}
	
	if AssigneeApplicable, ok := EvaluationMap["assigneeApplicable"].(bool); ok {
		o.AssigneeApplicable = &AssigneeApplicable
	}
    
	if releaseDateString, ok := EvaluationMap["releaseDate"].(string); ok {
		ReleaseDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", releaseDateString)
		o.ReleaseDate = &ReleaseDate
	}
	
	if assignedDateString, ok := EvaluationMap["assignedDate"].(string); ok {
		AssignedDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", assignedDateString)
		o.AssignedDate = &AssignedDate
	}
	
	if createdDateString, ok := EvaluationMap["createdDate"].(string); ok {
		CreatedDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", createdDateString)
		o.CreatedDate = &CreatedDate
	}
	
	if changedDateString, ok := EvaluationMap["changedDate"].(string); ok {
		ChangedDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", changedDateString)
		o.ChangedDate = &ChangedDate
	}
	
	if submittedDateString, ok := EvaluationMap["submittedDate"].(string); ok {
		SubmittedDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", submittedDateString)
		o.SubmittedDate = &SubmittedDate
	}
	
	if revisionCreatedDateString, ok := EvaluationMap["revisionCreatedDate"].(string); ok {
		RevisionCreatedDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", revisionCreatedDateString)
		o.RevisionCreatedDate = &RevisionCreatedDate
	}
	
	if Queue, ok := EvaluationMap["queue"].(map[string]interface{}); ok {
		QueueString, _ := json.Marshal(Queue)
		json.Unmarshal(QueueString, &o.Queue)
	}
	
	if MediaType, ok := EvaluationMap["mediaType"].([]interface{}); ok {
		MediaTypeString, _ := json.Marshal(MediaType)
		json.Unmarshal(MediaTypeString, &o.MediaType)
	}
	
	if Rescore, ok := EvaluationMap["rescore"].(bool); ok {
		o.Rescore = &Rescore
	}
    
	if conversationDateString, ok := EvaluationMap["conversationDate"].(string); ok {
		ConversationDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", conversationDateString)
		o.ConversationDate = &ConversationDate
	}
	
	if conversationEndDateString, ok := EvaluationMap["conversationEndDate"].(string); ok {
		ConversationEndDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", conversationEndDateString)
		o.ConversationEndDate = &ConversationEndDate
	}
	
	if NeverRelease, ok := EvaluationMap["neverRelease"].(bool); ok {
		o.NeverRelease = &NeverRelease
	}
    
	if Assigned, ok := EvaluationMap["assigned"].(bool); ok {
		o.Assigned = &Assigned
	}
    
	if dateAssigneeChangedString, ok := EvaluationMap["dateAssigneeChanged"].(string); ok {
		DateAssigneeChanged, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateAssigneeChangedString)
		o.DateAssigneeChanged = &DateAssigneeChanged
	}
	
	if ResourceId, ok := EvaluationMap["resourceId"].(string); ok {
		o.ResourceId = &ResourceId
	}
    
	if ResourceType, ok := EvaluationMap["resourceType"].(string); ok {
		o.ResourceType = &ResourceType
	}
    
	if Redacted, ok := EvaluationMap["redacted"].(bool); ok {
		o.Redacted = &Redacted
	}
    
	if AgentTeam, ok := EvaluationMap["agentTeam"].(map[string]interface{}); ok {
		AgentTeamString, _ := json.Marshal(AgentTeam)
		json.Unmarshal(AgentTeamString, &o.AgentTeam)
	}
	
	if IsScoringIndex, ok := EvaluationMap["isScoringIndex"].(bool); ok {
		o.IsScoringIndex = &IsScoringIndex
	}
    
	if AuthorizedActions, ok := EvaluationMap["authorizedActions"].([]interface{}); ok {
		AuthorizedActionsString, _ := json.Marshal(AuthorizedActions)
		json.Unmarshal(AuthorizedActionsString, &o.AuthorizedActions)
	}
	
	if HasAssistanceFailed, ok := EvaluationMap["hasAssistanceFailed"].(bool); ok {
		o.HasAssistanceFailed = &HasAssistanceFailed
	}
    
	if EvaluationSource, ok := EvaluationMap["evaluationSource"].(map[string]interface{}); ok {
		EvaluationSourceString, _ := json.Marshal(EvaluationSource)
		json.Unmarshal(EvaluationSourceString, &o.EvaluationSource)
	}
	
	if AiScoring, ok := EvaluationMap["aiScoring"].(map[string]interface{}); ok {
		AiScoringString, _ := json.Marshal(AiScoring)
		json.Unmarshal(AiScoringString, &o.AiScoring)
	}
	
	if SelfUri, ok := EvaluationMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Evaluation) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
