package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
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


	// MemberCount - The total number of members in the queue.
	MemberCount *int `json:"memberCount,omitempty"`


	// UserMemberCount - The number of user members (i.e., non-group members) in the queue.
	UserMemberCount *int `json:"userMemberCount,omitempty"`


	// JoinedMemberCount - The number of joined members in the queue.
	JoinedMemberCount *int `json:"joinedMemberCount,omitempty"`


	// MediaSettings - The media settings for the queue. Valid key values: CALL, CALLBACK, CHAT, EMAIL, MESSAGE, SOCIAL_EXPRESSION, VIDEO_COMM
	MediaSettings *map[string]Mediasetting `json:"mediaSettings,omitempty"`


	// RoutingRules - The routing rules for the queue, used for Preferred Agent Routing.
	RoutingRules *[]Routingrule `json:"routingRules,omitempty"`


	// Bullseye - The bullseye settings for the queue.
	Bullseye *Bullseye `json:"bullseye,omitempty"`


	// AcwSettings - The ACW settings for the queue.
	AcwSettings *Acwsettings `json:"acwSettings,omitempty"`


	// SkillEvaluationMethod - The skill evaluation method to use when routing conversations.
	SkillEvaluationMethod *string `json:"skillEvaluationMethod,omitempty"`


	// QueueFlow - The in-queue flow to use for call conversations waiting in queue.
	QueueFlow *Domainentityref `json:"queueFlow,omitempty"`


	// EmailInQueueFlow - The in-queue flow to use for email conversations waiting in queue.
	EmailInQueueFlow *Domainentityref `json:"emailInQueueFlow,omitempty"`


	// MessageInQueueFlow - The in-queue flow to use for message conversations waiting in queue.
	MessageInQueueFlow *Domainentityref `json:"messageInQueueFlow,omitempty"`


	// WhisperPrompt - The prompt used for whisper on the queue, if configured.
	WhisperPrompt *Domainentityref `json:"whisperPrompt,omitempty"`


	// OnHoldPrompt - The audio to be played when calls on this queue are on hold. If not configured, the default on-hold music will play.
	OnHoldPrompt *Domainentityref `json:"onHoldPrompt,omitempty"`


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

