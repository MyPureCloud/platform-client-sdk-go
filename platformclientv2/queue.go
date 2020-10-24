package platformclientv2
import (
	"time"
	"encoding/json"
)

// Queue
type Queue struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Division - The division to which this entity belongs.
	Division *Division `json:"division,omitempty"`


	// Description - The queue description.
	Description *string `json:"description,omitempty"`


	// DateCreated - The date the queue was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`


	// DateModified - The date of the last modification to the queue. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateModified *time.Time `json:"dateModified,omitempty"`


	// ModifiedBy - The ID of the user that last modified the queue.
	ModifiedBy *string `json:"modifiedBy,omitempty"`


	// CreatedBy - The ID of the user that created the queue.
	CreatedBy *string `json:"createdBy,omitempty"`


	// MemberCount - The number of users in the queue.
	MemberCount *int32 `json:"memberCount,omitempty"`


	// MediaSettings - The media settings for the queue. Valid key values: CALL, CALLBACK, CHAT, EMAIL, MESSAGE, SOCIAL_EXPRESSION, VIDEO_COMM
	MediaSettings *map[string]Mediasetting `json:"mediaSettings,omitempty"`


	// RoutingRules - The routing rules for the queue, used for routing to known or preferred agents.
	RoutingRules *[]Routingrule `json:"routingRules,omitempty"`


	// Bullseye - The bulls-eye settings for the queue.
	Bullseye *Bullseye `json:"bullseye,omitempty"`


	// AcwSettings - The ACW settings for the queue.
	AcwSettings *Acwsettings `json:"acwSettings,omitempty"`


	// SkillEvaluationMethod - The skill evaluation method to use when routing conversations.
	SkillEvaluationMethod *string `json:"skillEvaluationMethod,omitempty"`


	// QueueFlow - The in-queue flow to use for conversations waiting in queue.
	QueueFlow *Domainentityref `json:"queueFlow,omitempty"`


	// WhisperPrompt - The prompt used for whisper on the queue, if configured.
	WhisperPrompt *Domainentityref `json:"whisperPrompt,omitempty"`


	// AutoAnswerOnly - Specifies whether the configured whisper should play for all ACD calls, or only for those which are auto-answered.
	AutoAnswerOnly *bool `json:"autoAnswerOnly,omitempty"`


	// EnableTranscription - Indicates whether voice transcription is enabled for this queue.
	EnableTranscription *bool `json:"enableTranscription,omitempty"`


	// EnableManualAssignment - Indicates whether manual assignment is enabled for this queue.
	EnableManualAssignment *bool `json:"enableManualAssignment,omitempty"`


	// CallingPartyName - The name to use for caller identification for outbound calls from this queue.
	CallingPartyName *string `json:"callingPartyName,omitempty"`


	// CallingPartyNumber - The phone number to use for caller identification for outbound calls from this queue.
	CallingPartyNumber *string `json:"callingPartyNumber,omitempty"`


	// DefaultScripts - The default script Ids for the communication types.
	DefaultScripts *map[string]Script `json:"defaultScripts,omitempty"`


	// OutboundMessagingAddresses - The messaging addresses for the queue.
	OutboundMessagingAddresses *Queuemessagingaddresses `json:"outboundMessagingAddresses,omitempty"`


	// OutboundEmailAddress
	OutboundEmailAddress *Queueemailaddress `json:"outboundEmailAddress,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Queue) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
