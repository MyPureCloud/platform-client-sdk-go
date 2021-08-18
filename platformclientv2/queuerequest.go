package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Queuerequest
type Queuerequest struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name - The queue name
	Name *string `json:"name,omitempty"`


	// Division - The division to which this entity belongs.
	Division *Writabledivision `json:"division,omitempty"`


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


	// MemberCount - The total number of members (joined or unjoined) in the queue.
	MemberCount *int `json:"memberCount,omitempty"`


	// JoinedMemberCount - The number of joined members in the queue.
	JoinedMemberCount *int `json:"joinedMemberCount,omitempty"`


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

func (u *Queuerequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Queuerequest

	
	DateCreated := new(string)
	if u.DateCreated != nil {
		
		*DateCreated = timeutil.Strftime(u.DateCreated, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateCreated = nil
	}
	
	DateModified := new(string)
	if u.DateModified != nil {
		
		*DateModified = timeutil.Strftime(u.DateModified, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateModified = nil
	}
	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Division *Writabledivision `json:"division,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		
		ModifiedBy *string `json:"modifiedBy,omitempty"`
		
		CreatedBy *string `json:"createdBy,omitempty"`
		
		MemberCount *int `json:"memberCount,omitempty"`
		
		JoinedMemberCount *int `json:"joinedMemberCount,omitempty"`
		
		MediaSettings *map[string]Mediasetting `json:"mediaSettings,omitempty"`
		
		RoutingRules *[]Routingrule `json:"routingRules,omitempty"`
		
		Bullseye *Bullseye `json:"bullseye,omitempty"`
		
		AcwSettings *Acwsettings `json:"acwSettings,omitempty"`
		
		SkillEvaluationMethod *string `json:"skillEvaluationMethod,omitempty"`
		
		QueueFlow *Domainentityref `json:"queueFlow,omitempty"`
		
		WhisperPrompt *Domainentityref `json:"whisperPrompt,omitempty"`
		
		AutoAnswerOnly *bool `json:"autoAnswerOnly,omitempty"`
		
		EnableTranscription *bool `json:"enableTranscription,omitempty"`
		
		EnableManualAssignment *bool `json:"enableManualAssignment,omitempty"`
		
		CallingPartyName *string `json:"callingPartyName,omitempty"`
		
		CallingPartyNumber *string `json:"callingPartyNumber,omitempty"`
		
		DefaultScripts *map[string]Script `json:"defaultScripts,omitempty"`
		
		OutboundMessagingAddresses *Queuemessagingaddresses `json:"outboundMessagingAddresses,omitempty"`
		
		OutboundEmailAddress *Queueemailaddress `json:"outboundEmailAddress,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Name: u.Name,
		
		Division: u.Division,
		
		Description: u.Description,
		
		DateCreated: DateCreated,
		
		DateModified: DateModified,
		
		ModifiedBy: u.ModifiedBy,
		
		CreatedBy: u.CreatedBy,
		
		MemberCount: u.MemberCount,
		
		JoinedMemberCount: u.JoinedMemberCount,
		
		MediaSettings: u.MediaSettings,
		
		RoutingRules: u.RoutingRules,
		
		Bullseye: u.Bullseye,
		
		AcwSettings: u.AcwSettings,
		
		SkillEvaluationMethod: u.SkillEvaluationMethod,
		
		QueueFlow: u.QueueFlow,
		
		WhisperPrompt: u.WhisperPrompt,
		
		AutoAnswerOnly: u.AutoAnswerOnly,
		
		EnableTranscription: u.EnableTranscription,
		
		EnableManualAssignment: u.EnableManualAssignment,
		
		CallingPartyName: u.CallingPartyName,
		
		CallingPartyNumber: u.CallingPartyNumber,
		
		DefaultScripts: u.DefaultScripts,
		
		OutboundMessagingAddresses: u.OutboundMessagingAddresses,
		
		OutboundEmailAddress: u.OutboundEmailAddress,
		
		SelfUri: u.SelfUri,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Queuerequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