func (o *Queue) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Queue
	
	DateCreated := new(string)
	if o.DateCreated != nil {
		
		*DateCreated = timeutil.Strftime(o.DateCreated, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateCreated = nil
	}
	
	DateModified := new(string)
	if o.DateModified != nil {
		
		*DateModified = timeutil.Strftime(o.DateModified, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateModified = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Division *Division `json:"division,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		
		ModifiedBy *string `json:"modifiedBy,omitempty"`
		
		CreatedBy *string `json:"createdBy,omitempty"`
		
		MemberCount *int `json:"memberCount,omitempty"`
		
		UserMemberCount *int `json:"userMemberCount,omitempty"`
		
		JoinedMemberCount *int `json:"joinedMemberCount,omitempty"`
		
		MediaSettings *map[string]Mediasetting `json:"mediaSettings,omitempty"`
		
		RoutingRules *[]Routingrule `json:"routingRules,omitempty"`
		
		Bullseye *Bullseye `json:"bullseye,omitempty"`
		
		AcwSettings *Acwsettings `json:"acwSettings,omitempty"`
		
		SkillEvaluationMethod *string `json:"skillEvaluationMethod,omitempty"`
		
		QueueFlow *Domainentityref `json:"queueFlow,omitempty"`
		
		EmailInQueueFlow *Domainentityref `json:"emailInQueueFlow,omitempty"`
		
		MessageInQueueFlow *Domainentityref `json:"messageInQueueFlow,omitempty"`
		
		WhisperPrompt *Domainentityref `json:"whisperPrompt,omitempty"`
		
		OnHoldPrompt *Domainentityref `json:"onHoldPrompt,omitempty"`
		
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
		Id: o.Id,
		
		Name: o.Name,
		
		Division: o.Division,
		
		Description: o.Description,
		
		DateCreated: DateCreated,
		
		DateModified: DateModified,
		
		ModifiedBy: o.ModifiedBy,
		
		CreatedBy: o.CreatedBy,
		
		MemberCount: o.MemberCount,
		
		UserMemberCount: o.UserMemberCount,
		
		JoinedMemberCount: o.JoinedMemberCount,
		
		MediaSettings: o.MediaSettings,
		
		RoutingRules: o.RoutingRules,
		
		Bullseye: o.Bullseye,
		
		AcwSettings: o.AcwSettings,
		
		SkillEvaluationMethod: o.SkillEvaluationMethod,
		
		QueueFlow: o.QueueFlow,
		
		EmailInQueueFlow: o.EmailInQueueFlow,
		
		MessageInQueueFlow: o.MessageInQueueFlow,
		
		WhisperPrompt: o.WhisperPrompt,
		
		OnHoldPrompt: o.OnHoldPrompt,
		
		AutoAnswerOnly: o.AutoAnswerOnly,
		
		EnableTranscription: o.EnableTranscription,
		
		EnableManualAssignment: o.EnableManualAssignment,
		
		CallingPartyName: o.CallingPartyName,
		
		CallingPartyNumber: o.CallingPartyNumber,
		
		DefaultScripts: o.DefaultScripts,
		
		OutboundMessagingAddresses: o.OutboundMessagingAddresses,
		
		OutboundEmailAddress: o.OutboundEmailAddress,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Queue) UnmarshalJSON(b []byte) error {
	var QueueMap map[string]interface{}
	err := json.Unmarshal(b, &QueueMap)
	if err != nil {
		return err
	}
	
	if Id, ok := QueueMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Name, ok := QueueMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if Division, ok := QueueMap["division"].(map[string]interface{}); ok {
		DivisionString, _ := json.Marshal(Division)
		json.Unmarshal(DivisionString, &o.Division)
	}
	
	if Description, ok := QueueMap["description"].(string); ok {
		o.Description = &Description
	}
	
	if dateCreatedString, ok := QueueMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if dateModifiedString, ok := QueueMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if ModifiedBy, ok := QueueMap["modifiedBy"].(string); ok {
		o.ModifiedBy = &ModifiedBy
	}
	
	if CreatedBy, ok := QueueMap["createdBy"].(string); ok {
		o.CreatedBy = &CreatedBy
	}
	
	if MemberCount, ok := QueueMap["memberCount"].(float64); ok {
		MemberCountInt := int(MemberCount)
		o.MemberCount = &MemberCountInt
	}
	
	if UserMemberCount, ok := QueueMap["userMemberCount"].(float64); ok {
		UserMemberCountInt := int(UserMemberCount)
		o.UserMemberCount = &UserMemberCountInt
	}
	
	if JoinedMemberCount, ok := QueueMap["joinedMemberCount"].(float64); ok {
		JoinedMemberCountInt := int(JoinedMemberCount)
		o.JoinedMemberCount = &JoinedMemberCountInt
	}
	
	if MediaSettings, ok := QueueMap["mediaSettings"].(map[string]interface{}); ok {
		MediaSettingsString, _ := json.Marshal(MediaSettings)
		json.Unmarshal(MediaSettingsString, &o.MediaSettings)
	}
	
	if RoutingRules, ok := QueueMap["routingRules"].([]interface{}); ok {
		RoutingRulesString, _ := json.Marshal(RoutingRules)
		json.Unmarshal(RoutingRulesString, &o.RoutingRules)
	}
	
	if Bullseye, ok := QueueMap["bullseye"].(map[string]interface{}); ok {
		BullseyeString, _ := json.Marshal(Bullseye)
		json.Unmarshal(BullseyeString, &o.Bullseye)
	}
	
	if AcwSettings, ok := QueueMap["acwSettings"].(map[string]interface{}); ok {
		AcwSettingsString, _ := json.Marshal(AcwSettings)
		json.Unmarshal(AcwSettingsString, &o.AcwSettings)
	}
	
	if SkillEvaluationMethod, ok := QueueMap["skillEvaluationMethod"].(string); ok {
		o.SkillEvaluationMethod = &SkillEvaluationMethod
	}
	
	if QueueFlow, ok := QueueMap["queueFlow"].(map[string]interface{}); ok {
		QueueFlowString, _ := json.Marshal(QueueFlow)
		json.Unmarshal(QueueFlowString, &o.QueueFlow)
	}
	
	if EmailInQueueFlow, ok := QueueMap["emailInQueueFlow"].(map[string]interface{}); ok {
		EmailInQueueFlowString, _ := json.Marshal(EmailInQueueFlow)
		json.Unmarshal(EmailInQueueFlowString, &o.EmailInQueueFlow)
	}
	
	if MessageInQueueFlow, ok := QueueMap["messageInQueueFlow"].(map[string]interface{}); ok {
		MessageInQueueFlowString, _ := json.Marshal(MessageInQueueFlow)
		json.Unmarshal(MessageInQueueFlowString, &o.MessageInQueueFlow)
	}
	
	if WhisperPrompt, ok := QueueMap["whisperPrompt"].(map[string]interface{}); ok {
		WhisperPromptString, _ := json.Marshal(WhisperPrompt)
		json.Unmarshal(WhisperPromptString, &o.WhisperPrompt)
	}
	
	if OnHoldPrompt, ok := QueueMap["onHoldPrompt"].(map[string]interface{}); ok {
		OnHoldPromptString, _ := json.Marshal(OnHoldPrompt)
		json.Unmarshal(OnHoldPromptString, &o.OnHoldPrompt)
	}
	
	if AutoAnswerOnly, ok := QueueMap["autoAnswerOnly"].(bool); ok {
		o.AutoAnswerOnly = &AutoAnswerOnly
	}
	
	if EnableTranscription, ok := QueueMap["enableTranscription"].(bool); ok {
		o.EnableTranscription = &EnableTranscription
	}
	
	if EnableManualAssignment, ok := QueueMap["enableManualAssignment"].(bool); ok {
		o.EnableManualAssignment = &EnableManualAssignment
	}
	
	if CallingPartyName, ok := QueueMap["callingPartyName"].(string); ok {
		o.CallingPartyName = &CallingPartyName
	}
	
	if CallingPartyNumber, ok := QueueMap["callingPartyNumber"].(string); ok {
		o.CallingPartyNumber = &CallingPartyNumber
	}
	
	if DefaultScripts, ok := QueueMap["defaultScripts"].(map[string]interface{}); ok {
		DefaultScriptsString, _ := json.Marshal(DefaultScripts)
		json.Unmarshal(DefaultScriptsString, &o.DefaultScripts)
	}
	
	if OutboundMessagingAddresses, ok := QueueMap["outboundMessagingAddresses"].(map[string]interface{}); ok {
		OutboundMessagingAddressesString, _ := json.Marshal(OutboundMessagingAddresses)
		json.Unmarshal(OutboundMessagingAddressesString, &o.OutboundMessagingAddresses)
	}
	
	if OutboundEmailAddress, ok := QueueMap["outboundEmailAddress"].(map[string]interface{}); ok {
		OutboundEmailAddressString, _ := json.Marshal(OutboundEmailAddress)
		json.Unmarshal(OutboundEmailAddressString, &o.OutboundEmailAddress)
	}
	
	if SelfUri, ok := QueueMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Queue) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
